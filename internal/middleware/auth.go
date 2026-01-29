package middleware

import (
	"Url-shortener/internal/services"
	"context"
	"net/http"
)

// contextKey is a custom type for context keys
type contextKey string

// UserIDKey is the key for storing user ID in context
const UserIDKey contextKey = "userID"

// RequireAuth wraps handlers to ensure user is logged in
func RequireAuth(authService *services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_id")
			if err != nil {
				http.Redirect(w, r, "/login", 303)
			}
			user, err := authService.GetUserFromSession(cookie.Value)
			if err != nil {
				http.Redirect(w, r, "/login", 303)
			}
			ctx := context.WithValue(r.Context(), UserIDKey, user.ID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserID(r *http.Request) string {
	userID, _ := r.Context().Value(UserIDKey).(string)
	return userID
}
