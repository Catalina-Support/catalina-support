// Package services contains the operational reporting rules.
package services

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"catalina-support/backend/shared/auth"
)

var ErrForbidden = errors.New("auth.forbidden")

type Thresholds struct {
	Critical int `json:"critical"`
	High     int `json:"high"`
	Medium   int `json:"medium"`
	Low      int `json:"low"`
}

type Filters struct {
	From, To              time.Time
	Type, State, Priority string
	Snapshot              bool
}

type Row struct {
	Number            string     `json:"number"`
	Type              string     `json:"type"`
	Priority          string     `json:"priority"`
	State             string     `json:"state"`
	Subject           string     `json:"subject"`
	Category          string     `json:"category"`
	Assignee          string     `json:"assignee"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
	FirstResponseAt   *time.Time `json:"firstResponseAt"`
	FirstResolutionAt *time.Time `json:"firstResolutionAt"`
	ClosedAt          *time.Time `json:"closedAt"`
	AgeMinutes        int64      `json:"ageMinutes"`
	Reopens           int        `json:"reopens"`
}

type Source interface {
	Rows(Filters) ([]Row, error)
	Thresholds() (Thresholds, error)
}

type Service struct {
	source Source
	now    func() time.Time
}

func NewService(source Source) *Service { return &Service{source: source, now: time.Now} }

type Metric struct {
	Value     float64 `json:"value"`
	Previous  float64 `json:"previous"`
	Available bool    `json:"available"`
}
type Point struct {
	Key       string `json:"key"`
	Value     int    `json:"value"`
	Secondary int    `json:"secondary,omitempty"`
}
type Summary struct {
	Open                        Metric     `json:"open"`
	OpenCritical                Metric     `json:"openCritical"`
	Overdue                     Metric     `json:"overdue"`
	AverageFirstResponseMinutes Metric     `json:"averageFirstResponseMinutes"`
	AverageResolutionMinutes    Metric     `json:"averageResolutionMinutes"`
	Reopens                     Metric     `json:"reopens"`
	Unanswered                  int        `json:"unanswered"`
	CreatedClosed               []Point    `json:"createdClosed"`
	ByState                     []Point    `json:"byState"`
	ByPriority                  []Point    `json:"byPriority"`
	ByCategory                  []Point    `json:"byCategory"`
	ByAssignee                  []Point    `json:"byAssignee"`
	Grouping                    string     `json:"grouping"`
	Thresholds                  Thresholds `json:"thresholds"`
}

func (s *Service) allowed(identity auth.Identity) bool {
	return identity.IsAdministrator() || identity.Role == auth.RoleSoporte || identity.Role == auth.RoleDesarrollo
}
func scope(filters Filters, identity auth.Identity) Filters {
	if identity.Role == auth.RoleDesarrollo && !identity.IsAdministrator() {
		filters.Type = "internal"
	}
	return filters
}

func (s *Service) Summary(filters Filters, identity auth.Identity) (Summary, error) {
	if !s.allowed(identity) {
		return Summary{}, ErrForbidden
	}
	filters = scope(filters, identity)
	filters.Snapshot = true
	rows, err := s.source.Rows(filters)
	if err != nil {
		return Summary{}, err
	}
	duration := filters.To.Sub(filters.From)
	previous := filters
	previous.To = filters.From
	previous.From = filters.From.Add(-duration)
	previousRows, err := s.source.Rows(previous)
	if err != nil {
		return Summary{}, err
	}
	thresholds, err := s.source.Thresholds()
	if err != nil {
		return Summary{}, err
	}
	current := aggregate(rows, thresholds, filters.From, filters.To, s.now())
	prior := aggregate(previousRows, thresholds, previous.From, previous.To, s.now())
	current.Open.Previous = prior.Open.Value
	current.OpenCritical.Previous = prior.OpenCritical.Value
	current.Overdue.Previous = prior.Overdue.Value
	current.AverageFirstResponseMinutes.Previous = prior.AverageFirstResponseMinutes.Value
	current.AverageResolutionMinutes.Previous = prior.AverageResolutionMinutes.Value
	current.Reopens.Previous = prior.Reopens.Value
	return current, nil
}

func aggregate(rows []Row, thresholds Thresholds, from, to, now time.Time) Summary {
	result := Summary{Thresholds: thresholds, ByState: []Point{}, ByPriority: []Point{}, ByCategory: []Point{}, ByAssignee: []Point{}}
	days := to.Sub(from).Hours() / 24
	if days <= 31 {
		result.Grouping = "day"
	} else {
		result.Grouping = "week"
	}
	state, priority, category, assignee := map[string]int{}, map[string]int{}, map[string]int{}, map[string]int{}
	created, closed := map[string]*Point{}, map[string]*Point{}
	var firstTotal, resolutionTotal float64
	var firstCount, resolutionCount int
	for _, row := range rows {
		open := row.State != "cerrado"
		if open {
			result.Open.Value++
			result.Open.Available = true
			if row.Priority == "critical" {
				result.OpenCritical.Value++
				result.OpenCritical.Available = true
			}
		}
		if open && row.AgeMinutes > int64(thresholdFor(row.Priority, thresholds))*60 {
			result.Overdue.Value++
			result.Overdue.Available = true
		}
		createdInPeriod := !row.CreatedAt.Before(from) && row.CreatedAt.Before(to)
		if createdInPeriod && row.FirstResponseAt != nil {
			firstTotal += row.FirstResponseAt.Sub(row.CreatedAt).Minutes()
			firstCount++
		} else if createdInPeriod {
			result.Unanswered++
		}
		if row.FirstResolutionAt != nil && !row.FirstResolutionAt.Before(from) && row.FirstResolutionAt.Before(to) {
			resolutionTotal += row.FirstResolutionAt.Sub(row.CreatedAt).Minutes()
			resolutionCount++
		}
		result.Reopens.Value += float64(row.Reopens)
		result.Reopens.Available = true
		state[row.State]++
		priority[row.Priority]++
		category[blank(row.Category, "other")]++
		if open {
			assignee[blank(row.Assignee, "unassigned")]++
		}
		if createdInPeriod {
			key := periodKey(row.CreatedAt, result.Grouping)
			p := created[key]
			if p == nil {
				p = &Point{Key: key}
				created[key] = p
			}
			p.Value++
		}
		if row.ClosedAt != nil && !row.ClosedAt.Before(from) && row.ClosedAt.Before(to) {
			key := periodKey(*row.ClosedAt, result.Grouping)
			p := closed[key]
			if p == nil {
				p = &Point{Key: key}
				closed[key] = p
			}
			p.Secondary++
		}
	}
	result.Open.Available = true
	result.OpenCritical.Available = true
	result.Overdue.Available = true
	if firstCount > 0 {
		result.AverageFirstResponseMinutes = Metric{Value: firstTotal / float64(firstCount), Available: true}
	}
	if resolutionCount > 0 {
		result.AverageResolutionMinutes = Metric{Value: resolutionTotal / float64(resolutionCount), Available: true}
	}
	result.ByState = points(state, 0)
	result.ByPriority = points(priority, 0)
	result.ByCategory = points(category, 10)
	result.ByAssignee = points(assignee, 0)
	keys := map[string]bool{}
	for k := range created {
		keys[k] = true
	}
	for k := range closed {
		keys[k] = true
	}
	ordered := make([]string, 0, len(keys))
	for k := range keys {
		ordered = append(ordered, k)
	}
	sort.Strings(ordered)
	for _, k := range ordered {
		p := Point{Key: k}
		if created[k] != nil {
			p.Value = created[k].Value
		}
		if closed[k] != nil {
			p.Secondary = closed[k].Secondary
		}
		result.CreatedClosed = append(result.CreatedClosed, p)
	}
	return result
}
func thresholdFor(priority string, t Thresholds) int {
	switch priority {
	case "critical":
		return t.Critical
	case "high":
		return t.High
	case "low":
		return t.Low
	default:
		return t.Medium
	}
}
func blank(v, f string) string {
	if strings.TrimSpace(v) == "" {
		return f
	}
	return v
}
func periodKey(t time.Time, g string) string {
	if g == "week" {
		y, w := t.ISOWeek()
		return fmt.Sprintf("%04d-W%02d", y, w)
	}
	return t.Format("2006-01-02")
}
func points(values map[string]int, limit int) []Point {
	out := make([]Point, 0, len(values))
	for k, v := range values {
		out = append(out, Point{Key: k, Value: v})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Value == out[j].Value {
			return out[i].Key < out[j].Key
		}
		return out[i].Value > out[j].Value
	})
	if limit > 0 && len(out) > limit {
		other := 0
		for _, p := range out[limit:] {
			other += p.Value
		}
		out = append(out[:limit], Point{Key: "other", Value: other})
	}
	return out
}

func human(minutes int64, language string) string {
	h := minutes / 60
	m := minutes % 60
	if language == "es" {
		return fmt.Sprintf("%d h %d min", h, m)
	}
	return fmt.Sprintf("%d h %d min", h, m)
}
