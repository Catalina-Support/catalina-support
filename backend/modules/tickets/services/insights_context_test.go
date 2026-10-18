package services

import (
	"testing"
	"time"

	"catalina-support/backend/modules/tickets/repositories"
	"catalina-support/backend/shared/auth"
)

func TestUltimoHechoUsaElEventoFinalYSuComentarioVinculado(t *testing.T) {
	antes := time.Date(2026, 9, 8, 10, 0, 0, 0, time.UTC)
	despues := antes.Add(time.Minute)
	comentarios := []repositories.Comment{{ID: 7, AuthorID: 6, Body: "El incidente quedó solucionado.", CreatedAt: antes}}
	historial := []repositories.HistoryEntry{{ActorID: punteroInt64(6), Event: "cerrado", ToState: punteroTexto("cerrado"), CommentID: punteroInt64(7), CreatedAt: despues}}
	cuentas := map[int64]auth.Account{6: {ID: 6, Name: "Soporte", LastName: "Uno"}}

	hecho := ultimoHecho(comentarios, historial, cuentas, "es")
	if hecho.evento != "cerrado" || hecho.actor != "Soporte Uno" || hecho.estado != "cerrado" || hecho.comentario != "El incidente quedó solucionado." {
		t.Fatalf("hecho final incorrecto: %+v", hecho)
	}
}

func punteroInt64(valor int64) *int64   { return &valor }
func punteroTexto(valor string) *string { return &valor }
