package services

import (
	"archive/zip"
	"bytes"
	"io"
	"strings"
	"testing"
	"time"

	"catalina-support/backend/shared/auth"
)

type sourceStub struct {
	rows    []Row
	filters []Filters
}

func (s *sourceStub) Rows(f Filters) ([]Row, error) {
	s.filters = append(s.filters, f)
	return s.rows, nil
}
func (s *sourceStub) Thresholds() (Thresholds, error) {
	return Thresholds{Critical: 4, High: 24, Medium: 72, Low: 120}, nil
}

func TestDeveloperIsRestrictedToInternalTickets(t *testing.T) {
	source := &sourceStub{}
	service := NewService(source)
	_, err := service.Summary(Filters{From: time.Now().Add(-time.Hour), To: time.Now()}, auth.Identity{Role: auth.RoleDesarrollo})
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range source.filters {
		if f.Type != "internal" {
			t.Fatalf("developer scope = %q", f.Type)
		}
	}
}
func TestUserCannotReadDashboard(t *testing.T) {
	service := NewService(&sourceStub{})
	_, err := service.Summary(Filters{}, auth.Identity{Role: auth.RoleUsuario})
	if err != ErrForbidden {
		t.Fatalf("error = %v", err)
	}
}
func TestSummaryDoesNotTurnMissingDurationsIntoZero(t *testing.T) {
	now := time.Date(2026, 10, 18, 12, 0, 0, 0, time.UTC)
	source := &sourceStub{rows: []Row{{Number: "CS-1", State: "nuevo", Priority: "critical", CreatedAt: now.Add(-5 * time.Hour), UpdatedAt: now, AgeMinutes: 300}}}
	service := NewService(source)
	service.now = func() time.Time { return now }
	v, err := service.Summary(Filters{From: now.Add(-24 * time.Hour), To: now}, auth.Identity{Role: auth.RoleSoporte})
	if err != nil {
		t.Fatal(err)
	}
	if v.Overdue.Value != 1 || v.Unanswered != 1 || v.AverageFirstResponseMinutes.Available {
		t.Fatalf("unexpected summary: %+v", v)
	}
}
func TestXLSXContainsSixLocalizedAggregateSheets(t *testing.T) {
	now := time.Date(2026, 10, 18, 12, 0, 0, 0, time.UTC)
	source := &sourceStub{rows: []Row{{Number: "=cmd", Subject: "+formula", State: "nuevo", Priority: "medium", Category: "Software", Assignee: "Ada", CreatedAt: now.Add(-time.Hour), UpdatedAt: now}}}
	service := NewService(source)
	var out bytes.Buffer
	if err := service.XLSX(&out, Filters{From: now.Add(-24 * time.Hour), To: now.Add(time.Hour)}, auth.Identity{Role: auth.RoleSoporte}, "es", "America/Guayaquil"); err != nil {
		t.Fatal(err)
	}
	reader, err := zip.NewReader(bytes.NewReader(out.Bytes()), int64(out.Len()))
	if err != nil {
		t.Fatal(err)
	}
	files := map[string]string{}
	for _, file := range reader.File {
		opened, openErr := file.Open()
		if openErr != nil {
			t.Fatal(openErr)
		}
		body, readErr := io.ReadAll(opened)
		_ = opened.Close()
		if readErr != nil {
			t.Fatal(readErr)
		}
		files[file.Name] = string(body)
	}
	workbook := files["xl/workbook.xml"]
	for _, name := range []string{"Resumen", "Serie temporal", "Estado", "Prioridad", "Categoría", "Responsable"} {
		if !strings.Contains(workbook, `name="`+name+`"`) {
			t.Fatalf("missing sheet %q in %s", name, workbook)
		}
	}
	if len(files) < 11 || !strings.Contains(files["xl/worksheets/sheet2.xml"], `state="frozen"`) || !strings.Contains(files["xl/worksheets/sheet2.xml"], "<autoFilter") {
		t.Fatalf("workbook formatting is incomplete")
	}
	for _, body := range files {
		if strings.Contains(body, "=cmd") || strings.Contains(body, "+formula") {
			t.Fatalf("individual ticket data leaked into aggregate workbook")
		}
	}
}

func TestTimelineUsesDaysThroughThirtyOneDaysAndWeeksAfter(t *testing.T) {
	from := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	if got := aggregate(nil, Thresholds{}, from, from.AddDate(0, 0, 31), from).Grouping; got != "day" {
		t.Fatalf("31 days grouped as %q", got)
	}
	if got := aggregate(nil, Thresholds{}, from, from.AddDate(0, 0, 32), from).Grouping; got != "week" {
		t.Fatalf("32 days grouped as %q", got)
	}
}
