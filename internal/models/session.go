package models

import "time"

type Session struct {
	SessionID string
	UserID    string
	CreatedAt time.Time
}
