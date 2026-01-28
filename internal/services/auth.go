package services

import (
	"Url-shortener/internal/models"
	"Url-shortener/internal/store"
	"crypto/rand"
	"encoding/hex"
	// "fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	userStore    store.UserStore
	sessionStore store.SessionStore
}

func NewAuthService(us store.UserStore, ss store.SessionStore) *AuthService {
	return &AuthService{
		userStore:    us,
		sessionStore: ss,
	}
}

func generateID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (a *AuthService) Register(username, password string) error {
	// hashiing the password using bcrypt , dont know how does it work really but we'll get to that
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// create / fill
	user := models.User{
		ID:           generateID(),
		Username:     username,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}
	//creating a user
	return a.userStore.Create(user)
}

func (a *AuthService) Login(username, password string) (string, error) {

	user, err := a.userStore.GetByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", err
	}

	sessoin := models.Session{
		SessionID: generateID(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(10 * time.Hour), // whatever the number
	}
	if err := a.sessionStore.Create(sessoin); err != nil {
		return "", err
	}
	return sessoin.SessionID, nil
}

func (a *AuthService) Logout(sessionID string) error {
	return a.sessionStore.Delete(sessionID)
}
