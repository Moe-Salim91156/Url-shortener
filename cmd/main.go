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

/*
 * now when propagating the store, it could be anything, code wont change , just in main
 because of the interface
*/

// this is package level var
var UrlStore store.URLStore

// handler , ednpoints
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var urlModel models.UrlData

	if err := json.NewDecoder(r.Body).Decode(&urlModel); err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	urlModel.ShortCode = shortener.GenerateShortID()
	urlModel.CreationTime = time.Now()

	UrlStore.Save(urlModel)
	fmt.Fprint(w, urlModel.ShortCode)
}
func ResolveHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[1:]
	// lookup using Get method/func
	LongUrl, err := UrlStore.Get(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, LongUrl, 301)
}

func main() {

	UrlStore = store.NewInMemoryStorage()
	fmt.Println("Url shortner server running ")
	http.HandleFunc("/shorten", ShortenHandler)
	http.HandleFunc("/", ResolveHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("error running the http server")
	}
}
