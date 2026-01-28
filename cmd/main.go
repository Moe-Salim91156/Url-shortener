package main

import (
	"Url-shortener/internal/handlers"
	"Url-shortener/internal/store"
	"fmt"
	"log"
	"net/http"
)

/*
 * now when propagating the store, it could be anything, code wont change , just in main
 because of the interface
*/

func main() {

	// here we could switch later and make the url store a DB ,
	//and thats only the change that would happen
	UrlStore := store.NewInMemoryStorage()
	UrlHandler := handlers.NewURLHandler(UrlStore)

	fmt.Println("Url shortner server running ")
	http.HandleFunc("/shorten", UrlHandler.Shorten)
	http.HandleFunc("/", UrlHandler.Resolve)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("error running the http server")
	}
}
