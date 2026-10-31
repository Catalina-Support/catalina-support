package services

import (
	"errors"
	"strconv"
	"strings"

	"github.com/juanky-estevez/go-logs"

	"catalina-support/backend/modules/users/repositories"
	"catalina-support/backend/shared/auth"
)

// EnsureRequester devuelve la cuenta del solicitante y dice si acaba de crearla.
//
// Lo usa el alta de un ticket cuando Soporte escribe el correo de una persona que todavía no tiene
// cuenta. La cuenta local nace **sin contraseña y sin correo de alta**: identifica y avisa, pero no
// da acceso. Con AD se crea o vincula desde el directorio; con Keycloak no se puede buscar y se
// rechaza. Las reglas están en docs/modules/users.md, sección 18.
func (s *Service) EnsureRequester(email, name, lastName string, actor auth.Identity) (auth.Account, bool, error) {
	email = strings.TrimSpace(email)
	if !EmailIsValid(email) {
		return auth.Account{}, false, auth.ErrRequesterEmailInvalid
	}

	account, err := s.ByEmail(email)
	if err == nil {
		if !account.IsActive {
			return auth.Account{}, false, auth.ErrRequesterInactive
		}

		return account, false, nil
	}
	if !errors.Is(err, auth.ErrAccountNotFound) {
		return auth.Account{}, false, err
	}

	if s.access == nil {
		return auth.Account{}, false, auth.ErrRequesterDirectoryDown
	}
	access, err := s.access.Access()
	if err != nil {
		return auth.Account{}, false, err
	}

	switch access.Method {
	case auth.OriginLocal:
		return s.createLocalRequester(email, name, lastName)
	case auth.OriginAD:
		return s.createRequesterFromDirectory(email, actor)
	case auth.OriginKeycloak:
		return auth.Account{}, false, auth.ErrRequesterAccountExists
	default:
		return auth.Account{}, false, ErrOriginUnknown
	}
}

// createLocalRequester crea la cuenta mínima que identifica al solicitante. No manda correo ni
// establece contraseña: no le da acceso a la plataforma.
func (s *Service) createLocalRequester(email, name, lastName string) (auth.Account, bool, error) {
	name = strings.TrimSpace(name)
	lastName = strings.TrimSpace(lastName)
	if name == "" || lastName == "" {
		return auth.Account{}, false, auth.ErrRequesterNameRequired
	}

	nueva := &repositories.User{
		Name:     name,
		LastName: lastName,
		Email:    email,
		Role:     auth.RoleUsuario,
		Origin:   auth.OriginLocal,
		IsActive: true,
	}

	if err := s.repo.Create(nueva); err != nil {
		// Si otra petición ganó la carrera, la cuenta ya está y el ticket puede usarla.
		if existente, buscarErr := s.ByEmail(email); buscarErr == nil {
			if !existente.IsActive {
				return auth.Account{}, false, auth.ErrRequesterInactive
			}

			return existente, false, nil
		}

		return auth.Account{}, false, err
	}

	logs.LogSuccess("cuenta de solicitante creada desde el alta de un ticket: la #" + strconv.FormatInt(nueva.ID, 10))

	return nueva.Account(), true, nil
}

// createRequesterFromDirectory busca a la persona por su correo y la crea o vincula con los datos
// del directorio. Si no está, el ticket no se abre.
func (s *Service) createRequesterFromDirectory(email string, actor auth.Identity) (auth.Account, bool, error) {
	buscador, ok := s.directory.(DirectorySearch)
	if !ok || !buscador.DirectoryEnabled() {
		return auth.Account{}, false, auth.ErrRequesterDirectoryDown
	}

	personas, err := buscador.SearchDirectory(email, true, 20)
	if err != nil {
		return auth.Account{}, false, auth.ErrRequesterDirectoryDown
	}

	for i := range personas {
		if strings.EqualFold(personas[i].Email, email) || strings.EqualFold(personas[i].Username, email) {
			cuenta, err := s.CreateFromDirectory(personas[i].Username, auth.RoleUsuario, false, false, actor)
			if err != nil {
				if errors.Is(err, ErrDirectoryUnavailable) {
					return auth.Account{}, false, auth.ErrRequesterDirectoryDown
				}
				if errors.Is(err, ErrDirectoryNotFound) {
					return auth.Account{}, false, auth.ErrRequesterNotInDirectory
				}

				return auth.Account{}, false, err
			}

			return cuenta, true, nil
		}
	}

	return auth.Account{}, false, auth.ErrRequesterNotInDirectory
}
