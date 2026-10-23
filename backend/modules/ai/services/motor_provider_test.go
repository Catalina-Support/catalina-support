package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestLosCuatroAdaptadoresUsanSuContratoSinSalirALaRed(t *testing.T) {
	for _, provider := range []string{"openai", "deepseek", "openai-compatible", "claude"} {
		t.Run(provider, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if provider == "claude" {
					if r.URL.Path != "/v1/messages" || r.Header.Get("x-api-key") != "secreto" {
						t.Fatalf("contrato Claude incorrecto: %s, cabecera=%q", r.URL.Path, r.Header.Get("x-api-key"))
					}
					fmt.Fprint(w, `{"content":[{"text":"{\"text\":\"ok\"}"}]}`)
					return
				}
				if r.URL.Path != "/v1/chat/completions" || r.Header.Get("Authorization") != "Bearer secreto" {
					t.Fatalf("contrato compatible incorrecto: %s, cabecera=%q", r.URL.Path, r.Header.Get("Authorization"))
				}
				fmt.Fprint(w, `{"choices":[{"message":{"content":"{\"text\":\"ok\"}"}}]}`)
			}))
			defer server.Close()

			m := nuevoMotor(time.Second)
			if err := m.probarConfiguracion(context.Background(), server.URL, "modelo", provider, "bearer", "", "secreto", "es"); err != nil {
				t.Fatalf("el adaptador %s rechazó una respuesta válida: %v", provider, err)
			}
			text, _, err := m.pedirGlobal(context.Background(), Ajustes{
				URL: server.URL, Modelo: "modelo", Provider: provider, AuthType: "bearer",
				Credential: "secreto", Language: "es", Palabras: 60, Contexto: 4096,
			}, encargo{sistema: "resume", texto: "ticket falso"})
			if err != nil || text != "ok" {
				t.Fatalf("el resumen de %s no usó su adaptador: texto=%q err=%v", provider, text, err)
			}
		})
	}
}

func TestUnaRedireccionNoReenviaLaCredencial(t *testing.T) {
	var destinoLlamado atomic.Bool
	destino := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		destinoLlamado.Store(true)
	}))
	defer destino.Close()
	origen := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Location", destino.URL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	}))
	defer origen.Close()

	m := nuevoMotor(time.Second)
	err := m.probarConfiguracion(context.Background(), origen.URL, "modelo", "openai", "bearer", "", "no-debe-salir", "es")
	if !errors.Is(err, ErrSinMotor) || destinoLlamado.Load() {
		t.Fatalf("la redirección debía rechazarse sin llamar al destino: err=%v llamado=%v", err, destinoLlamado.Load())
	}
}

func TestUnaRespuestaMayorAlLimiteSeRechaza(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `{"choices":[{"message":{"content":"{\"text\":\"ok\"}"}}]}`+strings.Repeat(" ", maxRespuestaBytes))
	}))
	defer server.Close()

	m := nuevoMotor(time.Second)
	err := m.probarConfiguracion(context.Background(), server.URL, "modelo", "openai", "none", "", "", "es")
	if !errors.Is(err, ErrRespuestaInvalida) {
		t.Fatalf("se esperaba respuesta inválida por tamaño y llegó %v", err)
	}
}

func TestElIdiomaSeValidaConUnSoloReintentoLocal(t *testing.T) {
	var llamadas atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if llamadas.Add(1) == 1 {
			fmt.Fprint(w, `{"choices":[{"message":{"content":"{\"text\":\"The requester cannot access the application and needs help resetting the password.\"}"}}]}`)
			return
		}
		fmt.Fprint(w, `{"choices":[{"message":{"content":"{\"text\":\"La persona solicitante no puede entrar en la aplicación y necesita restablecer su contraseña.\"}"}}]}`)
	}))
	defer server.Close()

	m := nuevoMotor(time.Second)
	texto, _, err := m.pedirGlobal(context.Background(), Ajustes{URL: server.URL, Modelo: "modelo", Mode: "local", Language: "es", Palabras: 60, Contexto: 4096}, encargo{sistema: "resume", texto: "ticket"})
	if err != nil || !strings.HasPrefix(texto, "La persona") || llamadas.Load() != 2 {
		t.Fatalf("el reintento local no corrigió el idioma: texto=%q llamadas=%d err=%v", texto, llamadas.Load(), err)
	}
}

func TestElProveedorDePagoNoSeReintentaPorIdioma(t *testing.T) {
	var llamadas atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		llamadas.Add(1)
		fmt.Fprint(w, `{"choices":[{"message":{"content":"{\"text\":\"The requester cannot access the application and needs help resetting the password.\"}"}}]}`)
	}))
	defer server.Close()

	m := nuevoMotor(time.Second)
	_, _, err := m.pedirGlobal(context.Background(), Ajustes{URL: server.URL, Modelo: "modelo", Mode: "provider", Provider: "openai", Language: "es", Palabras: 60, Contexto: 4096}, encargo{sistema: "resume", texto: "ticket"})
	if !errors.Is(err, errIdiomaInvalido) || llamadas.Load() != 1 {
		t.Fatalf("el proveedor debía fallar sin reintento: llamadas=%d err=%v", llamadas.Load(), err)
	}
}

func TestElDetectorAceptaTextoBreveOTecnicoAmbiguo(t *testing.T) {
	for _, texto := range []string{"OK", "HTTP 500 SQL", "Qwen2.5 API"} {
		if !idiomaCompatible(texto, "es") || !idiomaCompatible(texto, "en") {
			t.Fatalf("el texto ambiguo %q no debe rechazarse", texto)
		}
	}
}

func TestTextoSignificativoRechazaMarcadoresSinContenido(t *testing.T) {
	for _, texto := range []string{"", "...", "—", "N/A", " / - _ "} {
		if textoSignificativo(texto) {
			t.Fatalf("%q no contiene un resumen significativo", texto)
		}
	}
	for _, texto := range []string{"HTTP 500", "El usuario abrió el ticket."} {
		if !textoSignificativo(texto) {
			t.Fatalf("%q sí contiene información", texto)
		}
	}
}

func TestUltimaAccionUsaEjemploDeAperturaYElReintentoSeRefuerza(t *testing.T) {
	primero := encargoDe(UltimaAccion, "Subject: Error\nConversation:\n- Ana (created)", 40, "en")
	if !strings.Contains(primero.usuario, "The requester opened the ticket") {
		t.Fatalf("el contrato no orienta el caso de apertura: %s", primero.usuario)
	}
	segundo := reforzarEncargo(primero, "en")
	if segundo.usuario == primero.usuario || !strings.Contains(segundo.usuario, "previous answer was invalid") {
		t.Fatalf("el reintento debe reforzar la instrucción: %s", segundo.usuario)
	}
}

func TestRedaccionProfesionalValidaEnEspanolEIngles(t *testing.T) {
	casos := []struct {
		nombre, borrador, propuesta string
	}{
		{
			nombre:    "español",
			borrador:  "reinicia la impresora y revisa si vuelve a fallar",
			propuesta: "He reiniciado la impresora y, por el momento, su funcionamiento parece haberse restablecido correctamente. Por favor, comprueba si ya puedes imprimir y avísanos si el problema vuelve a presentarse.",
		},
		{
			nombre:    "inglés",
			borrador:  "restart the printer and check if it fails again",
			propuesta: "I restarted the printer, and it now appears to be working correctly. Please try printing again and let us know if the issue returns so we can continue investigating.",
		},
	}
	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			if !redaccionProfesionalValida(caso.borrador, "Mañana informaremos el resultado después de revisar el servidor.", caso.propuesta) {
				t.Fatalf("se rechazó una propuesta profesional: %q", caso.propuesta)
			}
		})
	}
}

func TestMejorarRedaccionUsaContextoMinimoYDevuelveTextoPlano(t *testing.T) {
	var recibido string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		recibido = string(data)
		fmt.Fprint(w, `{"choices":[{"message":{"content":"{\"text\":\"Hemos comprobado que el acceso se encuentra disponible nuevamente para tu cuenta. Por favor, intenta ingresar y comunícanos si vuelves a encontrar alguna dificultad.\"}"}}]}`)
	}))
	defer server.Close()

	servicio := nuevoServicio(nil, Ajustes{
		URL: server.URL, Modelo: "modelo", Provider: "openai-compatible", Language: "es",
		Palabras: 40, Contexto: 4096, Espera: time.Second,
	})
	defer servicio.Parar()

	resultado, err := servicio.MejorarRedaccion(
		"hola ana ya puedes ingresar", "", "Ana Pérez", "main", "es",
	)
	if err != nil || !strings.HasPrefix(resultado.Texto, "Hemos comprobado") {
		t.Fatalf("resultado inesperado: texto=%q err=%v", resultado.Texto, err)
	}
	for _, esperado := range []string{"hola ana ya puedes ingresar", "Ana Pérez", "2 o 3 oraciones", "por favor", "Respuesta de ejemplo", "ticket principal de soporte"} {
		if !strings.Contains(recibido, esperado) {
			t.Fatalf("la petición no contiene %q: %s", esperado, recibido)
		}
	}
	for _, prohibido := range []string{"attachment", "conversation", "subject"} {
		if strings.Contains(strings.ToLower(recibido), prohibido) {
			t.Fatalf("la petición incluyó contexto no autorizado %q", prohibido)
		}
	}
}

func TestMejorarRedaccionRegeneraDesdeElBorradorYEvitandoSoloLaUltimaPropuesta(t *testing.T) {
	candidata := "Revisaremos el servidor de manera exhaustiva para comprobar su funcionamiento e identificar el origen del comportamiento informado. Mañana le comunicaremos el resultado obtenido y los siguientes pasos que correspondan."
	if !redaccionProfesionalValida("revisamos el servidor y mañana comunicamos el resultado", "Mañana informaremos el resultado después de revisar el servidor.", candidata) {
		t.Fatal("el ejemplo de regeneración debe superar la validación local")
	}
	var recibido string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		recibido = string(data)
		fmt.Fprint(w, `{"choices":[{"message":{"content":"{\"text\":\"Revisaremos el servidor de manera exhaustiva para comprobar su funcionamiento e identificar el origen del comportamiento informado. Mañana le comunicaremos el resultado obtenido y los siguientes pasos que correspondan.\"}"}}]}`)
	}))
	defer server.Close()

	servicio := nuevoServicio(nil, Ajustes{
		URL: server.URL, Modelo: "modelo", Provider: "openai-compatible", Language: "es",
		Palabras: 40, Contexto: 4096, Espera: time.Second,
	})
	defer servicio.Parar()

	resultado, err := servicio.MejorarRedaccion(
		"revisamos el servidor y mañana comunicamos el resultado",
		"Mañana informaremos el resultado después de revisar el servidor.",
		"Usuario Cuatro", "main", "es",
	)
	if err != nil || !strings.HasPrefix(resultado.Texto, "Revisaremos") {
		t.Fatalf("resultado inesperado: texto=%q err=%v", resultado.Texto, err)
	}
	for _, esperado := range []string{
		"revisamos el servidor y mañana comunicamos el resultado",
		"Mañana informaremos el resultado después de revisar el servidor.",
		"fuente de verdad",
		"paráfrasis claramente distinta",
	} {
		if !strings.Contains(recibido, esperado) {
			t.Fatalf("la regeneración no contiene %q: %s", esperado, recibido)
		}
	}
}

func TestMejorarRedaccionReintentaYRechazaSiLaSegundaSigueSiendoPobre(t *testing.T) {
	var llamadas atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		llamadas.Add(1)
		fmt.Fprint(w, `{"choices":[{"message":{"content":"{\"text\":\"hola necesito datos\"}"}}]}`)
	}))
	defer server.Close()

	servicio := nuevoServicio(nil, Ajustes{
		URL: server.URL, Modelo: "modelo", Provider: "openai-compatible", Language: "es",
		Palabras: 40, Contexto: 4096, Espera: time.Second,
	})
	defer servicio.Parar()

	resultado, err := servicio.MejorarRedaccion("hola necesito datos", "", "Ana", "main", "es")
	if !errors.Is(err, ErrRespuestaInvalida) {
		t.Fatalf("se esperaba rechazo de calidad, resultado=%+v err=%v", resultado, err)
	}
	if llamadas.Load() != 2 {
		t.Fatalf("se esperaban dos intentos: llamadas=%d", llamadas.Load())
	}
}

func TestConfiguradoNoCompruebaLaSaludDelMotor(t *testing.T) {
	servicio := nuevoServicio(nil, Ajustes{URL: "http://127.0.0.1:1", Modelo: "modelo"})
	defer servicio.Parar()
	if !servicio.Configurado() {
		t.Fatal("una dirección configurada debe conservar la capacidad aunque el motor esté caído")
	}
}
