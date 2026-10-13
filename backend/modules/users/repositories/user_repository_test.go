package repositories

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"catalina-support/backend/shared/config"
	"catalina-support/backend/shared/database"
)

type loggerDePrueba struct {
	errores []error
}

func (l *loggerDePrueba) LogMode(logger.LogLevel) logger.Interface { return l }
func (l *loggerDePrueba) Info(context.Context, string, ...any)     {}
func (l *loggerDePrueba) Warn(context.Context, string, ...any)     {}
func (l *loggerDePrueba) Error(_ context.Context, _ string, args ...any) {
	for _, arg := range args {
		if err, ok := arg.(error); ok {
			l.errores = append(l.errores, err)
		}
	}
}
func (l *loggerDePrueba) Trace(_ context.Context, _ time.Time, _ func() (string, int64), err error) {
	if err != nil {
		l.errores = append(l.errores, err)
	}
}

func TestBusquedaDeIdentidadAusenteNoSeRegistraComoError(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Skip("sin configuración de base de datos: " + err.Error())
	}

	db, err := database.Connect(cfg)
	if err != nil {
		t.Skip("la base de datos no responde: " + err.Error())
	}
	if sqlDB, err := db.DB(); err == nil {
		t.Cleanup(func() { _ = sqlDB.Close() })
	}
	if !db.Migrator().HasTable(&User{}) {
		t.Skip("falta la tabla users: aplica backend/migrations/v1.0.0.sql")
	}

	registro := &loggerDePrueba{}
	db = db.Session(&gorm.Session{Logger: registro})
	repositorio := NewUserRepository(db)
	sufijo := strconv.FormatInt(time.Now().UnixNano(), 36)

	if _, err := repositorio.FindByEmail("ausente-" + sufijo + "@example.invalid"); !errors.Is(err, ErrUserNotFound) {
		t.Fatalf("el correo ausente debería devolver %v y devolvió %v", ErrUserNotFound, err)
	}
	if _, err := repositorio.FindByExternalID("ad", "ausente-"+sufijo); !errors.Is(err, ErrUserNotFound) {
		t.Fatalf("la identidad ausente debería devolver %v y devolvió %v", ErrUserNotFound, err)
	}
	if len(registro.errores) != 0 {
		t.Fatalf("las ausencias esperadas se registraron como error: %v", registro.errores)
	}

	if err := db.Table("tabla_que_no_existe_para_prueba").Find(&[]struct{}{}).Error; err == nil {
		t.Fatal("la consulta inválida tenía que fallar")
	}
	if len(registro.errores) == 0 {
		t.Fatal("un error SQL real no llegó al logger")
	}
}
