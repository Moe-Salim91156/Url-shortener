package models

import (
	"time"
)

type UrlData struct {
	LongUrl      string    `json:"Url"`
	ShortCode    string    `json:"Short_Url"`
	CreationTime time.Time `json:"CreationTime"`
}
