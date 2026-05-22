package manager

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

const SessionCookieName = "session_id"

type SessionManager struct {
	persistence persistence.SessionPersistence
}

func NewSessionManager(persistence persistence.SessionPersistence) *SessionManager {
	return &SessionManager{persistence: persistence}
}

func (m *SessionManager) Save(ctx context.Context, session *persistence.Session) error {
	return m.persistence.Add(ctx, session)
}

func (m *SessionManager) Find(ctx context.Context, id string) (*persistence.Session, error) {
	return m.persistence.Get(ctx, id)
}

func (m *SessionManager) Remove(ctx context.Context, id string) error {
	return m.persistence.Delete(ctx, id)
}

func (m *SessionManager) Create(ctx context.Context, user entities.SessionUser) (*persistence.Session, error) {
	id, err := newSessionID()
	if err != nil {
		return nil, err
	}

	session := &persistence.Session{
		ID:   id,
		User: user,
	}

	if err := m.Save(ctx, session); err != nil {
		return nil, err
	}

	return session, nil
}

func newSessionID() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
