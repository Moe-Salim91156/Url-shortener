package store

import "Url-shortener/internal/models"
import "fmt"

var urls = make(map[string]models.UrlData)

func Save(data models.UrlData) {
	urls[data.ShortCode] = data
	//append into the map
}

func Get(shortCode string) (string, bool) {
	fmt.Printf("Looking for: '%s' in map with %d entries\n", shortCode, len(urls))
	data, ok := urls[shortCode]
	// lookup?
	return data.LongUrl, ok
}
