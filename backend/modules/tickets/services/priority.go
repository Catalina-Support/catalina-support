package services

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"catalina-support/backend/modules/tickets/repositories"
	"catalina-support/backend/shared/auth"
)

// ChangePriority cambia la prioridad del caso. El dato vive en el principal y el interno lo hereda.
func (s *Service) ChangePriority(number, value string, actor auth.Identity) (Ticket, string, error) {
	priority, err := normalizarPrioridad(value)
	if err != nil {
		return Ticket{}, "", err
	}
	number = strings.TrimSpace(number)
	isInternal := strings.HasPrefix(number, "INT-")
	if isInternal {
		if actor.Role != auth.RoleDesarrollo {
			return Ticket{}, "", ErrForbidden
		}
	} else if actor.Role != auth.RoleSoporte {
		return Ticket{}, "", ErrForbidden
	}

	var principal repositories.Ticket
	notify := false
	err = s.tickets.DB().Transaction(func(tx *gorm.DB) error {
		query := tx.Clauses(clause.Locking{Strength: "UPDATE"})
		if isInternal {
			var internal repositories.InternalTicket
			result := query.Where("number = ?", number).Limit(1).Find(&internal)
			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected == 0 {
				return ErrTicketNotFound
			}
			if internal.State == repositories.StateCerrado {
				return ErrClosed
			}
			result = query.Where("id = ?", internal.TicketID).Limit(1).Find(&principal)
			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected == 0 {
				return ErrTicketNotFound
			}
		} else {
			result := query.Where("number = ?", number).Limit(1).Find(&principal)
			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected == 0 {
				return ErrTicketNotFound
			}
		}
		if principal.State == repositories.StateCerrado {
			return ErrClosed
		}
		if principal.Priority == priority {
			return nil
		}

		before := principal.Priority
		principal.Priority = priority
		if priority == repositories.PriorityCritical && principal.CriticalNotifiedAt == nil {
			now := s.now()
			principal.CriticalNotifiedAt = &now
			notify = true
		}
		if err := s.tickets.UpdateTicket(tx, principal); err != nil {
			return err
		}
		return s.conversation.AddHistory(tx, repositories.Anotar(
			repositories.DestinoDePrincipal(principal.ID), &actor.ID, "prioridad", before, priority, "",
		))
	})
	if err != nil {
		return Ticket{}, "", err
	}

	warning := ""
	if notify {
		requester, accountErr := s.accounts.ByID(principal.RequesterID)
		if accountErr != nil || s.avisarPrioridadCritica(principal, requester) != nil {
			warning = "tickets.priority.notificationFailed"
		}
	}
	detail, err := s.ByNumber(number, actor)
	if err != nil {
		return Ticket{}, warning, err
	}
	return detail.Ticket, warning, nil
}

// avisarPrioridadCritica informa una sola vez a todo Soporte y permite conocer el fallo del envío.
func (s *Service) avisarPrioridadCritica(ticket repositories.Ticket, requester auth.Account) error {
	team, err := s.accounts.ActiveByRole(auth.RoleSoporte)
	if err != nil {
		return err
	}
	to := make([]string, 0, len(team))
	for _, person := range team {
		if person.IsActive {
			to = append(to, person.Email)
		}
	}
	if len(to) == 0 {
		return nil
	}
	return s.mailer.Send(CorreoTicketPrioridadCritica, s.idiomaGlobal(), to, map[string]string{
		"numero": ticket.Number, "asunto": ticket.Subject, "solicitante": requester.FullName(),
		"prioridad": repositories.PriorityCritical, "enlace": s.enlaceDelTicket(ticket.Number),
	})
}
