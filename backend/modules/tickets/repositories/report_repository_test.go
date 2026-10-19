package repositories

import (
	"testing"
	"time"

	dashboard "catalina-support/backend/modules/dashboard/services"
)

// El informe se prueba contra PostgreSQL real porque su contrato depende de ventanas, eventos y
// pausas de antigüedad que una base simulada no reproduce con fidelidad.
func TestReportRepositoryReadsOperationalFactsFromPostgreSQL(t *testing.T) {
	_, _, tx := repositorioDePrueba(t)
	defer tx.Rollback()

	report := NewReportRepository(tx)
	thresholds, err := report.Thresholds()
	if err != nil {
		t.Fatalf("Thresholds falló: %v", err)
	}
	if thresholds.Critical <= 0 || thresholds.High <= 0 || thresholds.Medium <= 0 || thresholds.Low <= 0 {
		t.Fatalf("umbrales inválidos: %+v", thresholds)
	}

	to := time.Now().UTC().Add(24 * time.Hour)
	rows, err := report.Rows(dashboard.Filters{From: to.AddDate(-10, 0, 0), To: to, Snapshot: true})
	if err != nil {
		t.Fatalf("Rows falló: %v", err)
	}
	if len(rows) == 0 {
		t.Fatal("el informe no devolvió los tickets sembrados")
	}
	for _, row := range rows {
		if row.Number == "" || (row.Type != "main" && row.Type != "internal") || row.AgeMinutes < 0 {
			t.Fatalf("fila operativa inválida: %+v", row)
		}
	}
}
