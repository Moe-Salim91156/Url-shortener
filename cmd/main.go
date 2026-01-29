package main

import (
	"Url-shortener/internal/handlers"
	"Url-shortener/internal/middleware"
	"Url-shortener/internal/services"
	"Url-shortener/internal/store"
	"fmt"
	"log"
	"net/http"
)

func main() {
	userStore := store.NewInMemoryUserStore()
	sessionStore := store.NewInMemorySessionStore()
	urlStore := store.NewInMemoryURLStorage()

	authService := services.NewAuthService(userStore, sessionStore)
	urlService := services.NewURLService(urlStore)

	authHandler := handlers.NewAuthHandler(authService)
	urlHandler := handlers.NewURLHandler(urlStore, urlService)
	dashboardHandler := handlers.NewDashboardHandler(urlService)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		urlHandler.Resolve(w, r)
	})
	http.HandleFunc("/login", authHandler.LoginPage)
	http.HandleFunc("/register", authHandler.RegisterPage)
	http.HandleFunc("/logout", authHandler.Logout)

	requireAuth := middleware.RequireAuth(authService)
	http.Handle("/dashboard", requireAuth(http.HandlerFunc(dashboardHandler.ShowDashboard)))
	http.Handle("/shorten", requireAuth(http.HandlerFunc(urlHandler.Shorten)))
	http.Handle("/delete/", requireAuth(http.HandlerFunc(urlHandler.Delete)))

	fmt.Println("🚀 Server starting on http://localhost:8000")
	fmt.Println("📝 Register at: http://localhost:8000/register")
	fmt.Println("🔑 Login at: http://localhost:8000/login")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
