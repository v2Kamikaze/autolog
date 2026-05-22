package persistence

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/v2code/autolog/internal/persistence/entities"
)

const defaultSessionTTL = 24 * time.Hour

const sessionKeyPrefix = "session:"

func BuildSessionKey(id string) string {
	return sessionKeyPrefix + id
}

type Session struct {
	ID   string
	User entities.SessionUser
}

type SessionPersistence interface {
	Add(ctx context.Context, session *Session) error
	Get(ctx context.Context, id string) (*Session, error)
	Delete(ctx context.Context, id string) error
}

type sessionPersistence struct {
	client *redis.Client
	ttl    time.Duration
}

func NewSessionPersistence(client *redis.Client) SessionPersistence {
	return &sessionPersistence{
		client: client,
		ttl:    defaultSessionTTL,
	}
}

func (s *sessionPersistence) Add(ctx context.Context, session *Session) error {
	payload, err := json.Marshal(session)
	if err != nil {
		return err
	}

	key := BuildSessionKey(session.ID)
	return s.client.Set(ctx, key, payload, s.ttl).Err()
}

func (s *sessionPersistence) Get(ctx context.Context, id string) (*Session, error) {
	key := BuildSessionKey(id)
	data, err := s.client.Get(ctx, key).Bytes()

	if err != nil {
		return nil, err
	}

	var session Session
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *sessionPersistence) Delete(ctx context.Context, id string) error {
	key := BuildSessionKey(id)
	return s.client.Del(ctx, key).Err()
}
