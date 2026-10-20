package services

import (
	"testing"

	"catalina-support/backend/shared/auth"
)

func TestDesarrolloSoloEscribeComentariosPrivadosEnElPrincipal(t *testing.T) {
	desarrollo := auth.Identity{Role: auth.RoleDesarrollo}
	if got := visibilidadEfectiva(desarrollo, false, VisibilityPublic); got != VisibilityPrivate {
		t.Fatalf("Desarrollo pidió public en el principal y quedó %q", got)
	}
	if got := visibilidadEfectiva(desarrollo, false, VisibilityPrivate); got != VisibilityPrivate {
		t.Fatalf("la nota privada de Desarrollo cambió a %q", got)
	}
	if got := visibilidadEfectiva(desarrollo, true, VisibilityPublic); got != VisibilityPublic {
		t.Fatalf("el comentario del ticket interno cambió a %q", got)
	}
	soporte := auth.Identity{Role: auth.RoleSoporte}
	if got := visibilidadEfectiva(soporte, false, VisibilityPublic); got != VisibilityPublic {
		t.Fatalf("el comentario público de Soporte cambió a %q", got)
	}
}
