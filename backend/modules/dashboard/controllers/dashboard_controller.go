// Package controllers exposes the Dashboard API.
package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"catalina-support/backend/modules/dashboard/services"
	"catalina-support/backend/shared/auth"
	"catalina-support/backend/shared/httpx"
)

type DashboardController struct {
	service  *services.Service
	timeZone func() string
	language func() (string, error)
}

func NewDashboardController(service *services.Service, timeZone func() string, language func() (string, error)) *DashboardController {
	return &DashboardController{service: service, timeZone: timeZone, language: language}
}

func (c *DashboardController) Summary(w http.ResponseWriter, r *http.Request) {
	f, ok := c.filters(w, r)
	if !ok {
		return
	}
	v, err := c.service.Summary(f, auth.MustFromContext(r.Context()))
	if err != nil {
		c.fail(w, err)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, v)
}
func (c *DashboardController) Export(w http.ResponseWriter, r *http.Request) {
	f, ok := c.filters(w, r)
	if !ok {
		return
	}
	language, _ := c.language()
	var output bytes.Buffer
	if err := c.service.XLSX(&output, f, auth.MustFromContext(r.Context()), language, c.timeZone()); err != nil {
		c.fail(w, err)
		return
	}
	loc, err := time.LoadLocation(c.timeZone())
	if err != nil {
		loc = time.UTC
	}
	filename := fmt.Sprintf("report_%s_%s.xlsx", f.From.In(loc).Format("2006-01-02"), f.To.In(loc).Add(-time.Nanosecond).Format("2006-01-02"))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(output.Bytes())
}

func (c *DashboardController) filters(w http.ResponseWriter, r *http.Request) (services.Filters, bool) {
	loc, err := time.LoadLocation(c.timeZone())
	if err != nil {
		loc = time.UTC
	}
	q := r.URL.Query()
	now := time.Now().In(loc)
	to := now
	from := now.AddDate(0, 0, -30)
	if v := q.Get("from"); v != "" {
		from, err = time.ParseInLocation("2006-01-02", v, loc)
		if err != nil {
			httpx.WriteError(w, 400, "dashboard.period.invalid")
			return services.Filters{}, false
		}
	}
	if v := q.Get("to"); v != "" {
		to, err = time.ParseInLocation("2006-01-02", v, loc)
		if err != nil {
			httpx.WriteError(w, 400, "dashboard.period.invalid")
			return services.Filters{}, false
		}
		to = to.AddDate(0, 0, 1)
	}
	if !from.Before(to) {
		httpx.WriteError(w, 400, "dashboard.period.invalid")
		return services.Filters{}, false
	}
	f := services.Filters{From: from.UTC(), To: to.UTC(), Type: strings.TrimSpace(q.Get("type")), State: strings.TrimSpace(q.Get("state")), Priority: strings.TrimSpace(q.Get("priority"))}
	return f, true
}
func (c *DashboardController) fail(w http.ResponseWriter, err error) {
	if err == services.ErrForbidden {
		httpx.WriteError(w, 403, "auth.forbidden")
		return
	}
	httpx.WriteError(w, 500, "dashboard.unavailable")
}
