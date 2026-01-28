package models

import (
	"time"
)

type UrlData struct {
	LongUrl   string `json:"Url"`
	ShortCode string `json:"Short_Url"`
	// there should be the owner ID , to specify which user owns which URLS
	// OwnerID string
	CreationTime time.Time `json:"CreationTime"`
}
