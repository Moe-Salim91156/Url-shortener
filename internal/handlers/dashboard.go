package handlers

import (
	"Url-shortener/internal/middleware"
	"Url-shortener/internal/services"
	"html/template"
	"net/http"
)

type DashboardHandler struct {
	urlService *services.URLService
}

func NewDashboardHandler(us *services.URLService) *DashboardHandler {
	return &DashboardHandler{urlService: us}
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
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	data := map[string]interface{}{
		"URLs":  urls,
		"Count": len(urls),
	}
	tmpl.Execute(w, data)
}
