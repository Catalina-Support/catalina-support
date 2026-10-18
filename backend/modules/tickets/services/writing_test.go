package services

import (
	"errors"
	"testing"

	"catalina-support/backend/shared/auth"
)

func TestImproveDraftRechazaPapelesSinPermisoAntesDeLeerElTicket(t *testing.T) {
	servicio := &Service{}
	for _, actor := range []auth.Identity{
		{Role: auth.RoleUsuario},
		{Role: auth.RoleAdministrador},
	} {
		_, err := servicio.ImproveDraft("CS-2026-0001", "comment", "borrador", "", actor)
		if !errors.Is(err, ErrForbidden) {
			t.Fatalf("el papel %s debía recibir forbidden y obtuvo %v", actor.Role, err)
		}
	}
}

func TestImproveDraftValidaBorradorYEditorAntesDelMotor(t *testing.T) {
	servicio := &Service{}
	actor := auth.Identity{Role: auth.RoleSoporte}
	casos := []struct {
		editor, borrador string
		esperado         error
	}{
		{"comment", "   ", ErrWritingDraftRequired},
		{"subject", "texto", ErrWritingEditorInvalid},
	}
	for _, caso := range casos {
		_, err := servicio.ImproveDraft("CS-2026-0001", caso.editor, caso.borrador, "", actor)
		if !errors.Is(err, caso.esperado) {
			t.Fatalf("editor=%q borrador=%q: se esperaba %v y llegó %v",
				caso.editor, caso.borrador, caso.esperado, err)
		}
	}
}
