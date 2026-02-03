package store

import "Url-shortener/internal/models"

type PasteStore interface {
	Save(paste models.Paste) error
	Get(shortCode string) (*models.Paste, error)
	GetByOwner(ownerID string) ([]models.Paste, error)
	Delete(shortCode string) error
}

type SessionStore interface {
	Create(session models.Session) error
	Get(sessionID string) (*models.Session, error)
	Delete(sessionID string) error
}
type URLStore interface {
	Save(data models.UrlData) error
	Get(shortCode string) (*models.UrlData, error)
	GetByOwner(OwnerID string) ([]models.UrlData, error)
	Delete(shortCode string) error
}

type UserStore interface {
	Create(user models.User) error
	GetByUsername(username string) (*models.User, error)
	GetByID(id string) (*models.User, error)
}
