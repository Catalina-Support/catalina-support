package services

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"
)

// ErrNotConfigured se devuelve cuando falta la configuración de SMTP. Es un error propio y no un
// fallo genérico: quien lo lea tiene que saber que el problema no es el correo, sino la
// configuración.
var ErrNotConfigured = errors.New("mail.smtp.notConfigured")

// Sender envía correos por SMTP, con la biblioteca estándar y sin dependencias.
type Sender struct {
	host      string
	port      string
	user      string
	password  string
	fromName  string
	fromEmail string
	timeout   time.Duration
	lookupIP  func(string) ([]net.IP, error)
	// instalacion da el correo saliente: **vive en la configuración**, como el directorio y Keycloak.
	// Ya no hay respaldo en el entorno (decisión del responsable, 2026-09-30).
	instalacion Instalacion
}

// NewSender construye el remitente a partir de la configuración del entorno.
func NewSender() *Sender {
	return &Sender{timeout: 15 * time.Second, lookupIP: net.LookupIP}
}

// SetInstalacion fija de dónde sale el correo saliente cuando la instalación tiene uno puesto. Lo
// llama el servicio, y el respaldo es lo del entorno.
func (s *Sender) SetInstalacion(instalacion Instalacion) { s.instalacion = instalacion }

// datos es el correo saliente de este envío: **el de la configuración**, que es donde vive. Sin él no
// hay con qué enviar, y el envío falla con su error claro en lugar de intentar una conexión a ninguna
// parte (docs/primer-arranque.md, sección 5).
func (s *Sender) datos() CorreoSaliente {
	if s.instalacion == nil {
		return CorreoSaliente{}
	}

	correo, _ := s.instalacion.SMTP()

	return correo
}

// Configured dice si hay con qué enviar. Sin esto, el envío falla con un error claro en lugar de
// intentar una conexión a ninguna parte.
func (s *Sender) Configured() bool {
	correo := s.datos()

	return correo.Host != "" && correo.Port != "" && correo.FromEmail != ""
}

// Send arma el mensaje y lo entrega.
//
// El tiempo límite es del código: un servidor de correo que no responde no puede dejar colgada una
// petición ni una tarea de fondo para siempre.
func (s *Sender) Send(to []string, subject, htmlBody string) error {
	if !s.Configured() {
		return ErrNotConfigured
	}
	if len(to) == 0 {
		return errors.New("mail.to.empty")
	}

	correo := s.datos()

	message, err := Build(correo.FromName, correo.FromEmail, to, subject, htmlBody)
	if err != nil {
		return err
	}

	return s.send(to, message, correo)
}

// Probar abre la conexión y **valida la autenticación, sin mandar ningún correo**: no hay MAIL FROM,
// ni RCPT TO, ni DATA. Es lo que necesita el asistente, donde no hay destinatario y sólo se quiere
// saber si el servidor contesta y si las credenciales valen (docs/primer-arranque.md, sección 3).
//
// Usa **la misma conexión y la misma autenticación que el envío** —`conectar`—, así que lo que se
// prueba es exactamente lo que se va a usar al mandar.
func (s *Sender) Probar(correo CorreoSaliente) error {
	if correo.Host == "" || correo.Port == "" {
		return ErrNotConfigured
	}

	cliente, err := s.conectar(correo)
	if err != nil {
		return err
	}
	defer cliente.Close()

	return cliente.Quit()
}

// SendWithConfig entrega un mensaje con una configuración todavía no guardada. Es el camino de la
// prueba real del asistente y de Configuración; no modifica la configuración activa del Sender.
func (s *Sender) SendWithConfig(to []string, subject, htmlBody string, correo CorreoSaliente) error {
	if correo.Host == "" || correo.Port == "" || correo.FromEmail == "" {
		return ErrNotConfigured
	}
	if len(to) == 0 {
		return errors.New("mail.to.empty")
	}
	message, err := Build(correo.FromName, correo.FromEmail, to, subject, htmlBody)
	if err != nil {
		return err
	}
	return s.send(to, message, correo)
}

// conectar deja un cliente SMTP **ya conectado y autenticado**, listo para entregar un mensaje o para
// comprobar la conexión. El cifrado directo (puerto 465) y STARTTLS (587) se resuelven aquí, con el
// mismo tiempo límite: un servidor que no responde no puede dejar colgada una petición.
func (s *Sender) conectar(correo CorreoSaliente) (*smtp.Client, error) {
	espera := s.timeout
	if espera == 0 {
		espera = 15 * time.Second
	}

	var cliente *smtp.Client

	if correo.SecurityMode == "tls_direct" {
		address := net.JoinHostPort(correo.Host, correo.Port)
		conexion, err := tls.DialWithDialer(
			&net.Dialer{Timeout: espera},
			"tcp",
			address,
			&tls.Config{ServerName: correo.Host, MinVersion: tls.VersionTLS12},
		)
		if err != nil {
			return nil, fmt.Errorf("mail.smtp.connect: %w", err)
		}

		cliente, err = smtp.NewClient(conexion, correo.Host)
		if err != nil {
			conexion.Close()
			return nil, fmt.Errorf("mail.smtp.handshake: %w", err)
		}
	} else {
		address := net.JoinHostPort(correo.Host, correo.Port)
		if correo.SecurityMode == "plain" {
			if !correo.InsecureAccepted {
				return nil, fmt.Errorf("mail.smtp.connect: insecure mode not accepted")
			}
			var err error
			address, err = direccionPrivada(correo.Host, correo.Port, s.lookupIP)
			if err != nil {
				return nil, fmt.Errorf("mail.smtp.connect: %w", err)
			}
		}
		conexion, err := net.DialTimeout("tcp", address, espera)
		if err != nil {
			return nil, fmt.Errorf("mail.smtp.connect: %w", err)
		}

		cliente, err = smtp.NewClient(conexion, correo.Host)
		if err != nil {
			conexion.Close()
			return nil, fmt.Errorf("mail.smtp.handshake: %w", err)
		}

		if correo.SecurityMode == "starttls" {
			soportado, _ := cliente.Extension("STARTTLS")
			if !soportado {
				cliente.Close()
				return nil, fmt.Errorf("mail.smtp.handshake: STARTTLS not supported")
			}
			if err := cliente.StartTLS(&tls.Config{ServerName: correo.Host, MinVersion: tls.VersionTLS12}); err != nil {
				cliente.Close()
				return nil, fmt.Errorf("mail.smtp.handshake: %w", err)
			}
		}
	}

	if correo.User != "" {
		auth := smtp.PlainAuth("", correo.User, correo.Password, correo.Host)
		if correo.SecurityMode == "plain" {
			auth = loginAuth{user: correo.User, password: correo.Password}
		}
		if err := cliente.Auth(auth); err != nil {
			cliente.Close()
			return nil, fmt.Errorf("mail.smtp.auth: %w", err)
		}
	}

	return cliente, nil
}

// loginAuth implementa el mecanismo LOGIN que anuncia el relay interno heredado. Sólo se selecciona
// después de validar el modo `plain`, su aceptación y la dirección privada; los modos cifrados
// conservan PLAIN dentro de TLS.
type loginAuth struct {
	user     string
	password string
}

func (a loginAuth) Start(*smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.user), nil
}

func (a loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if !more {
		return nil, nil
	}
	prompt := strings.ToLower(string(fromServer))
	if strings.Contains(prompt, "password") || strings.Contains(prompt, "pass") {
		return []byte(a.password), nil
	}
	return nil, errors.New("unexpected SMTP AUTH LOGIN challenge")
}

func (s *Sender) send(to []string, message []byte, correo CorreoSaliente) error {
	cliente, err := s.conectar(correo)
	if err != nil {
		return err
	}
	defer cliente.Close()

	if err := cliente.Mail(correo.FromEmail); err != nil {
		return fmt.Errorf("mail.smtp.from: %w", err)
	}
	for _, recipient := range to {
		if err := cliente.Rcpt(recipient); err != nil {
			return fmt.Errorf("mail.smtp.recipient: %w", err)
		}
	}

	writer, err := cliente.Data()
	if err != nil {
		return fmt.Errorf("mail.smtp.data: %w", err)
	}
	if _, err := writer.Write(message); err != nil {
		return fmt.Errorf("mail.smtp.write: %w", err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("mail.smtp.close: %w", err)
	}

	return cliente.Quit()
}

// direccionPrivada resuelve una sola vez y devuelve una de las direcciones ya comprobadas. Todas
// deben ser IPv4 privadas o loopback: una respuesta DNS mixta no puede sacar credenciales de la red.
func direccionPrivada(host, port string, lookup func(string) ([]net.IP, error)) (string, error) {
	direcciones, err := lookup(host)
	if err != nil {
		return "", err
	}
	if len(direcciones) == 0 {
		return "", errors.New("host without addresses")
	}
	var elegida net.IP
	for _, direccion := range direcciones {
		ipv4 := direccion.To4()
		if ipv4 == nil || (!ipv4.IsPrivate() && !ipv4.IsLoopback()) {
			return "", errors.New("unencrypted SMTP requires a private IPv4 destination")
		}
		if elegida == nil {
			elegida = ipv4
		}
	}
	return net.JoinHostPort(elegida.String(), port), nil
}
