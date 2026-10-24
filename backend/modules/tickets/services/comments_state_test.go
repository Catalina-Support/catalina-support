package services

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/gorm"

	"catalina-support/backend/modules/tickets/repositories"
	"catalina-support/backend/shared/auth"
)

// TestRespuestaDelSolicitanteReactivaPrincipal verifica el comportamiento completo contra
// PostgreSQL: comentario, estado e historial nacen juntos, y el actor del movimiento es quien
// solicitó el ticket (docs/modules/tickets.md, sección 27.3).
func TestRespuestaDelSolicitanteReactivaPrincipal(t *testing.T) {
	servicio, tx, _ := servicioDeObservadores(t)
	numero, principalID, solicitante := crearPrincipalDeUsuario(t, tx)
	servicio.SetAccounts(cuentasStub{solicitante.ID: solicitante})
	prepararPrincipalEnEspera(t, tx, principalID)
	historialAntes := contarTransicionesExactas(t, tx, principalID, repositories.StateEnEspera, repositories.StateEnProgreso)

	actor := auth.Identity{ID: solicitante.ID, Role: solicitante.Role}
	if _, err := servicio.Comment(numero, "Ya envié la información solicitada", actor); err != nil {
		t.Fatalf("el comentario del solicitante falló: %v", err)
	}

	if estado := estadoDelPrincipal(t, tx, principalID); estado != repositories.StateEnProgreso {
		t.Fatalf("el ticket quedó en %q; se esperaba %q", estado, repositories.StateEnProgreso)
	}
	if despues := contarTransicionesExactas(t, tx, principalID, repositories.StateEnEspera, repositories.StateEnProgreso); despues != historialAntes+1 {
		t.Fatalf("se esperaba una transición nueva y hubo %d (antes %d)", despues, historialAntes)
	}

	var actorID int64
	if err := tx.Raw(`SELECT actor_id FROM ticket_history
		WHERE ticket_id = ? AND from_state = ? AND to_state = ?
		ORDER BY id DESC LIMIT 1`, principalID, repositories.StateEnEspera, repositories.StateEnProgreso).
		Scan(&actorID).Error; err != nil {
		t.Fatalf("no se pudo leer el actor del historial: %v", err)
	}
	if actorID != solicitante.ID {
		t.Fatalf("el actor del movimiento es %d; se esperaba al solicitante %d", actorID, solicitante.ID)
	}
}

// TestRespuestaVaciaDelSolicitanteReactivaPrincipal cubre el comentario que nace sin cuerpo para
// recibir un adjunto después: sigue siendo una respuesta nueva y aplica la misma transición.
func TestRespuestaVaciaDelSolicitanteReactivaPrincipal(t *testing.T) {
	servicio, tx, _ := servicioDeObservadores(t)
	numero, principalID, solicitante := crearPrincipalDeUsuario(t, tx)
	servicio.SetAccounts(cuentasStub{solicitante.ID: solicitante})
	prepararPrincipalEnEspera(t, tx, principalID)

	if _, err := servicio.Comment(numero, "", auth.Identity{ID: solicitante.ID, Role: solicitante.Role}); err != nil {
		t.Fatalf("el comentario vacío del solicitante falló: %v", err)
	}
	if estado := estadoDelPrincipal(t, tx, principalID); estado != repositories.StateEnProgreso {
		t.Fatalf("el comentario vacío dejó el ticket en %q", estado)
	}
}

// TestComentarioAjenoNoReactivaPrincipal deja fijada la identidad exacta: ni siquiera otra persona
// autorizada a comentar en el principal sustituye al solicitante para esta automatización.
func TestComentarioAjenoNoReactivaPrincipal(t *testing.T) {
	servicio, tx, _ := servicioDeObservadores(t)
	numero, principalID, solicitante := crearPrincipalDeUsuario(t, tx)
	prepararPrincipalEnEspera(t, tx, principalID)
	soporte := crearCuentaParaComentario(t, tx, auth.RoleSoporte)
	servicio.SetAccounts(cuentasStub{solicitante.ID: solicitante, soporte.ID: soporte})

	if _, err := servicio.Comment(numero, "Seguimos pendientes de tu respuesta", auth.Identity{ID: soporte.ID, Role: soporte.Role}); err != nil {
		t.Fatalf("el comentario de Soporte falló: %v", err)
	}
	if estado := estadoDelPrincipal(t, tx, principalID); estado != repositories.StateEnEspera {
		t.Fatalf("un comentario ajeno cambió el estado a %q", estado)
	}
}

// TestRespuestaEnInternoNoReactivaPrincipal evita extender la regla al hilo interno: aunque quien
// comenta coincida con el solicitante, ese hilo tiene un ciclo de vida independiente.
func TestRespuestaEnInternoNoReactivaPrincipal(t *testing.T) {
	servicio, tx, _ := servicioDeObservadores(t)
	interno, principal := unInterno(t, tx)
	prepararPrincipalEnEspera(t, tx, principal.ID)

	// Para aislar la regla del tipo de hilo, la prueba hace coincidir temporalmente solicitante y un
	// actor autorizado en el interno. Todo se deshace al terminar la transacción de prueba.
	desarrollo := cuentaDeLaBase(t, tx, auth.RoleDesarrollo, true)
	if err := tx.Exec(`UPDATE tickets SET requester_id = ? WHERE id = ?`, desarrollo.ID, principal.ID).Error; err != nil {
		t.Fatalf("no se pudo preparar el solicitante: %v", err)
	}
	if err := tx.Exec(`UPDATE internal_tickets SET state = ? WHERE id = ?`, repositories.StateEnProgreso, interno.ID).Error; err != nil {
		t.Fatalf("no se pudo preparar el interno: %v", err)
	}

	if _, err := servicio.Comment(interno.Number, "Comentario en el hilo interno", auth.Identity{ID: desarrollo.ID, Role: desarrollo.Role}); err != nil {
		t.Fatalf("el comentario del interno falló: %v", err)
	}
	if estado := estadoDelPrincipal(t, tx, principal.ID); estado != repositories.StateEnEspera {
		t.Fatalf("un comentario interno cambió el principal a %q", estado)
	}
}

// TestEditarComentarioNoReactivaPrincipal demuestra que la automatización pertenece sólo al alta
// del comentario. Editar después una respuesta anterior no cuenta como una respuesta nueva.
func TestEditarComentarioNoReactivaPrincipal(t *testing.T) {
	servicio, tx, _ := servicioDeObservadores(t)
	numero, principalID, solicitante := crearPrincipalDeUsuario(t, tx)
	servicio.SetAccounts(cuentasStub{solicitante.ID: solicitante})
	actor := auth.Identity{ID: solicitante.ID, Role: solicitante.Role}

	if err := tx.Exec(`UPDATE tickets SET state = ? WHERE id = ?`, repositories.StateEnProgreso, principalID).Error; err != nil {
		t.Fatalf("no se pudo preparar el principal: %v", err)
	}
	comentario, err := servicio.Comment(numero, "Respuesta anterior", actor)
	if err != nil {
		t.Fatalf("no se pudo crear el comentario que se editará: %v", err)
	}
	prepararPrincipalEnEspera(t, tx, principalID)

	if _, err := servicio.EditComment(numero, comentario.ID, "Respuesta anterior corregida", actor); err != nil {
		t.Fatalf("no se pudo editar el comentario: %v", err)
	}
	if estado := estadoDelPrincipal(t, tx, principalID); estado != repositories.StateEnEspera {
		t.Fatalf("editar el comentario cambió el estado a %q", estado)
	}
}

// crearPrincipalDeUsuario construye todo el contrato que necesitan estas pruebas dentro de su
// transacción. No lee tickets ni cuentas de seeders y el rollback del helper de servicio lo elimina.
func crearPrincipalDeUsuario(t *testing.T, tx *gorm.DB) (string, int64, auth.Account) {
	t.Helper()

	solicitante := crearCuentaParaComentario(t, tx, auth.RoleUsuario)
	var categoriaID int64
	if err := tx.Raw(`SELECT id FROM ticket_categories WHERE normalized = 'general'`).Scan(&categoriaID).Error; err != nil {
		t.Fatalf("no se pudo leer la categoría de fábrica: %v", err)
	}
	if categoriaID == 0 {
		t.Fatal("no existe la categoría de fábrica de la migración")
	}

	var secuencia int
	if err := tx.Raw(`SELECT COALESCE(MAX(number_seq), 0) + 1 FROM tickets WHERE number_year = 2099`).Scan(&secuencia).Error; err != nil {
		t.Fatalf("no se pudo reservar el número del ticket de prueba: %v", err)
	}
	numero := fmt.Sprintf("TEST-2099-%04d", secuencia)
	var principalID int64
	if err := tx.Raw(`INSERT INTO tickets
		(number, number_year, number_seq, subject, description, state, priority, category_id, requester_id, created_by_id)
		VALUES (?, 2099, ?, 'Respuesta del solicitante', 'Contrato aislado de prueba', 'nuevo', 'medium', ?, ?, ?)
		RETURNING id`, numero, secuencia, categoriaID, solicitante.ID, solicitante.ID).Scan(&principalID).Error; err != nil {
		t.Fatalf("no se pudo crear el ticket de prueba: %v", err)
	}

	return numero, principalID, solicitante
}

func crearCuentaParaComentario(t *testing.T, tx *gorm.DB, rol string) auth.Account {
	t.Helper()
	unico := time.Now().UnixNano()
	cuenta := auth.Account{
		Name: "Prueba", LastName: "Comentarios", Role: rol, Origin: auth.OriginLocal, IsActive: true,
		Email: fmt.Sprintf("comentarios-%s-%d@example.test", rol, unico),
	}
	if err := tx.Raw(`INSERT INTO users (name, last_name, email, role, origin, is_active)
		VALUES (?, ?, ?, ?, ?, true) RETURNING id`, cuenta.Name, cuenta.LastName, cuenta.Email, cuenta.Role, cuenta.Origin).
		Scan(&cuenta.ID).Error; err != nil {
		t.Fatalf("no se pudo crear la cuenta %s de prueba: %v", rol, err)
	}
	return cuenta
}

func prepararPrincipalEnEspera(t *testing.T, tx *gorm.DB, principalID int64) {
	t.Helper()
	if err := tx.Exec(`UPDATE tickets SET state = ? WHERE id = ?`, repositories.StateEnEspera, principalID).Error; err != nil {
		t.Fatalf("no se pudo dejar el principal en espera: %v", err)
	}
}

func contarTransicionesExactas(t *testing.T, tx *gorm.DB, principalID int64, desde, hasta string) int64 {
	t.Helper()
	var total int64
	if err := tx.Raw(`SELECT count(*) FROM ticket_history
		WHERE ticket_id = ? AND from_state = ? AND to_state = ?`, principalID, desde, hasta).Scan(&total).Error; err != nil {
		t.Fatalf("no se pudo contar el historial de estado: %v", err)
	}
	return total
}
