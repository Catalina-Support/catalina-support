package services

import (
	"testing"

	"gorm.io/gorm"

	"catalina-support/backend/modules/tickets/repositories"
	"catalina-support/backend/shared/auth"
)

// TestRespuestaDelSolicitanteReactivaPrincipal verifica el comportamiento completo contra
// PostgreSQL: comentario, estado e historial nacen juntos, y el actor del movimiento es quien
// solicitó el ticket (docs/modules/tickets.md, sección 27.3).
func TestRespuestaDelSolicitanteReactivaPrincipal(t *testing.T) {
	servicio, tx, _ := servicioDeObservadores(t)
	numero, principalID, solicitanteID := unPrincipalCerrableDeUsuario(t, tx)
	prepararPrincipalEnEspera(t, tx, principalID)
	historialAntes := contarTransicionesExactas(t, tx, principalID, repositories.StateEnEspera, repositories.StateEnProgreso)

	cuenta := cuentaDeLaBasePorID(t, tx, solicitanteID)
	actor := auth.Identity{ID: cuenta.ID, Role: cuenta.Role}
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
	if actorID != solicitanteID {
		t.Fatalf("el actor del movimiento es %d; se esperaba al solicitante %d", actorID, solicitanteID)
	}
}

// TestRespuestaVaciaDelSolicitanteReactivaPrincipal cubre el comentario que nace sin cuerpo para
// recibir un adjunto después: sigue siendo una respuesta nueva y aplica la misma transición.
func TestRespuestaVaciaDelSolicitanteReactivaPrincipal(t *testing.T) {
	servicio, tx, _ := servicioDeObservadores(t)
	numero, principalID, solicitanteID := unPrincipalCerrableDeUsuario(t, tx)
	prepararPrincipalEnEspera(t, tx, principalID)
	cuenta := cuentaDeLaBasePorID(t, tx, solicitanteID)

	if _, err := servicio.Comment(numero, "", auth.Identity{ID: cuenta.ID, Role: cuenta.Role}); err != nil {
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
	numero, principalID, _ := unPrincipalCerrableDeUsuario(t, tx)
	prepararPrincipalEnEspera(t, tx, principalID)
	soporte := cuentaDeLaBase(t, tx, auth.RoleSoporte, true)

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
	numero, principalID, solicitanteID := unPrincipalCerrableDeUsuario(t, tx)
	cuenta := cuentaDeLaBasePorID(t, tx, solicitanteID)
	actor := auth.Identity{ID: cuenta.ID, Role: cuenta.Role}

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
