// Package repositories es el acceso a datos del módulo ai: la tabla de los dos resúmenes que redacta
// el motor de IA (docs/modules/ai.md, sección 4).
//
// **Sólo se escribe en `ai_insights`**: este módulo no escribe en las tablas de nadie más y nadie
// escribe en la suya (docs/modules/ai.md, sección 5).
package repositories

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Los dos campos que se le piden al motor. Son un valor cerrado y viven aquí para que la base y el
// código digan lo mismo, como los estados de un ticket.
const (
	KindSummary      = "summary"
	KindUltimaAccion = "ultima_accion"
)

const (
	SummaryContextVersion    = 2
	LastActionContextVersion = 2
)

func contextVersion(kind string) int {
	if kind == KindSummary {
		return SummaryContextVersion
	}
	return LastActionContextVersion
}

// Los cuatro estados de un campo.
//
// **Los dos fallos están separados a propósito**: `error` es una respuesta que no valía —culpa del
// modelo— y `sin_motor` es que hay motor configurado pero no contesta. No significan lo mismo ni se
// arreglan igual, así que no pueden compartir estado (docs/modules/ai.md, decisión 8).
const (
	StatePendiente = "pendiente"
	StateListo     = "listo"
	StateError     = "error"
	StateSinMotor  = "sin_motor"
)

// AIInsight es una fila de `ai_insights`: **un campo de un ticket**, con sus dos redacciones.
//
// Las columnas que la base admite nulas van con puntero, como en el resto del proyecto: así se
// distingue «no hay texto» de «el texto está vacío», y un campo `pendiente` se lee como lo que es.
type AIInsight struct {
	ID           int64  `gorm:"primaryKey"`
	TicketNumber string `gorm:"column:ticket_number"`
	Kind         string `gorm:"column:kind"`
	State        string `gorm:"column:state"`
	// Las dos redacciones, en la misma fila: el motor las devuelve juntas (decisión 3).
	TextEs   *string `gorm:"column:text_es"`
	TextEn   *string `gorm:"column:text_en"`
	Text     *string `gorm:"column:text"`
	Language *string `gorm:"column:language"`
	Provider *string `gorm:"column:provider"`
	// La clave del error (`ai.unavailable` o `ai.invalid`) cuando el estado no es `listo`.
	ErrorKey *string `gorm:"column:error_key"`
	// El modelo que lo escribió: dice con qué se generó.
	Model          *string    `gorm:"column:model"`
	Attempts       int        `gorm:"column:attempts"`
	RequestedAt    time.Time  `gorm:"column:requested_at"`
	GeneratedAt    *time.Time `gorm:"column:generated_at"`
	ContextVersion int        `gorm:"column:context_version"`
	Generation     int64      `gorm:"column:generation"`
	RecoveryCount  int        `gorm:"column:recovery_count"`
	RetryAt        *time.Time `gorm:"column:retry_at"`
}

func (r *AIRepository) GuardarResultadoGlobal(numero, tipo, text, language, provider, model string, generatedAt time.Time) error {
	version := contextVersion(tipo)
	row := AIInsight{TicketNumber: numero, Kind: tipo, State: StateListo, Text: &text, Language: &language, Provider: &provider, Model: &model, RequestedAt: generatedAt, GeneratedAt: &generatedAt, ContextVersion: version}
	return r.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "ticket_number"}, {Name: "kind"}}, DoUpdates: clause.Assignments(map[string]any{
		"state": StateListo, "text": text, "language": language, "provider": provider, "text_es": nil, "text_en": nil, "error_key": nil, "model": model, "generated_at": generatedAt, "context_version": version,
	})}).Create(&row).Error
}

// TableName fija el nombre de la tabla, que no se deduce del tipo.
func (AIInsight) TableName() string { return "ai_insights" }

// AIRepository lee y escribe los resúmenes del motor.
type AIRepository struct {
	db *gorm.DB
}

// NewAIRepository construye el repositorio.
func NewAIRepository(db *gorm.DB) *AIRepository {
	return &AIRepository{db: db}
}

// MarcarPendiente deja el campo apuntado como `pendiente` antes de encolarlo.
//
// Es un `INSERT ... ON CONFLICT`, y no un `UPDATE`: la fila del campo puede no existir todavía —es la
// primera vez que se pide para ese ticket— y volver a pedirlo **reescribe la misma fila** en lugar de
// acumular historial, que es lo que fija la clave única `(ticket_number, kind)` (sección 4).
//
// El texto anterior puede conservarse en la fila como dato diagnóstico, pero deja de ser vigente:
// la API no lo entrega mientras el estado sea pendiente ni después de un error.
func (r *AIRepository) MarcarPendiente(numero, tipo string) error {
	_, err := r.MarcarPendienteConGeneracion(numero, tipo)
	return err
}

// MarcarPendienteConGeneracion abre una revisión nueva y devuelve su número. El incremento ocurre
// en la misma escritura que marca el campo pendiente, por lo que dos peticiones concurrentes no
// pueden recibir la misma revisión.
func (r *AIRepository) MarcarPendienteConGeneracion(numero, tipo string) (int64, error) {
	ahora := time.Now()
	var generacion int64
	err := r.db.Raw(`
		INSERT INTO ai_insights (ticket_number, kind, state, requested_at, generation, recovery_count)
		VALUES (?, ?, ?, ?, 1, 0)
		ON CONFLICT (ticket_number, kind) DO UPDATE SET
			state = EXCLUDED.state,
			error_key = NULL,
			attempts = 0,
			requested_at = EXCLUDED.requested_at,
			generation = ai_insights.generation + 1,
			recovery_count = 0,
			retry_at = NULL
		RETURNING generation
	`, numero, tipo, StatePendiente, ahora).Scan(&generacion).Error
	return generacion, err
}

// PrepararRecuperacion abre el único ciclo diferido, sólo si la revisión que falló sigue vigente.
func (r *AIRepository) PrepararRecuperacion(numero, tipo string, generacion int64, noAntes time.Time) (int64, bool, error) {
	var nueva int64
	resultado := r.db.Raw(`
		UPDATE ai_insights SET
			state = ?, error_key = NULL, attempts = 0, requested_at = now(),
			generation = generation + 1, recovery_count = 1, retry_at = ?
		WHERE ticket_number = ? AND kind = ? AND generation = ? AND recovery_count = 0
		RETURNING generation
	`, StatePendiente, noAntes, numero, tipo, generacion).Scan(&nueva)
	return nueva, resultado.RowsAffected == 1, resultado.Error
}

// ErroresParaRecuperar devuelve únicamente errores del contrato anterior que todavía no recibieron
// el ciclo adicional. PrepararRecuperacion decide de forma atómica cuál sigue vigente.
func (r *AIRepository) ErroresParaRecuperar(limite int) ([]AIInsight, error) {
	var filas []AIInsight
	err := r.db.Where("state IN ? AND recovery_count = 0", []string{StateError, StateSinMotor}).
		Order("requested_at ASC").Limit(limite).Find(&filas).Error
	return filas, err
}

// GuardarResultado guarda las dos redacciones y deja el campo en `listo`.
//
// Los intentos **no se tocan** aquí: los lleva `SubirIntentos`, que es quien cuenta las llamadas al
// motor, y dejarlos a cero al guardar borraría esa cuenta.
func (r *AIRepository) GuardarResultado(numero, tipo, es, en, modelo string, generadoEn time.Time) error {
	fila := AIInsight{
		TicketNumber:   numero,
		Kind:           tipo,
		State:          StateListo,
		TextEs:         &es,
		TextEn:         &en,
		Model:          &modelo,
		RequestedAt:    generadoEn,
		GeneratedAt:    &generadoEn,
		ContextVersion: contextVersion(tipo),
	}

	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "ticket_number"}, {Name: "kind"}},
		DoUpdates: clause.Assignments(map[string]any{
			"state":           StateListo,
			"text_es":         es,
			"text_en":         en,
			"error_key":       nil,
			"model":           modelo,
			"generated_at":    generadoEn,
			"context_version": contextVersion(tipo),
		}),
	}).Create(&fila).Error
}

func (r *AIRepository) GuardarResultadoSiVigente(numero, tipo string, generacion int64, es, en, modelo string, generadoEn time.Time) (bool, error) {
	resultado := r.db.Model(&AIInsight{}).Where("ticket_number = ? AND kind = ? AND generation = ?", numero, tipo, generacion).
		Updates(map[string]any{"state": StateListo, "text_es": es, "text_en": en, "error_key": nil, "model": modelo, "generated_at": generadoEn, "context_version": contextVersion(tipo), "retry_at": nil})
	return resultado.RowsAffected == 1, resultado.Error
}

func (r *AIRepository) GuardarResultadoGlobalSiVigente(numero, tipo string, generacion int64, text, language, provider, model string, generatedAt time.Time) (bool, error) {
	resultado := r.db.Model(&AIInsight{}).Where("ticket_number = ? AND kind = ? AND generation = ?", numero, tipo, generacion).
		Updates(map[string]any{"state": StateListo, "text": text, "language": language, "provider": provider, "text_es": nil, "text_en": nil, "error_key": nil, "model": model, "generated_at": generatedAt, "context_version": contextVersion(tipo), "retry_at": nil})
	return resultado.RowsAffected == 1, resultado.Error
}

// MarcarError deja el campo en `error`: el motor contestó, pero lo que contestó no vale.
func (r *AIRepository) MarcarError(numero, tipo, clave string) error {
	return r.marcarFallo(numero, tipo, StateError, clave)
}

// MarcarSinMotor deja el campo en `sin_motor`: hay motor configurado pero no contesta. Es el estado
// de una instalación sin motor y el de un contenedor caído, que para quien lee son lo mismo.
func (r *AIRepository) MarcarSinMotor(numero, tipo, clave string) error {
	return r.marcarFallo(numero, tipo, StateSinMotor, clave)
}

// SubirIntentos suma uno al contador del campo.
//
// Es lo que corta el reintento infinito: sin este número, un contenedor caído se reintentaría para
// siempre cada vez que el backend arranca (decisión 8).
func (r *AIRepository) SubirIntentos(numero, tipo string) error {
	return r.db.Model(&AIInsight{}).
		Where("ticket_number = ? AND kind = ?", numero, tipo).
		UpdateColumn("attempts", gorm.Expr("attempts + 1")).Error
}

func (r *AIRepository) SubirIntentosSiVigente(numero, tipo string, generacion int64) (bool, error) {
	resultado := r.db.Model(&AIInsight{}).
		Where("ticket_number = ? AND kind = ? AND generation = ?", numero, tipo, generacion).
		UpdateColumn("attempts", gorm.Expr("attempts + 1"))
	return resultado.RowsAffected == 1, resultado.Error
}

func (r *AIRepository) MarcarFalloSiVigente(numero, tipo string, generacion int64, estado, clave string) (bool, error) {
	resultado := r.db.Model(&AIInsight{}).
		Where("ticket_number = ? AND kind = ? AND generation = ?", numero, tipo, generacion).
		Updates(map[string]any{"state": estado, "error_key": clave, "retry_at": nil})
	return resultado.RowsAffected == 1, resultado.Error
}

func (r *AIRepository) GeneracionVigente(numero, tipo string, generacion int64) (bool, error) {
	var total int64
	err := r.db.Model(&AIInsight{}).
		Where("ticket_number = ? AND kind = ? AND generation = ?", numero, tipo, generacion).
		Count(&total).Error
	return total == 1, err
}

// PorNumeros trae los resúmenes de varios tickets **de una vez**: es lo que hace que la lista se
// pinte con una sola consulta y no con una por fila (decisión 9).
func (r *AIRepository) PorNumeros(numeros []string) ([]AIInsight, error) {
	if len(numeros) == 0 {
		return nil, nil
	}

	var filas []AIInsight
	if err := r.db.Where("ticket_number IN ?", numeros).Find(&filas).Error; err != nil {
		return nil, err
	}

	return filas, nil
}

// Pendientes son los campos que quedaron a medias y se pueden volver a pedir: los `pendiente` de una
// cola que se perdió al reiniciar el backend y los `sin_motor` que **no han agotado sus intentos**
// (decisión 10).
//
// Los errores no salen de aquí: su único ciclo adicional se selecciona aparte mediante
// `ErroresParaRecuperar`, que persiste `recovery_count` y evita reintentos infinitos.
//
// El orden es por `requested_at`: lo que lleva más tiempo esperando se retoma primero.
func (r *AIRepository) Pendientes(maxIntentos int) ([]AIInsight, error) {
	var filas []AIInsight

	err := r.db.
		Where(
			"state = ? OR (state = ? AND attempts < ?)",
			StatePendiente, StateSinMotor, maxIntentos,
		).
		Order("requested_at ASC").
		Find(&filas).Error
	if err != nil {
		return nil, err
	}

	return filas, nil
}

// UltimasAccionesDesactualizadas devuelve un lote persistente de resultados creados con un
// contrato de contexto anterior. La versión cambia sólo al guardar el resultado nuevo, de modo que
// un reinicio vuelve a recoger lo que no alcanzó a terminar.
func (r *AIRepository) UltimasAccionesDesactualizadas(version, limite int) ([]AIInsight, error) {
	var filas []AIInsight
	err := r.db.Where("kind = ? AND context_version < ?", KindUltimaAccion, version).
		Order("requested_at ASC").Limit(limite).Find(&filas).Error
	return filas, err
}

func (r *AIRepository) NumerosConInsightVigente(kind string, version int) ([]string, error) {
	var numeros []string
	err := r.db.Model(&AIInsight{}).
		Where("kind = ? AND context_version >= ?", kind, version).
		Pluck("ticket_number", &numeros).Error
	return numeros, err
}

// marcarFallo apunta el estado y su clave conservando el último texto válido. El estado deja claro
// que está desactualizado y permite reintentar sin sustituir información útil por un error.
func (r *AIRepository) marcarFallo(numero, tipo, estado, clave string) error {
	fila := AIInsight{
		TicketNumber: numero,
		Kind:         tipo,
		State:        estado,
		ErrorKey:     &clave,
		RequestedAt:  time.Now(),
	}

	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "ticket_number"}, {Name: "kind"}},
		DoUpdates: clause.Assignments(map[string]any{
			"state":     estado,
			"error_key": clave,
		}),
	}).Create(&fila).Error
}
