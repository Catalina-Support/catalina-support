package services

import (
	"errors"
	"strings"
	"time"

	"github.com/juanky-estevez/go-logs"
)

// trabajo es lo que espera en la cola: un ticket, un campo y el texto que hay que leer.
type trabajo struct {
	numero       string
	tipo         Tipo
	texto        string
	generacion   int64
	recuperacion bool
	noAntes      time.Time
}

// clave identifica un trabajo: **uno por ticket y campo**. Pedir dos veces lo mismo no encola dos
// veces, y el encargo nuevo sustituye al que esperaba porque el texto viejo ya no sirve
// (docs/modules/ai.md, sección 6).
func (t trabajo) clave() string { return t.numero + "|" + string(t.tipo) }

// encolar deja el trabajo esperando y despierta a un trabajador.
//
// **No bloquea nunca**, y por eso el aviso es una señal y no el trabajo: si ya hay uno pendiente, el
// trabajador vaciará la cola entera cuando llegue, y quien pide un resumen no espera a nadie.
func (s *Service) encolar(t trabajo) {
	s.mu.Lock()
	if s.cerrada {
		s.mu.Unlock()
		return
	}
	s.cola[t.clave()] = t
	s.mu.Unlock()

	select {
	case s.aviso <- struct{}{}:
	default:
	}
}

// programar conserva la espera fuera de los trabajadores: dos recuperaciones diferidas no pueden
// ocupar los dos cupos del motor e impedir que otro ticket nuevo avance.
func (s *Service) programar(t trabajo) {
	espera := time.Until(t.noAntes)
	if t.noAntes.IsZero() || espera <= 0 {
		s.encolar(t)
		return
	}

	s.mu.Lock()
	if s.cerrada {
		s.mu.Unlock()
		return
	}
	s.enMarcha.Add(1)
	s.mu.Unlock()

	go func() {
		defer s.enMarcha.Done()
		if s.esperar(espera) {
			s.encolar(t)
		}
	}()
}

// siguiente saca el trabajo que toque.
//
// El orden no importa —son resúmenes sueltos y ninguno corre más que otro— y por eso no se mantiene
// una lista aparte del mapa: lo que importa es que no haya dos encargos del mismo campo a la vez.
func (s *Service) siguiente() (trabajo, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for clave, t := range s.cola {
		delete(s.cola, clave)
		return t, true
	}

	return trabajo{}, false
}

// trabajador vacía la cola mientras haya trabajo.
func (s *Service) trabajador() {
	defer s.enMarcha.Done()

	for {
		select {
		case <-s.parar:
			return
		case <-s.aviso:
		}

		for {
			t, hay := s.siguiente()
			if !hay {
				break
			}

			s.atender(t)
		}
	}
}

// atender le pide un campo al motor, con sus reintentos.
//
// Los dos fallos no se tratan igual, y es la decisión 8 tal cual: **un motor que no contesta se
// reintenta tres veces con espera creciente** —es un contenedor que puede estar arrancando, o lento—
// y **una respuesta que no vale se reintenta una sola vez** y después el campo queda en `error`,
// porque volver a preguntarle lo mismo a un modelo que ya contestó mal es pedirle que se repita.
func (s *Service) atender(t trabajo) {
	if !t.noAntes.IsZero() {
		espera := time.Until(t.noAntes)
		if espera > 0 && !s.esperar(espera) {
			return
		}
		if !s.trabajoVigente(t) {
			return
		}
		if s.fuente != nil {
			texto, err := s.fuente.TextoParaElMotor(t.numero)
			if err != nil {
				s.apuntarFalloDeTrabajo(t, EstadoError, ClaveInvalida)
				return
			}
			t.texto = texto
		}
	}
	// **El motor se resuelve aquí, en cada trabajo**: la dirección y el modelo salen de la
	// configuración de la instalación en este momento, así que cambiarlos en Configuración vale sin
	// reiniciar nada (docs/modules/ai.md).
	ajustes := s.ajustesDeAhora()

	for intento := 1; ; intento++ {
		pedido := encargoDe(t.tipo, t.texto, ajustes.Palabras, ajustes.Language)
		if intento > 1 {
			pedido = reforzarEncargo(pedido, ajustes.Language)
		}
		es, en, err := s.motor.pedir(s.ctx, ajustes, pedido)
		if err == nil && t.tipo == UltimaAccion && !ultimaAccionValida(es, t.texto) {
			err = ErrRespuestaInvalida
		}

		if err == nil {
			s.contarIntento(t)
			s.guardar(t, es, en, ajustes)
			return
		}

		// Un intento cortado por el apagado no es un intento del motor, y no se cuenta ni se apunta:
		// el campo se queda como estaba y la puesta al día del próximo arranque lo recogerá.
		if s.apagando() {
			return
		}

		s.contarIntento(t)

		if errors.Is(err, errIdiomaInvalido) {
			s.fallarORecuperar(t, EstadoError, ClaveInvalida)
			logs.LogWarning("el motor de IA contestó en otro idioma para el ticket " + t.numero + " (" + string(t.tipo) + ")")
			return
		}

		if errors.Is(err, ErrRespuestaInvalida) {
			if intento >= intentosRespuesta {
				if t.tipo == UltimaAccion {
					if respaldo, hay := respaldoDeUltimaAccion(t.texto, ajustes.Language, ajustes.Palabras); hay {
						s.guardar(t, respaldo, respaldo, ajustes)
						return
					}
				}
				if t.tipo == UltimaAccion && esUnTicketRecienAbierto(t.texto) {
					es, en := resumenDeApertura()
					if ajustes.Language == "en" {
						es = en
					}
					s.guardar(t, es, en, ajustes)
					return
				}
				s.fallarORecuperar(t, EstadoError, ClaveInvalida)
				// **Se registra la forma, nunca el contenido**: qué claves llegaron y cuánto miden. Sin
				// esto, «contestó algo que no vale» no se puede investigar; con esto, se sabe si faltó el
				// inglés, si vino vacío o si se cortó a mitad.
				logs.LogWarning("el motor de IA contestó algo que no vale para el ticket " +
					t.numero + " (" + string(t.tipo) + "): " + formaDe(err))
				return
			}

			// El reintento es inmediato: no es un motor lento, es un modelo que se ha explicado mal.
			continue
		}

		if intento >= intentosMotor {
			if t.tipo == UltimaAccion {
				if respaldo, hay := respaldoDeUltimaAccion(t.texto, ajustes.Language, ajustes.Palabras); hay {
					s.guardar(t, respaldo, respaldo, ajustes)
					return
				}
			}
			s.fallarORecuperar(t, EstadoSinMotor, ClaveSinMotor)
			logs.LogWarning("el motor de IA no contesta para el ticket " + t.numero + " (" + string(t.tipo) + ")")
			return
		}

		// La espera crece con cada intento: al motor que se ha quedado sin memoria hay que darle algo
		// más que al que ha tenido un hipo.
		if !s.esperar(s.esperaBase * time.Duration(intento)) {
			return
		}
	}
}

func marcasDelContexto(texto string) map[string]string {
	marcas := map[string]string{}
	for _, linea := range strings.Split(texto, "\n") {
		linea = strings.TrimSpace(linea)
		if !strings.HasPrefix(linea, "[") {
			continue
		}
		fin := strings.Index(linea, "]")
		if fin < 2 {
			continue
		}
		marcas[linea[1:fin]] = strings.TrimSpace(linea[fin+1:])
	}
	return marcas
}

func ultimaAccionValida(respuesta, contexto string) bool {
	m := marcasDelContexto(contexto)
	evento := strings.ToLower(m["TARGET_LAST_EVENT"])
	actor := strings.TrimSpace(m["TARGET_LAST_ACTOR"])
	if evento == "" {
		return true
	}
	texto := strings.ToLower(respuesta)
	if actor != "" && !strings.Contains(texto, strings.ToLower(actor)) {
		return false
	}
	palabras := map[string][]string{
		"cerrado": {"cerr", "clos"}, "resuelto": {"resolv", "solucion"},
		"reabierto": {"reab", "reopen"}, "asignado": {"asign", "assign"},
		"creado": {"abri", "cre", "open"}, "estado": {"estado", "state", "pasó", "moved"},
		"comentario": {"coment", "indic", "escrib", "said", "wrote", "reported"},
		"editado":    {"edit", "actualiz", "updat"}, "prioridad": {"prioridad", "priority"},
		"escalado": {"escal"}, "devuelto": {"devol", "return"},
		"observador_anadido": {"observ", "watch"}, "observador": {"observ", "watch"},
	}
	if esperadas := palabras[evento]; len(esperadas) > 0 {
		hay := false
		for _, palabra := range esperadas {
			if strings.Contains(texto, palabra) {
				hay = true
				break
			}
		}
		if !hay {
			return false
		}
	}
	razon := strings.ToLower(strings.TrimSpace(m["TARGET_REASON"]))
	return razon == "" || normalizarComparacion(texto) != normalizarComparacion(razon)
}

func normalizarComparacion(texto string) string {
	return strings.Join(strings.Fields(strings.Trim(strings.ToLower(texto), " .,:;!¡?¿")), " ")
}

func respaldoDeUltimaAccion(contexto, idioma string, palabras int) (string, bool) {
	m := marcasDelContexto(contexto)
	evento, actor := strings.ToLower(m["TARGET_LAST_EVENT"]), strings.TrimSpace(m["TARGET_LAST_ACTOR"])
	if evento == "" || actor == "" {
		return "", false
	}
	verbosES := map[string]string{
		"cerrado": "cerró el ticket", "resuelto": "resolvió el ticket", "reabierto": "reabrió el ticket",
		"asignado": "actualizó la asignación", "creado": "abrió el ticket", "estado": "cambió el estado del ticket",
		"comentario": "añadió un comentario", "editado": "actualizó el ticket", "prioridad": "cambió la prioridad del ticket",
		"escalado": "escaló el ticket", "devuelto": "devolvió el ticket", "observador_anadido": "añadió un observador",
		"observador": "quitó un observador",
	}
	verbosEN := map[string]string{
		"cerrado": "closed the ticket", "resuelto": "resolved the ticket", "reabierto": "reopened the ticket",
		"asignado": "updated the assignment", "creado": "opened the ticket", "estado": "changed the ticket status",
		"comentario": "added a comment", "editado": "updated the ticket", "prioridad": "changed the ticket priority",
		"escalado": "escalated the ticket", "devuelto": "returned the ticket", "observador_anadido": "added a watcher",
		"observador": "removed a watcher",
	}
	verbo := verbosES[evento]
	separador := ": "
	if idioma == "en" {
		verbo = verbosEN[evento]
	}
	if verbo == "" {
		return "", false
	}
	frase := actor + " " + verbo
	if comentario := strings.TrimSpace(m["TARGET_LAST_COMMENT"]); comentario != "" {
		frase += separador + comentario
	}
	return recortarPalabras(frase, palabras), true
}

func esUnTicketRecienAbierto(texto string) bool {
	lineas := strings.Split(texto, "\n")
	eventos := make([]string, 0, 2)
	for _, linea := range lineas {
		linea = strings.TrimSpace(linea)
		if strings.HasPrefix(linea, "-") {
			eventos = append(eventos, strings.ToLower(linea))
		}
	}
	if len(eventos) != 1 {
		return false
	}
	return strings.Contains(eventos[0], "(created)") || strings.Contains(eventos[0], "(creado)")
}

func resumenDeApertura() (string, string) {
	return "El ticket fue abierto y se describió el problema reportado.", "The ticket was opened and the reported problem was described."
}

// guardar escribe las dos redacciones en la fila del campo.
func (s *Service) guardar(t trabajo, es, en string, ajustes Ajustes) {
	var err error
	guardado := true
	if global, ok := s.store.(interface {
		GuardarResultadoGlobalSiVigente(string, string, int64, string, string, string, string, time.Time) (bool, error)
	}); ok && t.generacion > 0 && ajustes.Language != "" {
		guardado, err = global.GuardarResultadoGlobalSiVigente(t.numero, string(t.tipo), t.generacion, es, ajustes.Language, ajustes.Provider, ajustes.Modelo, s.now())
	} else if vigente, ok := s.store.(interface {
		GuardarResultadoSiVigente(string, string, int64, string, string, string, time.Time) (bool, error)
	}); ok && t.generacion > 0 {
		guardado, err = vigente.GuardarResultadoSiVigente(t.numero, string(t.tipo), t.generacion, es, en, ajustes.Modelo, s.now())
	} else if global, ok := s.store.(interface {
		GuardarResultadoGlobal(string, string, string, string, string, string, time.Time) error
	}); ok && ajustes.Language != "" {
		err = global.GuardarResultadoGlobal(t.numero, string(t.tipo), es, ajustes.Language, ajustes.Provider, ajustes.Modelo, s.now())
	} else {
		err = s.store.GuardarResultado(t.numero, string(t.tipo), es, en, ajustes.Modelo, s.now())
	}
	if err != nil {
		logs.LogError("no se pudo guardar el resumen del ticket " + t.numero + " (" + string(t.tipo) + "): " + err.Error())
		return
	}
	if !guardado {
		logs.LogInfo("se descartó un resumen obsoleto del ticket " + t.numero + " (" + string(t.tipo) + ")")
		return
	}

	// Se registran el número, el campo y el estado, y **nunca el texto del ticket** ni el que escribe
	// el motor: los logs son para saber qué pasó, no para leer los tickets de nadie.
	logs.LogSuccess("resumen escrito para el ticket " + t.numero + " (" + string(t.tipo) + ")")
}

// contarIntento suma uno al contador del campo. Es lo que corta el reintento infinito entre reinicios
// del backend (decisión 8).
// formaDe describe **sin el contenido** por qué no valió una respuesta: qué claves llegaron y
// cuánto midió cada una.
//
// Es lo que hace investigable el «contestó algo que no vale»: si falta el inglés, si llegó vacío o si
// el modelo se quedó a medias. El texto del ticket **no** se registra nunca.
func formaDe(err error) string {
	var deRespuesta *respuestaInvalida
	if errors.As(err, &deRespuesta) {
		return deRespuesta.forma
	}

	return "sin detalle"
}

func (s *Service) contarIntento(t trabajo) {
	var err error
	if vigente, ok := s.store.(interface {
		SubirIntentosSiVigente(string, string, int64) (bool, error)
	}); ok && t.generacion > 0 {
		_, err = vigente.SubirIntentosSiVigente(t.numero, string(t.tipo), t.generacion)
	} else {
		err = s.store.SubirIntentos(t.numero, string(t.tipo))
	}
	if err != nil {
		logs.LogError("no se pudo contar el intento del resumen del ticket " + t.numero + ": " + err.Error())
	}
}

func (s *Service) trabajoVigente(t trabajo) bool {
	if comprobable, ok := s.store.(interface {
		GeneracionVigente(string, string, int64) (bool, error)
	}); ok && t.generacion > 0 {
		vigente, err := comprobable.GeneracionVigente(t.numero, string(t.tipo), t.generacion)
		return err == nil && vigente
	}
	return true
}

func (s *Service) fallarORecuperar(t trabajo, estado Estado, clave string) {
	if !t.recuperacion && t.generacion > 0 && s.fuente != nil {
		if recuperable, ok := s.store.(interface {
			PrepararRecuperacion(string, string, int64, time.Time) (int64, bool, error)
		}); ok {
			noAntes := s.now().Add(s.esperaRecuperacion)
			nueva, preparada, err := recuperable.PrepararRecuperacion(t.numero, string(t.tipo), t.generacion, noAntes)
			if err != nil {
				logs.LogError("no se pudo programar la recuperación del resumen del ticket " + t.numero + ": " + err.Error())
			} else if preparada {
				s.programar(trabajo{numero: t.numero, tipo: t.tipo, generacion: nueva, recuperacion: true, noAntes: noAntes})
				return
			} else {
				return // Otra revisión ya reemplazó a ésta.
			}
		}
	}
	s.apuntarFalloDeTrabajo(t, estado, clave)
}

func (s *Service) apuntarFalloDeTrabajo(t trabajo, estado Estado, clave string) {
	if vigente, ok := s.store.(interface {
		MarcarFalloSiVigente(string, string, int64, string, string) (bool, error)
	}); ok && t.generacion > 0 {
		_, err := vigente.MarcarFalloSiVigente(t.numero, string(t.tipo), t.generacion, string(estado), clave)
		if err != nil {
			logs.LogError("no se pudo apuntar el resumen del ticket " + t.numero + " (" + string(t.tipo) + "): " + err.Error())
		}
		return
	}
	s.apuntarFallo(t.numero, t.tipo, estado, clave)
}

// esperar duerme entre intentos y **se despierta si el backend se está apagando**: apagar no puede
// quedarse esperando a que venza una espera.
func (s *Service) esperar(espera time.Duration) bool {
	temporizador := time.NewTimer(espera)
	defer temporizador.Stop()

	select {
	case <-temporizador.C:
		return true
	case <-s.parar:
		return false
	}
}

// apagando dice si el servicio está cerrando.
func (s *Service) apagando() bool {
	select {
	case <-s.parar:
		return true
	default:
		return false
	}
}
