package main

import (
	"Url-shortener/internal/models"
	"Url-shortener/internal/shortener"
	"Url-shortener/internal/store"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// handler , ednpoints
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var urlModel models.UrlData

	if err := json.NewDecoder(r.Body).Decode(&urlModel); err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	urlModel.ShortCode = shortener.GenerateShortID()
	urlModel.CreationTime = time.Now()

	store.Save(urlModel)
	fmt.Fprint(w, urlModel.ShortCode)
}
func ResolveHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[1:]
	// lookup using Get method/func
	LongUrl, ok := store.Get(slug)
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, LongUrl, 301)
}

func main() {

	fmt.Println("Url shortner server running ")
	http.HandleFunc("/shorten", ShortenHandler)
	http.HandleFunc("/", ResolveHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("error running the http server")
	}
}
