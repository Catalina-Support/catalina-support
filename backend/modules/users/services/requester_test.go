package services

import (
	"errors"
	"strconv"
	"testing"
	"time"

	"catalina-support/backend/modules/users/repositories"
	"catalina-support/backend/shared/auth"
	"catalina-support/backend/shared/config"
	"catalina-support/backend/shared/database"
)

type accessDePrueba string

func (a accessDePrueba) Access() (config.Access, error) {
	return config.Access{Method: string(a)}, nil
}

type directorioDePrueba struct {
	personas []auth.DirectoryAccount
	activo   bool
}

func (d directorioDePrueba) Knows(string) (bool, error) { return false, nil }

func (d directorioDePrueba) SearchDirectory(string, bool, int) ([]auth.DirectoryAccount, error) {
	return d.personas, nil
}

func (d directorioDePrueba) DirectoryEnabled() bool { return d.activo }

func servicioDeSolicitantes(t *testing.T) (*Service, func()) {
	t.Helper()

	cfg, err := config.Load()
	if err != nil {
		t.Skip("sin configuración de base de datos: " + err.Error())
	}
	db, err := database.Connect(cfg)
	if err != nil {
		t.Skip("la base de datos no responde: " + err.Error())
	}
	if !db.Migrator().HasTable(&repositories.User{}) {
		t.Skip("falta la tabla users: aplica las migraciones")
	}

	servicio := NewService(repositories.NewUserRepository(db))

	return servicio, func() {
		if sqlDB, err := db.DB(); err == nil {
			_ = sqlDB.Close()
		}
	}
}

func TestEnsureRequesterCreaCuentaLocalSinAcceso(t *testing.T) {
	servicio, cerrar := servicioDeSolicitantes(t)
	defer cerrar()
	servicio.SetAccess(accessDePrueba(auth.OriginLocal))

	sufijo := strconv.FormatInt(time.Now().UnixNano(), 36)
	correo := "solicitante-" + sufijo + "@example.invalid"

	cuenta, creada, err := servicio.EnsureRequester(correo, "Ana", "Pérez", auth.Identity{ID: 1, Role: auth.RoleSoporte})
	if err != nil {
		t.Fatalf("el alta local falló: %v", err)
	}
	if !creada {
		t.Fatal("la cuenta tenía que marcarse como creada")
	}
	if cuenta.Origin != auth.OriginLocal || cuenta.Role != auth.RoleUsuario || !cuenta.IsActive {
		t.Fatalf("la cuenta local no quedó como debía: %+v", cuenta)
	}

	usuario, err := servicio.repo.FindByEmail(correo)
	if err != nil {
		t.Fatalf("la cuenta creada no se puede leer: %v", err)
	}
	if usuario.PasswordHash != nil {
		t.Fatal("la cuenta de solicitante no puede tener contraseña")
	}

	_, creadaOtraVez, err := servicio.EnsureRequester(correo, "", "", auth.Identity{ID: 1, Role: auth.RoleSoporte})
	if err != nil || creadaOtraVez {
		t.Fatalf("una cuenta existente se reutiliza sin crearla otra vez: creada=%v err=%v", creadaOtraVez, err)
	}
}

func TestEnsureRequesterExigeNombresParaLaCuentaLocal(t *testing.T) {
	servicio, cerrar := servicioDeSolicitantes(t)
	defer cerrar()
	servicio.SetAccess(accessDePrueba(auth.OriginLocal))

	sufijo := strconv.FormatInt(time.Now().UnixNano(), 36)
	_, _, err := servicio.EnsureRequester(
		"sinnombre-"+sufijo+"@example.invalid",
		"",
		"",
		auth.Identity{ID: 1, Role: auth.RoleSoporte},
	)
	if !errors.Is(err, auth.ErrRequesterNameRequired) {
		t.Fatalf("se esperaba ErrRequesterNameRequired y llegó %v", err)
	}
}

func TestEnsureRequesterRechazaKeycloakSinCuenta(t *testing.T) {
	servicio, cerrar := servicioDeSolicitantes(t)
	defer cerrar()
	servicio.SetAccess(accessDePrueba(auth.OriginKeycloak))

	sufijo := strconv.FormatInt(time.Now().UnixNano(), 36)
	_, _, err := servicio.EnsureRequester(
		"keycloak-"+sufijo+"@example.invalid",
		"Ana",
		"Pérez",
		auth.Identity{ID: 1, Role: auth.RoleSoporte},
	)
	if !errors.Is(err, auth.ErrRequesterAccountExists) {
		t.Fatalf("se esperaba ErrRequesterAccountExists y llegó %v", err)
	}
}

func TestEnsureRequesterRechazaADSiNoEstáEnElDirectorio(t *testing.T) {
	servicio, cerrar := servicioDeSolicitantes(t)
	defer cerrar()
	servicio.SetAccess(accessDePrueba(auth.OriginAD))
	servicio.SetDirectory(directorioDePrueba{activo: true})

	sufijo := strconv.FormatInt(time.Now().UnixNano(), 36)
	_, _, err := servicio.EnsureRequester(
		"ad-"+sufijo+"@example.invalid",
		"Ana",
		"Pérez",
		auth.Identity{ID: 1, Role: auth.RoleSoporte},
	)
	if !errors.Is(err, auth.ErrRequesterNotInDirectory) {
		t.Fatalf("se esperaba ErrRequesterNotInDirectory y llegó %v", err)
	}
}

func TestEnsureRequesterCreaDesdeADConLosDatosDelDirectorio(t *testing.T) {
	servicio, cerrar := servicioDeSolicitantes(t)
	defer cerrar()
	servicio.SetAccess(accessDePrueba(auth.OriginAD))

	sufijo := strconv.FormatInt(time.Now().UnixNano(), 36)
	correo := "ad-alta-" + sufijo + "@example.invalid"
	servicio.SetDirectory(directorioDePrueba{
		activo: true,
		personas: []auth.DirectoryAccount{{
			Origin:     auth.OriginAD,
			ExternalID: "ext-" + sufijo,
			Username:   "ad-" + sufijo,
			Email:      correo,
			Name:       "Ana",
			LastName:   "Pérez",
		}},
	})

	cuenta, creada, err := servicio.EnsureRequester(correo, "", "", auth.FactoryIdentity())
	if err != nil {
		t.Fatalf("el alta desde AD falló: %v", err)
	}
	if !creada || cuenta.Origin != auth.OriginAD || cuenta.Role != auth.RoleUsuario || cuenta.HasPassword() {
		t.Fatalf("la cuenta de AD no quedó como debía: creada=%v cuenta=%+v", creada, cuenta)
	}
	if cuenta.Name != "Ana" || cuenta.LastName != "Pérez" {
		t.Fatalf("los datos del directorio no mandaron: %+v", cuenta)
	}
}
