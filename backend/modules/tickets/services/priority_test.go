package services

import (
	"errors"
	"testing"

	"catalina-support/backend/modules/tickets/repositories"
	"catalina-support/backend/shared/auth"
)

type correoPrioridad struct {
	envios int
	fallo  error
}

func (c *correoPrioridad) Send(string, string, []string, map[string]string) error {
	c.envios++
	return c.fallo
}
func (*correoPrioridad) SendAsync(string, string, []string, map[string]string) {}
func (*correoPrioridad) Link(path string) string                               { return path }

func TestCambiarPrioridadEsTransaccionalYElAvisoEsUnico(t *testing.T) {
	servicio, tx, cuentas := servicioDeCierre(t)
	numero, id, requesterID := unPrincipalCerrable(t, tx)
	soporte := cuentaDeLaBase(t, tx, auth.RoleSoporte, true)
	solicitante := cuentaDeLaBasePorID(t, tx, requesterID)
	cuentas[soporte.ID], cuentas[solicitante.ID] = soporte, solicitante
	servicio.SetAccounts(cuentas)
	correo := &correoPrioridad{}
	servicio.SetMailer(correo)

	if err := tx.Model(&repositories.Ticket{}).Where("id = ?", id).
		Updates(map[string]any{"priority": repositories.PriorityMedium, "critical_notified_at": nil}).Error; err != nil {
		t.Fatal(err)
	}
	actor := auth.Identity{ID: soporte.ID, Role: auth.RoleSoporte}
	cambiado, aviso, err := servicio.ChangePriority(numero, repositories.PriorityCritical, actor)
	if err != nil || aviso != "" || cambiado.Priority != repositories.PriorityCritical {
		t.Fatalf("el cambio crítico falló: ticket=%+v aviso=%q err=%v", cambiado, aviso, err)
	}
	if correo.envios != 1 {
		t.Fatalf("se esperaba un aviso y salieron %d", correo.envios)
	}

	if _, _, err = servicio.ChangePriority(numero, repositories.PriorityLow, actor); err != nil {
		t.Fatal(err)
	}
	if _, _, err = servicio.ChangePriority(numero, repositories.PriorityCritical, actor); err != nil {
		t.Fatal(err)
	}
	if _, _, err = servicio.ChangePriority(numero, repositories.PriorityCritical, actor); err != nil {
		t.Fatal(err)
	}
	if correo.envios != 1 {
		t.Fatalf("volver a crítica no debe avisar otra vez: salieron %d avisos", correo.envios)
	}

	var historial int64
	if err := tx.Model(&repositories.HistoryEntry{}).
		Where("ticket_id = ? AND event = ?", id, "prioridad").Count(&historial).Error; err != nil {
		t.Fatal(err)
	}
	if historial != 3 {
		t.Fatalf("se esperaban tres cambios en historial y hay %d", historial)
	}
}

func TestPermisosParaCambiarPrioridad(t *testing.T) {
	servicio := &Service{}
	casos := []struct {
		nombre string
		numero string
		actor  auth.Identity
	}{
		{"el usuario no cambia el principal", "CS-2026-0001", auth.Identity{Role: auth.RoleUsuario}},
		{"desarrollo no cambia el principal", "CS-2026-0001", auth.Identity{Role: auth.RoleDesarrollo}},
		{"el administrador no cambia el principal", "CS-2026-0001", auth.Identity{Role: auth.RoleAdministrador}},
		{"soporte no cambia desde el interno", "INT-CS-2026-0001", auth.Identity{Role: auth.RoleSoporte}},
		{"el usuario no cambia el interno", "INT-CS-2026-0001", auth.Identity{Role: auth.RoleUsuario}},
		{"el administrador no cambia el interno", "INT-CS-2026-0001", auth.Identity{Role: auth.RoleAdministrador}},
	}
	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			if _, _, err := servicio.ChangePriority(caso.numero, repositories.PriorityHigh, caso.actor); !errors.Is(err, ErrForbidden) {
				t.Fatalf("se esperaba %v y salió %v", ErrForbidden, err)
			}
		})
	}
}

func TestDesarrolloCambiaLaPrioridadDesdeElInterno(t *testing.T) {
	servicio, tx, cuentas := servicioDeCierre(t)
	interno, principal := internoAbierto(t, tx)
	desarrollador := cuentaDeLaBase(t, tx, auth.RoleDesarrollo, true)
	cuentas[desarrollador.ID] = desarrollador
	servicio.SetAccounts(cuentas)

	cambiado, aviso, err := servicio.ChangePriority(
		interno.Number,
		repositories.PriorityHigh,
		auth.Identity{ID: desarrollador.ID, Role: auth.RoleDesarrollo},
	)
	if err != nil || aviso != "" || cambiado.Priority != repositories.PriorityHigh {
		t.Fatalf("el cambio desde el interno falló: ticket=%+v aviso=%q err=%v", cambiado, aviso, err)
	}

	var guardada string
	if err := tx.Raw(`SELECT priority FROM tickets WHERE id = ?`, principal.ID).Scan(&guardada).Error; err != nil {
		t.Fatal(err)
	}
	if guardada != repositories.PriorityHigh {
		t.Fatalf("el principal conserva %q y se esperaba %q", guardada, repositories.PriorityHigh)
	}
}

func TestFalloDelAvisoNoRevierteLaPrioridad(t *testing.T) {
	servicio, tx, cuentas := servicioDeCierre(t)
	numero, id, requesterID := unPrincipalCerrable(t, tx)
	soporte := cuentaDeLaBase(t, tx, auth.RoleSoporte, true)
	solicitante := cuentaDeLaBasePorID(t, tx, requesterID)
	cuentas[soporte.ID], cuentas[solicitante.ID] = soporte, solicitante
	servicio.SetAccounts(cuentas)
	servicio.SetMailer(&correoPrioridad{fallo: errors.New("smtp caído")})
	if err := tx.Model(&repositories.Ticket{}).Where("id = ?", id).
		Updates(map[string]any{"priority": repositories.PriorityMedium, "critical_notified_at": nil}).Error; err != nil {
		t.Fatal(err)
	}

	ticket, aviso, err := servicio.ChangePriority(numero, repositories.PriorityCritical, auth.Identity{ID: soporte.ID, Role: auth.RoleSoporte})
	if err != nil || ticket.Priority != repositories.PriorityCritical {
		t.Fatalf("el correo no debe revertir el cambio: ticket=%+v err=%v", ticket, err)
	}
	if aviso != "tickets.priority.notificationFailed" {
		t.Fatalf("se esperaba el aviso del fallo SMTP y salió %q", aviso)
	}
}

func TestUnTicketCerradoCongelaLaPrioridad(t *testing.T) {
	servicio, tx, _ := servicioDeCierre(t)
	numero, id, _ := unPrincipalCerrable(t, tx)
	soporte := cuentaDeLaBase(t, tx, auth.RoleSoporte, true)
	if err := tx.Model(&repositories.Ticket{}).Where("id = ?", id).
		Update("state", repositories.StateCerrado).Error; err != nil {
		t.Fatal(err)
	}

	if _, _, err := servicio.ChangePriority(numero, repositories.PriorityHigh, auth.Identity{ID: soporte.ID, Role: auth.RoleSoporte}); !errors.Is(err, ErrClosed) {
		t.Fatalf("se esperaba %v y salió %v", ErrClosed, err)
	}
}
