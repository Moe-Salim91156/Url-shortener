package models

import (
	"time"
)

type UrlData struct {
	LongUrl      string `json:"Url"`
	ShortCode    string
	CreationTime time.Time
}
