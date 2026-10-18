package services

import (
	"strings"

	"catalina-support/backend/shared/auth"
)

// ImproveDraft autoriza el editor, obtiene el destinatario desde la base y delega únicamente el
// contexto mínimo. No persiste ni el borrador ni el resultado.
func (s *Service) ImproveDraft(number, editor, draft, previousSuggestion string, actor auth.Identity) (WritingSuggestion, error) {
	if !actor.Can(auth.RoleSoporte, auth.RoleDesarrollo) {
		return WritingSuggestion{}, ErrForbidden
	}
	borrador := strings.TrimSpace(draft)
	if borrador == "" {
		return WritingSuggestion{}, ErrWritingDraftRequired
	}
	if editor != "description" && editor != "comment" {
		return WritingSuggestion{}, ErrWritingEditorInvalid
	}
	if s.writing == nil || s.accounts == nil || s.language == nil {
		return WritingSuggestion{}, ErrWritingUnavailable
	}

	principal, interno, _, err := s.porNumero(number)
	if err != nil {
		return WritingSuggestion{}, err
	}
	detalle, err := s.ticketDe(number, actor)
	if err != nil {
		return WritingSuggestion{}, err
	}
	if editor == "description" {
		if interno || !puedeEditar(actor, detalle) {
			return WritingSuggestion{}, ErrForbidden
		}
	} else if !puedeComentar(actor, detalle) {
		return WritingSuggestion{}, ErrForbidden
	}

	destinatarioID := principal.RequesterID
	tipo := "main"
	if interno {
		internoDB, err := s.tickets.InternalByNumber(strings.TrimSpace(number))
		if err != nil {
			return WritingSuggestion{}, traducirTicket(err)
		}
		destinatarioID = internoDB.CreatedByID
		tipo = "internal"
	}
	destinatario, err := s.accounts.ByID(destinatarioID)
	if err != nil {
		return WritingSuggestion{}, err
	}
	idioma, err := s.language.Language()
	if err != nil {
		return WritingSuggestion{}, err
	}

	return s.writing.ImproveDraft(
		borrador, strings.TrimSpace(previousSuggestion), destinatario.FullName(), tipo, idioma,
	)
}
