package services

import (
	"strings"
	"testing"
)

func TestUltimaAccionExigeActorAccionYNoCopiaLaRazon(t *testing.T) {
	contexto := `[TARGET_REASON] El aviso de caducidad no existe: hay que añadirlo al módulo de acceso.
[TARGET_LAST_EVENT] cerrado
[TARGET_LAST_ACTOR] Soporte Uno
[TARGET_LAST_STATE] cerrado
[TARGET_LAST_COMMENT] El incidente quedó solucionado.`

	if ultimaAccionValida("El aviso de caducidad no existe: hay que añadirlo al módulo de acceso.", contexto) {
		t.Fatal("copiar la razón no puede ser una última acción válida")
	}
	if ultimaAccionValida("El ticket fue cerrado.", contexto) {
		t.Fatal("la última acción debe conservar el nombre visible del actor")
	}
	if !ultimaAccionValida("Soporte Uno cerró el ticket porque el incidente quedó solucionado.", contexto) {
		t.Fatal("una frase con actor y cierre debería ser válida")
	}
}

func TestRecorteJerarquizadoConservaHechoObjetivoYFinal(t *testing.T) {
	texto := "[TARGET_LAST_EVENT] cerrado\n[TARGET_LAST_ACTOR] Soporte Uno\n" + strings.Repeat("historial antiguo\n", 600) + "evento final"
	recortado := recortarJerarquizado(texto)
	if !strings.Contains(recortado, "[TARGET_LAST_EVENT] cerrado") || !strings.Contains(recortado, "evento final") {
		t.Fatalf("el recorte perdió anclas: %q", recortado)
	}
}

func TestRespaldoDeUltimaAccionCombinaEventoYComentario(t *testing.T) {
	contexto := `[TARGET_LAST_EVENT] cerrado
[TARGET_LAST_ACTOR] Soporte Uno
[TARGET_LAST_STATE] cerrado
[TARGET_LAST_COMMENT] El incidente quedó solucionado.`

	texto, hay := respaldoDeUltimaAccion(contexto, "es", 40)
	if !hay || texto != "Soporte Uno cerró el ticket: El incidente quedó solucionado." {
		t.Fatalf("respaldo inesperado: %q, hay=%v", texto, hay)
	}
}

func TestRespaldoDeUltimaAccionCubreLosMovimientosSinComentario(t *testing.T) {
	casos := map[string]string{
		"editado": "actualizó el ticket", "prioridad": "cambió la prioridad del ticket",
		"escalado": "escaló el ticket", "devuelto": "devolvió el ticket",
		"observador_anadido": "añadió un observador", "observador": "quitó un observador",
	}
	for evento, accion := range casos {
		texto, hay := respaldoDeUltimaAccion("[TARGET_LAST_EVENT] "+evento+"\n[TARGET_LAST_ACTOR] Soporte Uno", "es", 40)
		if !hay || texto != "Soporte Uno "+accion {
			t.Errorf("%s: respaldo inesperado %q (hay=%v)", evento, texto, hay)
		}
	}
}
