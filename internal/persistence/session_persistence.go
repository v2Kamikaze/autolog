package persistence

type Session struct {
	ID   string
	Data any
}

type SessionPersistence interface {
	Add(session *Session) error
	Get(id string) (*Session, error)
	Delete(id string) error
}

type inMemorySessionPersistence struct {
	cache map[string]*Session
}

func NewSessionPersistence() SessionPersistence {
	return &inMemorySessionPersistence{
		cache: make(map[string]*Session),
	}
}

func (s *inMemorySessionPersistence) Add(session *Session) error {
	s.cache[session.ID] = session
	return nil
}

func (s *inMemorySessionPersistence) Get(id string) (*Session, error) {
	if v, ok := s.cache[id]; ok {
		return v, nil
	}
	return nil, nil
}

func (s *inMemorySessionPersistence) Delete(sessionID string) error {
	delete(s.cache, sessionID)
	return nil
}
