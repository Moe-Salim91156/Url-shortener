package store

import "Url-shortener/internal/models"
import "fmt"

// treating this as a sepereate storage service
// means that if we would add actual DB , the interface will take care of the rest of the code, no changing reqiured(much/collateral)
type URLStore interface {
	Save(data models.UrlData) error
	Get(shortCode string) (*models.UrlData, error)
	GetByOwner(OwnerID string) ([]models.UrlData, error)
	Delete(shortCode string) error
}

type InMemoryStorage struct {
	urls map[string]models.UrlData
}

// data{}
// init new struct
func NewInMemoryURLStorage() *InMemoryStorage {
	return &InMemoryStorage{
		urls: make(map[string]models.UrlData),
	}
}

// method tied to structs now istead of generals
func (m *InMemoryStorage) Save(data models.UrlData) error {
	m.urls[data.ShortCode] = data
	return nil
}

func (m *InMemoryStorage) Get(shortCode string) (*models.UrlData, error) {
	data, ok := m.urls[shortCode]
	if !ok {
		return nil, fmt.Errorf("URL NOT FOUND")
	}
	return &data, nil
}

func (m *InMemoryStorage) GetByOwner(OwnerID string) ([]models.UrlData, error) {
	// loop through map , if USERID matches url.owenerID
	var result []models.UrlData
	for _, url := range m.urls {
		if url.OwnerID == OwnerID {
			result = append(result, url)
			// append to result to return the slice
		}
	}
	return result, nil
}
func (m *InMemoryStorage) Delete(shortCode string) error {
	delete(m.urls, shortCode)
	return nil
}
