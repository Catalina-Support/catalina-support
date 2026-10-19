// Package repositories exposes the ticket facts needed by operational reporting.
package repositories

import (
	"strings"

	"catalina-support/backend/modules/dashboard/services"
	"gorm.io/gorm"
)

type ReportRepository struct{ db *gorm.DB }

func NewReportRepository(db *gorm.DB) *ReportRepository { return &ReportRepository{db: db} }

func (r *ReportRepository) Thresholds() (services.Thresholds, error) {
	var row struct{ Critical, High, Medium, Low int }
	err := r.db.Raw(`SELECT report_overdue_critical_hours critical, report_overdue_high_hours high,
report_overdue_medium_hours medium, report_overdue_low_hours low FROM ticket_settings WHERE id=1`).Scan(&row).Error
	return services.Thresholds{Critical: row.Critical, High: row.High, Medium: row.Medium, Low: row.Low}, err
}

func (r *ReportRepository) Rows(f services.Filters) ([]services.Row, error) {
	thresholds, err := r.Thresholds()
	if err != nil {
		return nil, err
	}
	where := []string{"created_at < ?"}
	args := []any{f.From, f.To, f.To}
	if !f.Snapshot {
		where = append(where, "created_at >= ?")
		args = append(args, f.From)
	}
	if f.Type != "" && f.Type != "all" {
		where = append(where, "ticket_type = ?")
		args = append(args, f.Type)
	}
	if f.State != "" {
		where = append(where, "state = ?")
		args = append(args, f.State)
	}
	if f.Priority != "" {
		where = append(where, "priority = ?")
		args = append(args, f.Priority)
	}
	query := `WITH params AS (
 SELECT ?::timestamptz from_time, ?::timestamptz to_time
), state_events AS (
 SELECT 'main' ticket_type,t.id entity_id,t.created_at event_at,'nuevo' state FROM tickets t
 UNION ALL
 SELECT 'internal',i.id,i.created_at,'nuevo' FROM internal_tickets i
 UNION ALL
 SELECT CASE WHEN th.ticket_id IS NOT NULL THEN 'main' ELSE 'internal' END,
        coalesce(th.ticket_id,th.internal_ticket_id),th.created_at,th.to_state
 FROM ticket_history th
 WHERE th.to_state IS NOT NULL AND th.event <> 'creado'
), state_spans AS (
 SELECT ticket_type,entity_id,event_at,state,
        lead(event_at,1,least(now(),(SELECT to_time FROM params))) OVER (PARTITION BY ticket_type,entity_id ORDER BY event_at) next_at
 FROM state_events WHERE event_at < (SELECT to_time FROM params)
), active_age AS (
 SELECT ticket_type,entity_id,
        CAST(coalesce(sum(extract(epoch FROM (next_at-event_at))/60)
          FILTER (WHERE state NOT IN ('resuelto','cerrado') AND next_at > event_at),0) AS bigint) age_minutes
 FROM state_spans GROUP BY ticket_type,entity_id
), state_at_cutoff AS (
 SELECT DISTINCT ON (ticket_type,entity_id) ticket_type,entity_id,state
 FROM state_events WHERE event_at < (SELECT to_time FROM params)
 ORDER BY ticket_type,entity_id,event_at DESC
), report_rows AS (
 SELECT t.id entity_id,'main' ticket_type,t.number,t.priority,coalesce(s.state,t.state) state,t.subject,t.category_id,t.assignee_id,
 c.name category,trim(concat(u.name,' ',u.last_name)) assignee,t.created_at,t.updated_at,
 (SELECT min(tc.created_at) FROM ticket_comments tc JOIN users au ON au.id=tc.author_id,params p WHERE tc.ticket_id=t.id AND tc.deleted_at IS NULL AND au.role='soporte' AND tc.created_at<p.to_time) first_response_at,
 (SELECT min(th.created_at) FROM ticket_history th,params p WHERE th.ticket_id=t.id AND th.event='resuelto' AND th.created_at<p.to_time) first_resolution_at,
 (SELECT max(th.created_at) FROM ticket_history th,params p WHERE th.ticket_id=t.id AND th.event='cerrado' AND th.created_at<p.to_time) closed_at,
 (SELECT count(*) FROM ticket_history th,params p WHERE th.ticket_id=t.id AND th.event='reabierto' AND th.created_at>=p.from_time AND th.created_at<p.to_time) reopens,
	 coalesce(a.age_minutes,0) age_minutes
 FROM tickets t JOIN ticket_categories c ON c.id=t.category_id LEFT JOIN users u ON u.id=t.assignee_id
 LEFT JOIN active_age a ON a.ticket_type='main' AND a.entity_id=t.id
 LEFT JOIN state_at_cutoff s ON s.ticket_type='main' AND s.entity_id=t.id
 UNION ALL
 SELECT i.id,'internal',i.number,t.priority,coalesce(s.state,i.state),t.subject,t.category_id,i.assignee_id,
 c.name,trim(concat(u.name,' ',u.last_name)),i.created_at,i.updated_at,
 (SELECT min(tc.created_at) FROM ticket_comments tc JOIN users au ON au.id=tc.author_id,params p WHERE tc.internal_ticket_id=i.id AND tc.deleted_at IS NULL AND au.role='desarrollo' AND tc.created_at<p.to_time),
 (SELECT min(th.created_at) FROM ticket_history th,params p WHERE th.internal_ticket_id=i.id AND th.event='resuelto' AND th.created_at<p.to_time),
 (SELECT max(th.created_at) FROM ticket_history th,params p WHERE th.internal_ticket_id=i.id AND th.event='cerrado' AND th.created_at<p.to_time),
 (SELECT count(*) FROM ticket_history th,params p WHERE th.internal_ticket_id=i.id AND th.event='reabierto' AND th.created_at>=p.from_time AND th.created_at<p.to_time),
	 coalesce(a.age_minutes,0)
 FROM internal_tickets i JOIN tickets t ON t.id=i.ticket_id JOIN ticket_categories c ON c.id=t.category_id LEFT JOIN users u ON u.id=i.assignee_id
 LEFT JOIN active_age a ON a.ticket_type='internal' AND a.entity_id=i.id
 LEFT JOIN state_at_cutoff s ON s.ticket_type='internal' AND s.entity_id=i.id
) SELECT number,ticket_type AS type,priority,state,subject,category,assignee,created_at,updated_at,first_response_at,first_resolution_at,closed_at,age_minutes,reopens
FROM report_rows WHERE ` + strings.Join(where, " AND ") + `
ORDER BY (age_minutes > CASE priority WHEN 'critical' THEN CAST(? AS bigint) WHEN 'high' THEN CAST(? AS bigint) WHEN 'medium' THEN CAST(? AS bigint) ELSE CAST(? AS bigint) END) DESC,
CASE priority WHEN 'critical' THEN 4 WHEN 'high' THEN 3 WHEN 'medium' THEN 2 ELSE 1 END DESC, updated_at DESC`
	args = append(args, thresholds.Critical*60, thresholds.High*60, thresholds.Medium*60, thresholds.Low*60)
	var rows []services.Row
	err = r.db.Raw(query, args...).Scan(&rows).Error
	return rows, err
}
