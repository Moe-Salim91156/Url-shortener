package handlers

import (
	"Url-shortener/internal/middleware"
	"Url-shortener/internal/models"
	"Url-shortener/internal/services"
	"Url-shortener/internal/shortener"
	"Url-shortener/internal/store"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type URLHandler struct {
	store      store.URLStore
	urlService *services.URLService
}

func NewURLHandler(s store.URLStore, Us *services.URLService) *URLHandler {
	return &URLHandler{store: s,
		urlService: Us}
}

func (h *URLHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	var urlModel models.UrlData

	if err := json.NewDecoder(r.Body).Decode(&urlModel); err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}
	userID := middleware.GetUserID(r)
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	urlModel.ShortCode = shortener.GenerateShortID()
	urlModel.OwnerID = userID
	urlModel.CreationTime = time.Now()

	if err := h.store.Save(urlModel); err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"shortCode": urlModel.ShortCode,
		"shortUrl":  "http://localhost:8000/" + urlModel.ShortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *URLHandler) Resolve(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[1:]

	if slug == "" || slug == "login" || slug == "register" || slug == "dashboard" ||
		strings.HasPrefix(slug, "shorten") || strings.HasPrefix(slug, "delete") || slug == "logout" {
		return
	}

	urlData, err := h.store.Get(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, urlData.LongUrl, 301)
}

func (h *URLHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := middleware.GetUserID(r)
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	shortCode := strings.TrimPrefix(r.URL.Path, "/delete/")

	if err := h.urlService.DeleteURL(shortCode, userID); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
