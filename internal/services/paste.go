package services

import (
	"Url-shortener/internal/models"
	"Url-shortener/internal/shortener"
	"Url-shortener/internal/store"
	"fmt"
	"time"
)

type PasteService struct {
	PasteStore store.PasteStore
}

func NewPasteService(ps store.PasteStore) *PasteService {
	return &PasteService{
		PasteStore: ps,
	}
}

func (s *PasteService) CreatePaste(content, title, ownerID string) (string, error) {
	shortCode := shortener.GenerateShortID()
	paste := models.Paste{
		Content:      content,
		Title:        title,
		ShortCode:    shortCode,
		OwnerID:      ownerID,
		CreationTime: time.Now(),
	}
	err := s.PasteStore.Save(paste)
	if err != nil {
		return "", fmt.Errorf("failed to save paste")
	}
	return shortCode, nil
}

func (s *PasteService) GetUserPastes(userID string) ([]models.Paste, error) {
	return s.PasteStore.GetByOwner(userID)
}

func (s *PasteService) DeletePaste(shortCode, userID string) error {
	paste, err := s.PasteStore.Get(shortCode)
	if err != nil {
		return fmt.Errorf("paste not found")
	}
	if paste.OwnerID != userID {
		return fmt.Errorf("unauthorized to delete this paste")
	}
	return s.PasteStore.Delete(shortCode)
}

func (s *PasteService) GetPaste(shortCode string) (*models.Paste, error) {
	return s.PasteStore.Get(shortCode)
}
