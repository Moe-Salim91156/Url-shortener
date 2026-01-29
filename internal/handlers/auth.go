package handlers

import (
	"Url-shortener/internal/services"
	"html/template"
	"net/http"
	"time"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(as *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: as}
}

func (h *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
		return
	}
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		sessionID, err := h.authService.Login(username, password)
		if err != nil {
			tmpl := template.Must(template.ParseFiles("templates/login.html"))
			tmpl.Execute(w, map[string]string{"Error": "Invalid credentials"})
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(10 * time.Hour),
		})
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func (h *AuthHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/register.html"))
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		err := h.authService.Register(username, password)
		if err != nil {
			tmpl := template.Must(template.ParseFiles("templates/register.html"))
			tmpl.Execute(w, map[string]string{"Error": err.Error()})
			return
		}
		sessionID, _ := h.authService.Login(username, password)
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(10 * time.Hour),
		})
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		h.authService.Logout(cookie.Value)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
