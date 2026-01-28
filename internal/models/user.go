package models

import "time"

type User struct {
	ID           string
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}

// so we gonna create the user , to be specific for each user.
// must own its own URLS

// func NewUser() *User {
// 	return &User{
// 		ID:           "",
// 		Username:     "",
// 		PasswordHash: "",
// 		CreatedAt:    time.Now(),
// 	}
// }
