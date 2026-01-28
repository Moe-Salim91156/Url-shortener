package store

// create user , get user -> by username , by ID
import (
	"Url-shortener/internal/models"
	"fmt"
)

// when user register , program must store its data like the url_store manner
type UserStore interface {
	Create(user models.User) error
	GetByUsername(username string) (*models.User, error)
	GetByID(id string) (*models.User, error)
}

type InMemoryUserStore struct {
	UsersByName map[string]models.User //for retrieving by username
	UsersByID   map[string]models.User // for retrieving by ID if needed by session
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		UsersByName: make(map[string]models.User),
		UsersByID:   make(map[string]models.User),
	}
}

func (U *InMemoryUserStore) Create(user models.User) error {
	if _, exists := U.UsersByName[user.Username]; exists {
		// checking if user exits by getting username
		return fmt.Errorf("User ALready exits")
	}
	// now store in both maps
	U.UsersByID[user.ID] = user
	U.UsersByName[user.Username] = user
	return nil
}

func (U *InMemoryUserStore) GetByID(id string) (*models.User, error) {
	user, ok := U.UsersByID[id]
	if !ok {
		return nil, fmt.Errorf("ID for User Not Found")
	}
	return &user, nil
}

func (U *InMemoryUserStore) GetByUsername(username string) (*models.User, error) {
	user, ok := U.UsersByName[username]
	if !ok {
		return nil, fmt.Errorf("Username invalid")
	}
	return &user, nil
}
