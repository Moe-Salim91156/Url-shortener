package store

import "Url-shortener/internal/models"
import "fmt"

// treating this as a sepereate storage service
// means that if we would add actual DB , the interface will take care of the rest of the code, no changing reqiured(much/collateral)
type URLStore interface {
	Save(data models.UrlData) error
	Get(shortCode string) (string, error)
}

type InMemoryStorage struct {
	urls map[string]models.UrlData
}

// data{}
// init new struct
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		urls: make(map[string]models.UrlData),
	}
}

// method tied to structs now istead of generals
func (m *InMemoryStorage) Save(data models.UrlData) error {
	m.urls[data.ShortCode] = data
	return nil
}

func (m *InMemoryStorage) Get(shortCode string) (string, error) {
	data, ok := m.urls[shortCode]
	if !ok {
		return "", fmt.Errorf("URL NOT FOUND")
	}
	return data.LongUrl, nil
}
