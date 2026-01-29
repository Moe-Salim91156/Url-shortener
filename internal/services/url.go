package services

// decide how to use the store/interfaces , business logic , BRain

import (
	"Url-shortener/internal/models"
	"Url-shortener/internal/shortener"
	"Url-shortener/internal/store"
	"fmt"
	"time"
)

type URLService struct {
	URLStore store.URLStore
}

func NewURLService(us store.URLStore) *URLService {
	return &URLService{
		URLStore: us,
	}
}

// create URL , plus USER now
func (s *URLService) CreateShortURL(LongURL string, ownerID string) (string, error) {
	// createa the short url using genereate short url
	// add the url Data
	// save ->
	// return the generaetd URL (Short)
	shortCode := shortener.GenerateShortID()
	URLData := models.UrlData{
		LongUrl:      LongURL,
		ShortCode:    shortCode,
		OwnerID:      ownerID,
		CreationTime: time.Now(),
	}
	err := s.URLStore.Save(URLData)
	if err != nil {
		return "", fmt.Errorf("saving URl Data struct failed")
	}
	return shortCode, nil
}
func (s *URLService) GetUserURLs(userID string) ([]models.UrlData, error) {

	return s.URLStore.GetByOwner(userID)
}

func (s *URLService) DeleteURL(shortCode, userID string) error {
	urlData, err := s.URLStore.Get(shortCode)
	if err != nil {
		return fmt.Errorf("failed to get URL in DELETE URL")
	}
	if urlData.OwnerID != userID {
		return fmt.Errorf("You are not authorized to do this , stupid !")
	}
	return s.URLStore.Delete(shortCode)
}

func (s *URLService) ResolveShortCode(shortCode string) (string, error) {
	urlData, err := s.URLStore.Get(shortCode)
	if err != nil {
		return "", fmt.Errorf("Error In resolving Short code")
	}
	return urlData.LongUrl, nil
}
