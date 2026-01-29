package main

import (
	"Url-shortener/internal/handlers"
	"Url-shortener/internal/services"
	"Url-shortener/internal/store"
	"fmt"
	// "log"
	"net/http"
)

/*
* now when propagating the store, it could be anything, code wont change , just in main
because of the interface
*/
func main() {
	// Auth test (keep this)
	userStore := store.NewInMemoryUserStore()
	sessionStore := store.NewInMemorySessionStore()
	authService := services.NewAuthService(userStore, sessionStore)

	err := authService.Register("alice", "password123")
	fmt.Println("Register error:", err)

	sessionID, err := authService.Login("alice", "password123")
	fmt.Println("Session ID:", sessionID)
	fmt.Println("Login error:", err)

	// NEW: URL Service test
	fmt.Println("\n=== Testing URL Service ===")
	urlStore := store.NewInMemoryURLStorage()
	urlService := services.NewURLService(urlStore)

	// Test 1: Create URL
	code, err := urlService.CreateShortURL("https://google.com", "user_alice")
	fmt.Printf("Created: %s, Error: %v\n", code, err)

	// Test 2: Get user's URLs
	urls, _ := urlService.GetUserURLs("user_alice")
	fmt.Printf("Alice has %d URLs\n", len(urls))

	// Test 3: Unauthorized delete (should fail)
	err = urlService.DeleteURL(code, "user_bob")
	fmt.Printf("Bob delete (should fail): %v\n", err)

	// Test 4: Authorized delete (should work)
	err = urlService.DeleteURL(code, "user_alice")
	fmt.Printf("Alice delete (should work): %v\n", err)

	// Test 5: Check it's gone
	urls, _ = urlService.GetUserURLs("user_alice")
	fmt.Printf("Alice now has %d URLs\n", len(urls))

	// Keep server running
	fmt.Println("\n=== Starting Server ===")
	urlHandler := handlers.NewURLHandler(urlStore)
	http.HandleFunc("/shorten", urlHandler.Shorten)
	http.HandleFunc("/", urlHandler.Resolve)
	http.ListenAndServe(":8000", nil)
}
