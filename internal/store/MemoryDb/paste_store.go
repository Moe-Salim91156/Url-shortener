package store

import (
	"Url-shortener/internal/models"
	"fmt"
)

type InMemoryPasteStore struct {
	pastes map[string]models.Paste
}

func NewInMemoryPasteStore() *InMemoryPasteStore {
	return &InMemoryPasteStore{
		pastes: make(map[string]models.Paste),
	}
}

func (s *InMemoryPasteStore) Save(paste models.Paste) error {
	s.pastes[paste.ShortCode] = paste
	return nil
}

func (s *InMemoryPasteStore) Get(shortCode string) (*models.Paste, error) {
	paste, ok := s.pastes[shortCode]
	if !ok {
		return nil, fmt.Errorf("paste not found")
	}
	return &paste, nil
}

func (s *InMemoryPasteStore) GetByOwner(ownerID string) ([]models.Paste, error) {
	var result []models.Paste
	for _, paste := range s.pastes {
		if paste.OwnerID == ownerID {
			result = append(result, paste)
		}
	}
	return result, nil
}

func (s *InMemoryPasteStore) Delete(shortCode string) error {
	delete(s.pastes, shortCode)
	return nil
}
