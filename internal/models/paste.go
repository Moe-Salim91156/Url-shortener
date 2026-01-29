package models

import (
	"time"
)

type Paste struct {
	Content      string    `json:"content"`
	ShortCode    string    `json:"short_code"`
	Title        string    `json:"title"`
	OwnerID      string    `json:"owner_id"`
	CreationTime time.Time `json:"creation_time"`
}
