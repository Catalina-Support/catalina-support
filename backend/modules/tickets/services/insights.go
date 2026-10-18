package services

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/juanky-estevez/go-logs"

	"catalina-support/backend/modules/tickets/repositories"
	"catalina-support/backend/shared/auth"
)

// Los dos campos que redacta el motor de IA (`docs/modules/ai.md`). Son los mismos nombres que usa
// su tabla, y se declaran aquí para no escribirlos de dos maneras.
const (
	TipoSummary      = "summary"
	TipoUltimaAccion = "ultima_accion"
)

// Resumen es un texto del motor en el idioma global, con su estado y la clave del error.
type Resumen struct {
	Estado   string
	Text     string
	Language string
	ErrorKey string
}

// Resumenes son los dos campos de un ticket.
type Resumenes struct {
	Summary      Resumen
	UltimaAccion Resumen
}

// EntradaDeResumen es el encargo que se le hace al motor: un ticket, un campo y el texto ya armado.
type EntradaDeResumen struct {
	Numero string
	// Tipo es `motivo` o `ultima_accion`.
	Tipo string
	// Texto lo arma este módulo, que es quien tiene el ticket: asunto, descripción, adjuntos y
	// conversación, ya en texto plano (`textoParaElMotor`).
	Texto string
}

// Insights es lo que este módulo necesita del motor de IA.
//
// Se declara aquí, en quien lo usa, y lo cumple el módulo `ai` sin saber que existe: así ninguno de
// los dos importa al otro y quien los une es `main.go` (`docs/arquitectura.md`, sección 4). La misma
// forma que tienen `Accounts` y `Mailer`.
type Insights interface {
	// Pedir apunta un resumen y **vuelve enseguida**: no espera al motor (decisión 2).
	Pedir(EntradaDeResumen)
	// De devuelve lo que hay de esos tickets, **en bloque**: una consulta por página, no una por fila
	// (decisión 9).
	De(numeros []string) (map[string]Resumenes, error)
}

// SetInsights engancha el motor de IA. Si no se llama —o se llama con nada—, los dos campos se quedan
// sin texto y **todo lo demás funciona igual**: la mesa de ayuda no depende del motor.
func (s *Service) SetInsights(insights Insights) {
	s.insights = insights
}

// pedirAlMotor le pide al motor los resúmenes de ese ticket, con el texto ya armado.
//
// **No devuelve error a propósito**: quien crea un ticket, comenta o lo mueve no puede quedarse
// esperando a un contenedor que puede estar caído. Si el texto no se puede armar, o el motor no está,
// se deja dicho en el log y el campo se queda como estaba (decisión 2).
func (s *Service) pedirAlMotor(numero string, tipos ...string) {
	if s.insights == nil {
		return
	}

	texto, err := s.textoParaElMotor(numero)
	if err != nil {
		logs.LogWarning("no se pudo armar el texto para el resumen del ticket " + numero)
		return
	}

	for _, tipo := range tipos {
		s.insights.Pedir(EntradaDeResumen{Numero: numero, Tipo: tipo, Texto: texto})
	}
}

// TextoParaElMotor arma el texto de un ticket, y lo usa el módulo `ai` para **volver a pedir** un
// resumen que quedó a medias cuando el backend se reinició: la cola vive en memoria y el texto no se
// guarda, así que hay que volver a construirlo (docs/modules/ai.md, decisión 10).
//
// Se declara aquí y `ai` lo cumple por su cuenta: ninguno de los dos módulos importa al otro.
func (s *Service) TextoParaElMotor(numero string) (string, error) {
	return s.textoParaElMotor(numero)
}

// NumerosParaRegenerarInsights enumera los dos tipos sin cruzar la frontera de datos: `ai`
// decide versiones, pero `tickets` sigue siendo quien conoce dónde viven sus números.
func (s *Service) NumerosParaRegenerarInsights() ([]string, error) {
	type fila struct{ Number string }
	var filas []fila
	err := s.tickets.DB().Raw(`
		SELECT number FROM tickets
		UNION ALL
		SELECT number FROM internal_tickets
		ORDER BY number
	`).Scan(&filas).Error
	if err != nil {
		return nil, err
	}
	numeros := make([]string, 0, len(filas))
	for _, actual := range filas {
		numeros = append(numeros, actual.Number)
	}
	return numeros, nil
}

// textoParaElMotor arma lo que lee el motor: **el ticket en texto plano**.
//
// Lo que se manda y por qué está en `docs/modules/ai.md`, sección 3: el asunto, la descripción y la
// conversación —comentarios e historial, de lo más antiguo a lo más reciente—, **el HTML convertido a
// texto** y **los adjuntos por su nombre**, nunca su contenido. El recorte, si el texto es largo, lo
// hace el módulo `ai` (decisión 6), que es quien conoce el tope: aquí se manda lo que hay.
func (s *Service) textoParaElMotor(numero string) (string, error) {
	idioma := "es"
	if s.language != nil {
		if actual, err := s.language.Language(); err == nil && actual == "en" {
			idioma = "en"
		}
	}
	principal, esInterno, hay, err := s.porNumero(numero)
	if err != nil {
		return "", err
	}

	interno, tieneInterno, err := s.tickets.InternalByTicket(principal.ID)
	if err != nil {
		return "", err
	}

	destinoPrincipal := repositories.DestinoDePrincipal(principal.ID)
	comentariosPrincipal, adjuntosPrincipal, historialPrincipal, err := s.conversacionDe(destinoPrincipal)
	if err != nil {
		return "", err
	}

	var comentariosInterno []repositories.Comment
	var adjuntosInterno []repositories.Attachment
	var historialInterno []repositories.HistoryEntry
	if tieneInterno {
		comentariosInterno, adjuntosInterno, historialInterno, err = s.conversacionDe(repositories.DestinoDeInterno(interno.ID))
		if err != nil {
			return "", err
		}
	}
	observadoresPrincipal, err := s.conversation.Observers(s.tickets.DB(), destinoPrincipal)
	if err != nil {
		return "", err
	}
	var observadoresInterno []repositories.TicketObserver
	if tieneInterno {
		observadoresInterno, err = s.conversation.Observers(s.tickets.DB(), repositories.DestinoDeInterno(interno.ID))
		if err != nil {
			return "", err
		}
	}

	comentarios := comentariosPrincipal
	adjuntos := adjuntosPrincipal
	historial := historialPrincipal
	estadoObjetivo := principal.State
	razonObjetivo := TextoPlano(principal.Description)
	if esInterno {
		comentarios = comentariosInterno
		adjuntos = adjuntosInterno
		historial = historialInterno
		estadoObjetivo = interno.State
		razonObjetivo = TextoPlano(interno.EscalationReason)
	}

	// Los nombres de las personas que aparecen: para el motor, «Soporte pidió una captura» dice más que
	// «la cuenta 42 pidió una captura». Son los mismos datos que ya se ven en el ticket.
	ids := []int64{principal.RequesterID, principal.CreatedByID}
	if principal.AssigneeID != nil {
		ids = append(ids, *principal.AssigneeID)
	}
	if tieneInterno {
		ids = append(ids, interno.CreatedByID)
		if interno.AssigneeID != nil {
			ids = append(ids, *interno.AssigneeID)
		}
	}
	for _, grupo := range [][]repositories.Comment{comentariosPrincipal, comentariosInterno} {
		for i := range grupo {
			ids = append(ids, grupo[i].AuthorID)
		}
	}
	for _, grupo := range [][]repositories.HistoryEntry{historialPrincipal, historialInterno} {
		for i := range grupo {
			if grupo[i].ActorID != nil {
				ids = append(ids, *grupo[i].ActorID)
			}
		}
	}
	for _, grupo := range [][]repositories.TicketObserver{observadoresPrincipal, observadoresInterno} {
		for _, observador := range grupo {
			ids = append(ids, observador.AccountID, observador.AddedByID)
		}
	}

	porID, err := s.accounts.ByIDs(ids)
	if err != nil {
		return "", err
	}

	ultimo := ultimoHecho(comentarios, historial, porID, idioma)
	var texto strings.Builder
	fmt.Fprintf(&texto, "[TARGET_NUMBER] %s\n", numero)
	fmt.Fprintf(&texto, "[TARGET_STATE] %s\n", estadoObjetivo)
	fmt.Fprintf(&texto, "[TARGET_REASON] %s\n", razonObjetivo)
	fmt.Fprintf(&texto, "[TARGET_LAST_EVENT] %s\n", ultimo.evento)
	fmt.Fprintf(&texto, "[TARGET_LAST_ACTOR] %s\n", ultimo.actor)
	fmt.Fprintf(&texto, "[TARGET_LAST_STATE] %s\n", ultimo.estado)
	fmt.Fprintf(&texto, "[TARGET_LAST_COMMENT] %s\n\n", ultimo.comentario)
	fmt.Fprintf(&texto, "Requester: %s\nCreated by: %s\n", nombreDeCuenta(principal.RequesterID, porID), nombreDeCuenta(principal.CreatedByID, porID))
	if principal.AssigneeID != nil {
		fmt.Fprintf(&texto, "Main assignee: %s\n", nombreDeCuenta(*principal.AssigneeID, porID))
	}
	if tieneInterno {
		fmt.Fprintf(&texto, "Internal created by: %s\n", nombreDeCuenta(interno.CreatedByID, porID))
		if interno.AssigneeID != nil {
			fmt.Fprintf(&texto, "Internal assignee: %s\n", nombreDeCuenta(*interno.AssigneeID, porID))
		}
	}

	if esInterno {
		escribirContexto(&texto, "INTERNAL", interno.Number, interno.State, principal.Subject,
			"", interno.EscalationReason, comentariosInterno, adjuntosInterno, historialInterno, porID, idioma)
		escribirFechas(&texto, interno.CreatedAt, interno.UpdatedAt, interno.ResolvedAt, interno.ClosedAt)
		escribirObservadores(&texto, observadoresInterno, porID)
	}
	escribirContexto(&texto, "MAIN", principal.Number, principal.State, principal.Subject,
		principal.Description, "", comentariosPrincipal, adjuntosPrincipal, historialPrincipal, porID, idioma)
	escribirFechas(&texto, principal.CreatedAt, principal.UpdatedAt, principal.ResolvedAt, principal.ClosedAt)
	escribirObservadores(&texto, observadoresPrincipal, porID)
	if tieneInterno && !esInterno {
		escribirContexto(&texto, "INTERNAL", interno.Number, interno.State, principal.Subject,
			"", interno.EscalationReason, comentariosInterno, adjuntosInterno, historialInterno, porID, idioma)
		escribirFechas(&texto, interno.CreatedAt, interno.UpdatedAt, interno.ResolvedAt, interno.ClosedAt)
		escribirObservadores(&texto, observadoresInterno, porID)
	}

	if s.categories != nil {
		if categoria, e := s.categories.CategoryByID(principal.CategoryID); e == nil {
			fmt.Fprintf(&texto, "\nCategory: %s\n", categoria.Name)
		}
		if tags, e := s.categories.TagsByTicket(principal.ID); e == nil && len(tags) > 0 {
			fmt.Fprintf(&texto, "Tags: %s\n", strings.Join(tags, ", "))
		}
	}
	fmt.Fprintf(&texto, "Priority: %s\n", principal.Priority)
	_ = hay
	_ = adjuntos

	return texto.String(), nil
}

func escribirFechas(texto *strings.Builder, creada, actualizada time.Time, resuelta, cerrada *time.Time) {
	fmt.Fprintf(texto, "Created at: %s\nUpdated at: %s\n", creada.Format(time.RFC3339), actualizada.Format(time.RFC3339))
	if resuelta != nil {
		fmt.Fprintf(texto, "Resolved at: %s\n", resuelta.Format(time.RFC3339))
	}
	if cerrada != nil {
		fmt.Fprintf(texto, "Closed at: %s\n", cerrada.Format(time.RFC3339))
	}
}

func nombreDeCuenta(id int64, porID map[int64]auth.Account) string {
	if cuenta, ok := porID[id]; ok {
		return cuenta.FullName()
	}
	return ""
}

func escribirObservadores(texto *strings.Builder, observadores []repositories.TicketObserver, porID map[int64]auth.Account) {
	nombres := make([]string, 0, len(observadores))
	for _, observador := range observadores {
		if cuenta, ok := porID[observador.AccountID]; ok {
			nombres = append(nombres, cuenta.FullName())
		}
	}
	if len(nombres) > 0 {
		fmt.Fprintf(texto, "Observers: %s\n", strings.Join(nombres, ", "))
	}
}

type hechoObjetivo struct{ evento, actor, estado, comentario string }

func ultimoHecho(comentarios []repositories.Comment, historial []repositories.HistoryEntry, porID map[int64]auth.Account, idioma string) hechoObjetivo {
	ultimo := hechoObjetivo{}
	var cuando time.Time
	porComentario := make(map[int64]repositories.Comment, len(comentarios))
	for _, comentario := range comentarios {
		porComentario[comentario.ID] = comentario
		if comentario.CreatedAt.After(cuando) {
			cuando = comentario.CreatedAt
			ultimo = hechoObjetivo{"comentario", quienEs(&comentario.AuthorID, porID, idioma), "", TextoPlano(comentario.Body)}
		}
	}
	for _, entrada := range historial {
		if entrada.CreatedAt.Before(cuando) {
			continue
		}
		cuando = entrada.CreatedAt
		ultimo = hechoObjetivo{entrada.Event, quienEs(entrada.ActorID, porID, idioma), textoDePuntero(entrada.ToState), ""}
		if entrada.CommentID != nil {
			if comentario, ok := porComentario[*entrada.CommentID]; ok && comentario.DeletedAt == nil {
				ultimo.comentario = TextoPlano(comentario.Body)
			}
		}
	}
	return ultimo
}

func textoDePuntero(valor *string) string {
	if valor == nil {
		return ""
	}
	return *valor
}

func escribirContexto(texto *strings.Builder, tipo, numero, estado, asunto, descripcion, motivo string,
	comentarios []repositories.Comment, adjuntos []repositories.Attachment, historial []repositories.HistoryEntry,
	porID map[int64]auth.Account, idioma string) {
	labels := etiquetasDelContexto(idioma)
	fmt.Fprintf(texto, "[%s TICKET]\nNumber: %s\n%s: %s\n%s: %s\n", tipo, numero, labels.estado, estado, labels.asunto, asunto)
	if descripcion != "" {
		fmt.Fprintf(texto, "%s: %s\n", labels.descripcion, TextoPlano(descripcion))
	}
	if motivo != "" {
		fmt.Fprintf(texto, "%s: %s\n", labels.escalado, TextoPlano(motivo))
	}
	if nombres := nombresDe(adjuntos); nombres != "" {
		fmt.Fprintf(texto, "%s: %s\n", labels.adjuntos, nombres)
	}
	fmt.Fprintf(texto, "%s:\n", labels.conversacion)
	for _, renglon := range lineasDe(comentarios, historial, porID, idioma) {
		fmt.Fprintf(texto, "- %s\n", renglon)
	}
}

type etiquetasContexto struct{ asunto, descripcion, estado, escalado, adjuntos, conversacion string }

func etiquetasDelContexto(idioma string) etiquetasContexto {
	if idioma == "en" {
		return etiquetasContexto{"Subject", "Description", "State", "Escalation reason", "Attachments", "Conversation (oldest to newest)"}
	}
	return etiquetasContexto{"Asunto", "Descripción", "Estado", "Summary del escalado", "Adjuntos", "Conversación (de lo más antiguo a lo más reciente)"}
}

// nombresDe junta los nombres de los adjuntos, sin repetir y en el orden en que se subieron.
func nombresDe(adjuntos []repositories.Attachment) string {
	vistos := make(map[string]bool, len(adjuntos))
	nombres := make([]string, 0, len(adjuntos))

	for _, adjunto := range adjuntos {
		if adjunto.Filename == "" || vistos[adjunto.Filename] {
			continue
		}

		vistos[adjunto.Filename] = true
		nombres = append(nombres, adjunto.Filename)
	}

	return strings.Join(nombres, ", ")
}

// lineasDe intercala los comentarios y lo que hizo el sistema, como la línea de tiempo de la pantalla.
func lineasDe(
	comentarios []repositories.Comment,
	historial []repositories.HistoryEntry,
	porID map[int64]auth.Account,
	idioma string,
) []string {
	type conFecha struct {
		cuando time.Time
		texto  string
	}

	renglones := make([]conFecha, 0, len(comentarios)+len(historial))

	for i := range comentarios {
		comentario := comentarios[i]
		cuerpo := TextoPlano(comentario.Body)
		if comentario.DeletedAt != nil {
			cuerpo = "(comentario borrado)"
			if idioma == "en" {
				cuerpo = "(deleted comment)"
			}
		}

		renglones = append(renglones, conFecha{
			cuando: comentario.CreatedAt,
			texto:  fmt.Sprintf("%s (%s): %s", quienEs(&comentario.AuthorID, porID, idioma), map[bool]string{true: "comment", false: "comentario"}[idioma == "en"], cuerpo),
		})
	}

	for i := range historial {
		entrada := historial[i]

		detalle := ""
		if entrada.Detail != nil {
			detalle = *entrada.Detail
		}

		renglones = append(renglones, conFecha{
			cuando: entrada.CreatedAt,
			texto:  fmt.Sprintf("%s (%s): %s", quienEs(entrada.ActorID, porID, idioma), nombreDeEvento(entrada.Event, idioma), detalle),
		})
	}

	sort.SliceStable(renglones, func(i, j int) bool { return renglones[i].cuando.Before(renglones[j].cuando) })

	lineas := make([]string, 0, len(renglones))
	for _, renglon := range renglones {
		lineas = append(lineas, strings.TrimSpace(renglon.texto))
	}

	return lineas
}

func nombreDeEvento(evento, idioma string) string {
	if idioma != "en" {
		return evento
	}
	traducidos := map[string]string{
		"estado": "state changed", "resuelto": "resolved", "cerrado": "closed", "reabierto": "reopened",
		"creado": "created", "asignado": "assigned", "escalado": "escalated", "devuelto": "returned",
		"observador_agregado": "watcher added", "observador_eliminado": "watcher removed",
	}
	if traducido, ok := traducidos[evento]; ok {
		return traducido
	}
	return evento
}

// quienEs es el nombre de esa cuenta, o «el sistema» cuando la entrada no tiene autor.
func quienEs(id *int64, porID map[int64]auth.Account, idioma string) string {
	if id == nil {
		if idioma == "en" {
			return "the system"
		}
		return "el sistema"
	}

	cuenta, hay := porID[*id]
	if !hay {
		if idioma == "en" {
			return "someone"
		}
		return "alguien"
	}

	return cuenta.FullName()
}

// pegarResumenes escribe los dos campos en una lista de tickets, **de una sola consulta**.
//
// Recibe la lista por su sitio en memoria —un `[]Ticket`— y no una copia: es lo que hace que el mismo
// trabajo valga para una página entera y para un solo ticket.
func (s *Service) pegarResumenes(tickets []Ticket) {
	if len(tickets) == 0 {
		return
	}

	numeros := make([]string, 0, len(tickets))
	for i := range tickets {
		numeros = append(numeros, tickets[i].Number)
	}

	resumenes := s.resumenesDe(numeros)
	for i := range tickets {
		if resumen, hay := resumenes[tickets[i].Number]; hay {
			tickets[i].Resumen = resumen
		}
	}
}

// resumenesDe lee los dos campos de una lista de tickets, **en una sola consulta**.
//
// Va aparte de `aTicket` porque la lista se arma en dos pasadas: primero los tickets y después sus
// resúmenes, para no preguntar una vez por fila (decisión 9). Si el motor no está, devuelve el mapa
// vacío y los campos se quedan sin texto.
func (s *Service) resumenesDe(numeros []string) map[string]Resumenes {
	if s.insights == nil || len(numeros) == 0 {
		return nil
	}

	resumenes, err := s.insights.De(numeros)
	if err != nil {
		logs.LogWarning("no se pudieron leer los resúmenes de los tickets")
		return nil
	}

	return resumenes
}
