package manager

import (
	"net/http"

	"github.com/v2code/autolog/internal/auth"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type SessionPrincipal struct {
	SessionID string
	User      entities.SessionUser
}

func (p *SessionPrincipal) Principal() *SessionPrincipal {
	return p
}

type SessionAuthManager struct {
	sessions   *SessionManager
	cookieName string
}

func NewSessionAuthManager(sessions *SessionManager) auth.AuthManager[*SessionPrincipal] {
	return &SessionAuthManager{
		sessions:   sessions,
		cookieName: SessionCookieName,
	}
}

func (m *SessionAuthManager) Authenticate(req *http.Request) (auth.Principal[*SessionPrincipal], error) {
	cookie, err := req.Cookie(m.cookieName)
	if err != nil {
		return nil, auth.ErrUnauthorized
	}

	session, err := m.sessions.Find(req.Context(), cookie.Value)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, auth.ErrSessionNotFound
	}

	return &SessionPrincipal{
		SessionID: session.ID,
		User:      session.User,
	}, nil
}
