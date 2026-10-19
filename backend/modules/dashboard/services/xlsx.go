package services

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"catalina-support/backend/shared/auth"
)

type workbookCell struct {
	text   string
	number *float64
	date   *time.Time
}

type workbookSheet struct {
	name string
	rows [][]workbookCell
}

// XLSX exports only the aggregated information shown by the Dashboard.
func (s *Service) XLSX(w io.Writer, filters Filters, identity auth.Identity, language, timeZone string) error {
	summary, err := s.Summary(filters, identity)
	if err != nil {
		return err
	}
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		location = time.UTC
	}
	return writeWorkbook(w, dashboardSheets(summary, filters, language, location))
}

func dashboardSheets(s Summary, f Filters, language string, location *time.Location) []workbookSheet {
	tr := dashboardTranslations(language)
	to := f.To.In(location).Add(-time.Nanosecond)
	rows := [][]workbookCell{{txt(tr.section), txt(tr.metric), txt(tr.value), txt(tr.minutes), txt(tr.previous), txt(tr.previousMinutes), txt(tr.available)}}
	filters := [][2]string{
		{tr.from, f.From.In(location).Format("2006-01-02")},
		{tr.to, to.Format("2006-01-02")},
		{tr.typeLabel, localizedValue("type", f.Type, language, tr.all)},
		{tr.stateLabel, localizedValue("state", f.State, language, tr.all)},
		{tr.priorityLabel, localizedValue("priority", f.Priority, language, tr.all)},
	}
	for _, item := range filters {
		rows = append(rows, []workbookCell{txt(tr.filters), txt(item[0]), txt(item[1]), txt(""), txt(""), txt(""), txt(tr.yes)})
	}
	metrics := []struct {
		name     string
		m        Metric
		duration bool
	}{
		{tr.open, s.Open, false}, {tr.openCritical, s.OpenCritical, false}, {tr.overdue, s.Overdue, false},
		{tr.firstResponse, s.AverageFirstResponseMinutes, true}, {tr.resolution, s.AverageResolutionMinutes, true},
		{tr.reopens, s.Reopens, false},
	}
	for _, item := range metrics {
		if item.duration {
			rows = append(rows, []workbookCell{txt(tr.indicators), txt(item.name), txt(human(int64(item.m.Value), language)), num(item.m.Value), txt(human(int64(item.m.Previous), language)), num(item.m.Previous), txt(yesNo(item.m.Available, tr))})
		} else {
			rows = append(rows, []workbookCell{txt(tr.indicators), txt(item.name), num(item.m.Value), txt(""), num(item.m.Previous), txt(""), txt(yesNo(item.m.Available, tr))})
		}
	}
	rows = append(rows, []workbookCell{txt(tr.indicators), txt(tr.unanswered), num(float64(s.Unanswered)), txt(""), txt(""), txt(""), txt(tr.yes)})

	return []workbookSheet{
		{name: tr.summary, rows: rows},
		{name: tr.timeline, rows: timelineRows(s.CreatedClosed, s.Grouping, tr, location)},
		{name: tr.state, rows: pointRows(s.ByState, "state", tr, language)},
		{name: tr.priority, rows: pointRows(s.ByPriority, "priority", tr, language)},
		{name: tr.category, rows: pointRows(s.ByCategory, "category", tr, language)},
		{name: tr.assignee, rows: pointRows(s.ByAssignee, "assignee", tr, language)},
	}
}

func timelineRows(points []Point, grouping string, tr dashboardText, location *time.Location) [][]workbookCell {
	rows := [][]workbookCell{{txt(tr.period), txt(tr.periodISO), txt(tr.created), txt(tr.closed)}}
	for _, point := range points {
		date, ok := periodStart(point.Key, grouping, location)
		if ok {
			rows = append(rows, []workbookCell{dateCell(date), txt(date.Format("2006-01-02")), num(float64(point.Value)), num(float64(point.Secondary))})
		} else {
			rows = append(rows, []workbookCell{txt(point.Key), txt(point.Key), num(float64(point.Value)), num(float64(point.Secondary))})
		}
	}
	return rows
}

func pointRows(points []Point, kind string, tr dashboardText, language string) [][]workbookCell {
	rows := [][]workbookCell{{txt(tr.label), txt(tr.value)}}
	for _, point := range points {
		rows = append(rows, []workbookCell{txt(localizedValue(kind, point.Key, language, point.Key)), num(float64(point.Value))})
	}
	return rows
}

func periodStart(key, grouping string, location *time.Location) (time.Time, bool) {
	if grouping == "day" {
		parsed, err := time.ParseInLocation("2006-01-02", key, location)
		return parsed, err == nil
	}
	var year, week int
	if _, err := fmt.Sscanf(key, "%d-W%d", &year, &week); err != nil {
		return time.Time{}, false
	}
	jan4 := time.Date(year, 1, 4, 0, 0, 0, 0, location)
	monday := jan4.AddDate(0, 0, -((int(jan4.Weekday()) + 6) % 7))
	return monday.AddDate(0, 0, (week-1)*7), true
}

func writeWorkbook(w io.Writer, sheets []workbookSheet) error {
	zipWriter := zip.NewWriter(w)
	files := map[string]string{
		"[Content_Types].xml":        contentTypes(len(sheets)),
		"_rels/.rels":                packageRelationships,
		"xl/workbook.xml":            workbookXML(sheets),
		"xl/_rels/workbook.xml.rels": workbookRelationships(len(sheets)),
		"xl/styles.xml":              stylesXML,
	}
	for index, sheet := range sheets {
		files[fmt.Sprintf("xl/worksheets/sheet%d.xml", index+1)] = worksheetXML(sheet.rows)
	}
	for name, body := range files {
		entry, err := zipWriter.Create(name)
		if err != nil {
			return err
		}
		if _, err = io.WriteString(entry, body); err != nil {
			return err
		}
	}
	return zipWriter.Close()
}

func worksheetXML(rows [][]workbookCell) string {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><worksheet xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main"><sheetViews><sheetView workbookViewId="0"><pane ySplit="1" topLeftCell="A2" activePane="bottomLeft" state="frozen"/></sheetView></sheetViews><cols><col min="1" max="10" width="24" customWidth="1"/></cols><sheetData>`)
	for rowIndex, row := range rows {
		fmt.Fprintf(&b, `<row r="%d">`, rowIndex+1)
		for columnIndex, cell := range row {
			reference := columnName(columnIndex+1) + strconv.Itoa(rowIndex+1)
			switch {
			case cell.date != nil:
				fmt.Fprintf(&b, `<c r="%s" s="2"><v>%.8f</v></c>`, reference, excelDate(*cell.date))
			case cell.number != nil:
				fmt.Fprintf(&b, `<c r="%s"%s><v>%s</v></c>`, reference, headerStyle(rowIndex), strconv.FormatFloat(*cell.number, 'f', -1, 64))
			default:
				fmt.Fprintf(&b, `<c r="%s" t="inlineStr"%s><is><t xml:space="preserve">%s</t></is></c>`, reference, headerStyle(rowIndex), xmlText(cell.text))
			}
		}
		b.WriteString(`</row>`)
	}
	lastColumn := "A"
	if len(rows) > 0 && len(rows[0]) > 0 {
		lastColumn = columnName(len(rows[0]))
	}
	fmt.Fprintf(&b, `</sheetData><autoFilter ref="A1:%s%d"/></worksheet>`, lastColumn, maxInt(1, len(rows)))
	return b.String()
}

func contentTypes(count int) string {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types"><Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/><Default Extension="xml" ContentType="application/xml"/><Override PartName="/xl/workbook.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"/><Override PartName="/xl/styles.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"/>`)
	for index := 1; index <= count; index++ {
		fmt.Fprintf(&b, `<Override PartName="/xl/worksheets/sheet%d.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"/>`, index)
	}
	b.WriteString(`</Types>`)
	return b.String()
}

func workbookXML(sheets []workbookSheet) string {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><workbook xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"><sheets>`)
	for index, sheet := range sheets {
		fmt.Fprintf(&b, `<sheet name="%s" sheetId="%d" r:id="rId%d"/>`, xmlText(sheet.name), index+1, index+1)
	}
	b.WriteString(`</sheets></workbook>`)
	return b.String()
}

func workbookRelationships(count int) string {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">`)
	for index := 1; index <= count; index++ {
		fmt.Fprintf(&b, `<Relationship Id="rId%d" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet" Target="worksheets/sheet%d.xml"/>`, index, index)
	}
	fmt.Fprintf(&b, `<Relationship Id="rId%d" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles" Target="styles.xml"/></Relationships>`, count+1)
	return b.String()
}

const packageRelationships = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships"><Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="xl/workbook.xml"/></Relationships>`
const stylesXML = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><styleSheet xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main"><fonts count="2"><font><sz val="11"/><name val="Arial"/></font><font><b/><color rgb="FFFFFFFF"/><sz val="11"/><name val="Arial"/></font></fonts><fills count="3"><fill><patternFill patternType="none"/></fill><fill><patternFill patternType="gray125"/></fill><fill><patternFill patternType="solid"><fgColor rgb="FF1D4ED8"/><bgColor indexed="64"/></patternFill></fill></fills><borders count="1"><border><left/><right/><top/><bottom/><diagonal/></border></borders><cellStyleXfs count="1"><xf numFmtId="0" fontId="0" fillId="0" borderId="0"/></cellStyleXfs><cellXfs count="3"><xf numFmtId="0" fontId="0" fillId="0" borderId="0" xfId="0"/><xf numFmtId="0" fontId="1" fillId="2" borderId="0" xfId="0" applyFont="1" applyFill="1"/><xf numFmtId="14" fontId="0" fillId="0" borderId="0" xfId="0" applyNumberFormat="1"/></cellXfs></styleSheet>`

func txt(value string) workbookCell         { return workbookCell{text: value} }
func num(value float64) workbookCell        { return workbookCell{number: &value} }
func dateCell(value time.Time) workbookCell { return workbookCell{date: &value} }
func headerStyle(row int) string {
	if row == 0 {
		return ` s="1"`
	}
	return ""
}
func excelDate(value time.Time) float64 {
	return value.Sub(time.Date(1899, 12, 30, 0, 0, 0, 0, value.Location())).Hours() / 24
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func columnName(value int) string {
	var result string
	for value > 0 {
		value--
		result = string(rune('A'+value%26)) + result
		value /= 26
	}
	return result
}
func xmlText(value string) string {
	var b bytes.Buffer
	for _, r := range value {
		switch r {
		case '&':
			b.WriteString("&amp;")
		case '<':
			b.WriteString("&lt;")
		case '>':
			b.WriteString("&gt;")
		case '"':
			b.WriteString("&quot;")
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}

type dashboardText struct {
	summary, timeline, state, priority, category, assignee                                     string
	section, metric, value, minutes, previous, previousMinutes, available, filters, indicators string
	from, to, typeLabel, stateLabel, priorityLabel, all                                        string
	open, openCritical, overdue, firstResponse, resolution, reopens, unanswered                string
	period, periodISO, created, closed, label, yes, no                                         string
}

func dashboardTranslations(language string) dashboardText {
	if language == "es" {
		return dashboardText{"Resumen", "Serie temporal", "Estado", "Prioridad", "Categoría", "Responsable", "Sección", "Métrica", "Valor", "Minutos", "Periodo anterior", "Minutos anteriores", "Disponible", "Filtros", "Indicadores", "Desde", "Hasta", "Tipo", "Estado", "Prioridad", "Todos", "Abiertos", "Críticos abiertos", "Atrasados", "Primera respuesta", "Resolución", "Reaperturas", "Sin respuesta", "Periodo", "Periodo ISO", "Creados", "Cerrados", "Etiqueta", "Sí", "No"}
	}
	return dashboardText{"Summary", "Timeline", "State", "Priority", "Category", "Assignee", "Section", "Metric", "Value", "Minutes", "Previous period", "Previous minutes", "Available", "Filters", "Indicators", "From", "To", "Type", "State", "Priority", "All", "Open", "Open critical", "Overdue", "First response", "Resolution", "Reopens", "Unanswered", "Period", "Period ISO", "Created", "Closed", "Label", "Yes", "No"}
}

func yesNo(value bool, tr dashboardText) string {
	if value {
		return tr.yes
	}
	return tr.no
}
func localizedValue(kind, key, language, fallback string) string {
	if strings.TrimSpace(key) == "" {
		return fallback
	}
	values := map[string]map[string][2]string{
		"type":     {"main": {"Main", "Principal"}, "internal": {"Internal", "Interno"}},
		"state":    {"nuevo": {"New", "Nuevo"}, "en progreso": {"In progress", "En progreso"}, "en espera": {"Waiting", "Esperando"}, "escalado": {"Escalated", "Escalado"}, "resuelto": {"Resolved", "Resuelto"}, "cerrado": {"Closed", "Cerrado"}},
		"priority": {"critical": {"Critical", "Crítica"}, "high": {"High", "Alta"}, "medium": {"Medium", "Media"}, "low": {"Low", "Baja"}},
		"category": {"other": {"Other", "Otras"}},
		"assignee": {"unassigned": {"Unassigned", "Sin responsable"}},
	}
	if pair, ok := values[kind][key]; ok {
		if language == "es" {
			return pair[1]
		}
		return pair[0]
	}
	return key
}
