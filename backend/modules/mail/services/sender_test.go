package services

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"testing"
)

// servidorSMTPDePrueba es un servidor de correo mínimo, **en memoria y en la propia máquina**: existe
// para comprobar que la prueba de la conexión habla el protocolo y no manda nada. No sale a la red.
type servidorSMTPDePrueba struct {
	ln          net.Listener
	rechazaAuth bool

	mu       sync.Mutex
	comandos []string
}

func nuevoServidorSMTPDePrueba(t *testing.T, rechazaAuth bool) *servidorSMTPDePrueba {
	t.Helper()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("no se pudo abrir el servidor de prueba: %v", err)
	}

	servidor := &servidorSMTPDePrueba{ln: ln, rechazaAuth: rechazaAuth}
	go servidor.atender()

	t.Cleanup(func() { ln.Close() })

	return servidor
}

func (s *servidorSMTPDePrueba) puerto() string {
	_, puerto, _ := net.SplitHostPort(s.ln.Addr().String())
	return puerto
}

func (s *servidorSMTPDePrueba) apuntar(linea string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.comandos = append(s.comandos, linea)
}

// apuntado devuelve todo lo que el cliente ha dicho, para poder afirmar qué hizo y qué no.
func (s *servidorSMTPDePrueba) apuntado() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return strings.Join(s.comandos, "\n")
}

func (s *servidorSMTPDePrueba) atender() {
	for {
		conexion, err := s.ln.Accept()
		if err != nil {
			return
		}

		go s.atenderConexion(conexion)
	}
}

func (s *servidorSMTPDePrueba) atenderConexion(conexion net.Conn) {
	defer conexion.Close()

	lector := bufio.NewReader(conexion)
	fmt.Fprint(conexion, "220 prueba\r\n")

	esperaContrasena := false
	enDatos := false
	for {
		linea, err := lector.ReadString('\n')
		if err != nil {
			return
		}

		comando := strings.TrimRight(linea, "\r\n")
		s.apuntar(comando)
		if enDatos {
			if comando == "." {
				enDatos = false
				fmt.Fprint(conexion, "250 2.0.0 queued\r\n")
			}
			continue
		}
		if esperaContrasena {
			if s.rechazaAuth {
				fmt.Fprint(conexion, "535 5.7.8 Authentication credentials invalid\r\n")
			} else {
				fmt.Fprint(conexion, "235 2.7.0 Authentication successful\r\n")
			}
			esperaContrasena = false
			continue
		}

		switch {
		case strings.HasPrefix(strings.ToUpper(comando), "EHLO"):
			// **Sin STARTTLS**: la prueba va en claro contra la propia máquina y con AUTH anunciado.
			fmt.Fprint(conexion, "250-prueba\r\n250-AUTH LOGIN PLAIN\r\n250 OK\r\n")
		case strings.HasPrefix(strings.ToUpper(comando), "HELO"):
			fmt.Fprint(conexion, "250 prueba\r\n")
		case strings.HasPrefix(strings.ToUpper(comando), "AUTH"):
			if strings.HasPrefix(strings.ToUpper(comando), "AUTH LOGIN") {
				esperaContrasena = true
				fmt.Fprint(conexion, "334 UGFzc3dvcmQ6\r\n")
				continue
			}
			if s.rechazaAuth {
				fmt.Fprint(conexion, "535 5.7.8 Authentication credentials invalid\r\n")
				continue
			}
			fmt.Fprint(conexion, "235 2.7.0 Authentication successful\r\n")
		case strings.HasPrefix(strings.ToUpper(comando), "QUIT"):
			fmt.Fprint(conexion, "221 Bye\r\n")
			return
		case strings.EqualFold(comando, "DATA"):
			enDatos = true
			fmt.Fprint(conexion, "354 End data with <CR><LF>.<CR><LF>\r\n")
		default:
			fmt.Fprint(conexion, "250 OK\r\n")
		}
	}
}

func TestProbarEnvioEntregaRemitenteDestinatarioYCuerpo(t *testing.T) {
	servidor := nuevoServidorSMTPDePrueba(t, false)
	servicio := NewService(nil, NewSender(), "")
	err := servicio.ProbarEnvio(CorreoSaliente{
		Host: "127.0.0.1", Port: servidor.puerto(), SecurityMode: "plain", InsecureAccepted: true,
		FromName: "AMC Soporte", FromEmail: "soporte@empresa.com",
	}, "destino@empresa.com", "es", "Catalina AMC")
	if err != nil {
		t.Fatalf("el envío de prueba no debería fallar: %v", err)
	}
	dicho := servidor.apuntado()
	for _, esperado := range []string{"MAIL FROM:<soporte@empresa.com>", "RCPT TO:<destino@empresa.com>", "DATA", "Subject: [prueba] Correo de prueba de Catalina AMC", "No es necesario responder"} {
		if !strings.Contains(dicho, esperado) {
			t.Fatalf("falta %q en:\n%s", esperado, dicho)
		}
	}
}

// Sin servidor no hay nada que probar.
func TestProbarSinConfiguracion(t *testing.T) {
	remitente := NewSender()

	if err := remitente.Probar(CorreoSaliente{}); !errors.Is(err, ErrNotConfigured) {
		t.Fatalf("se esperaba %v y llegó %v", ErrNotConfigured, err)
	}
}

// La prueba **conecta y autentica, pero no manda ningún correo**: no hay MAIL FROM, ni RCPT TO, ni
// DATA (docs/primer-arranque.md, sección 3).
func TestProbarConectaYAutenticaSinEnviar(t *testing.T) {
	servidor := nuevoServidorSMTPDePrueba(t, false)
	remitente := NewSender()

	err := remitente.Probar(CorreoSaliente{
		Host:             "127.0.0.1",
		Port:             servidor.puerto(),
		SecurityMode:     "plain",
		InsecureAccepted: true,
		User:             "mesa",
		Password:         "secreta",
	})
	if err != nil {
		t.Fatalf("la prueba no debería fallar: %v", err)
	}

	dicho := servidor.apuntado()
	for _, esperado := range []string{"EHLO", "AUTH"} {
		if !strings.Contains(dicho, esperado) {
			t.Fatalf("la prueba debería haber dicho %s:\n%s", esperado, dicho)
		}
	}
	for _, prohibido := range []string{"MAIL FROM", "RCPT TO", "DATA"} {
		if strings.Contains(dicho, prohibido) {
			t.Fatalf("la prueba no puede mandar %s:\n%s", prohibido, dicho)
		}
	}
	// Y la contraseña no viaja en claro por el protocolo.
	if strings.Contains(dicho, "secreta") {
		t.Fatal("la contraseña no puede salir en claro")
	}
}

// Sin usuario no se autentica: se comprueba sólo que el servidor contesta.
func TestProbarSinCredenciales(t *testing.T) {
	servidor := nuevoServidorSMTPDePrueba(t, false)
	remitente := NewSender()

	if err := remitente.Probar(CorreoSaliente{Host: "127.0.0.1", Port: servidor.puerto(), SecurityMode: "plain", InsecureAccepted: true}); err != nil {
		t.Fatalf("la prueba no debería fallar: %v", err)
	}

	if strings.Contains(servidor.apuntado(), "AUTH") {
		t.Fatal("sin usuario no hay nada que autenticar")
	}
}

// Si el servidor rechaza las credenciales, la prueba falla con la clave de autenticación.
func TestProbarConCredencialesRechazadas(t *testing.T) {
	servidor := nuevoServidorSMTPDePrueba(t, true)
	remitente := NewSender()

	err := remitente.Probar(CorreoSaliente{
		Host:             "127.0.0.1",
		Port:             servidor.puerto(),
		SecurityMode:     "plain",
		InsecureAccepted: true,
		User:             "mesa",
		Password:         "mala",
	})
	if err == nil {
		t.Fatal("con las credenciales rechazadas la prueba tiene que fallar")
	}
	if !strings.Contains(err.Error(), "mail.smtp.auth") {
		t.Fatalf("el fallo debería llevar la clave de autenticación: %v", err)
	}
}

func TestPlainExigeAceptacion(t *testing.T) {
	servidor := nuevoServidorSMTPDePrueba(t, false)
	err := NewSender().Probar(CorreoSaliente{Host: "127.0.0.1", Port: servidor.puerto(), SecurityMode: "plain"})
	if err == nil || !strings.Contains(err.Error(), "not accepted") {
		t.Fatalf("se esperaba rechazo sin aceptación y llegó %v", err)
	}
}

func TestStartTLSNoCaeATextoPlano(t *testing.T) {
	servidor := nuevoServidorSMTPDePrueba(t, false)
	err := NewSender().Probar(CorreoSaliente{Host: "127.0.0.1", Port: servidor.puerto(), SecurityMode: "starttls"})
	if err == nil || !strings.Contains(err.Error(), "STARTTLS not supported") {
		t.Fatalf("se esperaba rechazo sin STARTTLS y llegó %v", err)
	}
	if strings.Contains(servidor.apuntado(), "AUTH") {
		t.Fatal("STARTTLS ausente no puede degradarse ni autenticar")
	}
}

func TestPlainRechazaResolucionMixta(t *testing.T) {
	remitente := NewSender()
	remitente.lookupIP = func(string) ([]net.IP, error) {
		return []net.IP{net.ParseIP("172.20.0.10"), net.ParseIP("203.0.113.10")}, nil
	}
	err := remitente.Probar(CorreoSaliente{Host: "relay.example", Port: "587", SecurityMode: "plain", InsecureAccepted: true})
	if err == nil || !strings.Contains(err.Error(), "private IPv4") {
		t.Fatalf("se esperaba rechazo de DNS mixto y llegó %v", err)
	}
}

func TestPlainRechazaIPv6(t *testing.T) {
	_, err := direccionPrivada("relay.example", "587", func(string) ([]net.IP, error) {
		return []net.IP{net.ParseIP("fd00::1")}, nil
	})
	if err == nil || !strings.Contains(err.Error(), "private IPv4") {
		t.Fatalf("se esperaba rechazo de IPv6 y llegó %v", err)
	}
}
