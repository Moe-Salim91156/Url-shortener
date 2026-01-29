package handlers

import (
	"Url-shortener/internal/middleware"
	"Url-shortener/internal/services"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
)

type PasteHandler struct {
	pasteService *services.PasteService
}

func NewPasteHandler(ps *services.PasteService) *PasteHandler {
	return &PasteHandler{
		pasteService: ps,
	}
}

func (h *PasteHandler) CreatePaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Content string `json:"content"`
		Title   string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userID := middleware.GetUserID(r)
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	shortCode, err := h.pasteService.CreatePaste(req.Content, req.Title, userID)
	if err != nil {
		http.Error(w, "Failed to create paste", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"shortCode": shortCode,
		"pasteUrl":  "http://localhost:8000/paste/" + shortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PasteHandler) ViewPaste(w http.ResponseWriter, r *http.Request) {
	shortCode := strings.TrimPrefix(r.URL.Path, "/paste/")

	paste, err := h.pasteService.GetPaste(shortCode)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/paste_view.html"))
	tmpl.Execute(w, paste)
}

func (h *PasteHandler) DeletePaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := middleware.GetUserID(r)
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	shortCode := strings.TrimPrefix(r.URL.Path, "/delete-paste/")

	if err := h.pasteService.DeletePaste(shortCode, userID); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
