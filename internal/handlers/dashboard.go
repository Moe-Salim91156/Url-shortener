package handlers

import (
	"Url-shortener/internal/middleware"
	"Url-shortener/internal/services"
	"html/template"
	"net/http"
)

type DashboardHandler struct {
	urlService   *services.URLService
	pasteService *services.PasteService
}

func NewDashboardHandler(us *services.URLService, ps *services.PasteService) *DashboardHandler {
	return &DashboardHandler{
		urlService:   us,
		pasteService: ps,
	}
}

func (h *DashboardHandler) ShowDashboard(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	if userID == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	urls, err := h.urlService.GetUserURLs(userID)
	if err != nil {
		http.Error(w, "Failed to fetch URLs", http.StatusInternalServerError)
		return
	}

	pastes, err := h.pasteService.GetUserPastes(userID)
	if err != nil {
		http.Error(w, "Failed to fetch pastes", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	data := map[string]interface{}{
		"URLs":       urls,
		"Count":      len(urls),
		"Pastes":     pastes,
		"PasteCount": len(pastes),
	}
	tmpl.Execute(w, data)
}
