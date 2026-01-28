package handlers

import (
	"Url-shortener/internal/models"
	"Url-shortener/internal/shortener"
	"Url-shortener/internal/store"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type URLHandler struct {
	store store.URLStore
}

func NewURLHandler(s store.URLStore) *URLHandler {
	return &URLHandler{store: s}
}

func (h *URLHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	var urlModel models.UrlData

	if err := json.NewDecoder(r.Body).Decode(&urlModel); err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	urlModel.ShortCode = shortener.GenerateShortID()
	urlModel.CreationTime = time.Now()

	h.store.Save(urlModel)
	fmt.Fprint(w, urlModel.ShortCode)
}

func (h *URLHandler) Resolve(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[1:]

	LongUrl, err := h.store.Get(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, LongUrl, 301)
}
