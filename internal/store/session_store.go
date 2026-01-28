package store

import (
	"Url-shortener/internal/models"
	"fmt"
	"time"
)

type SessionStore interface {
	Create(session models.Session) error
	Get(sessionID string) (*models.Session, error)
	Delete(sessionID string) error
}
type InMemorySessionStore struct {
	sessions map[string]models.Session
	// key is SESSION_ID
}

// same contracts/architecture as before
func NewInMemorySessionStore() *InMemorySessionStore {
	return &InMemorySessionStore{
		sessions: make(map[string]models.Session),
	}
}

func (S *InMemorySessionStore) Create(session models.Session) error {
	S.sessions[session.SessionID] = session
	return nil
}

func (S *InMemorySessionStore) Get(sessionID string) (*models.Session, error) {
	Session, exists := S.sessions[sessionID]

	// if time of session is expired , this is extra
	if !exists || time.Now().After(Session.CreatedAt) {
		return nil, fmt.Errorf("session not found OR Expired")
	}
	return &Session, nil
}

func (S *InMemorySessionStore) Delete(sessionID string) error {
	delete(S.sessions, sessionID)
	return nil
}
