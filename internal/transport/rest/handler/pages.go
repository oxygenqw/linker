package handler

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/Oxygenss/linker/internal/services"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

type PagesHandlerImpl struct {
	logger    *logger.Logger
	service   *services.Service
	templates map[string]*template.Template
}

func NewPagesHandler(service *services.Service, logger *logger.Logger) *PagesHandlerImpl {
	files := []string{
		"./web/pages/home.html",
		"./web/pages/login.html",
	}

	templates := make(map[string]*template.Template)
	for _, file := range files {
		name := filepath.Base(file)
		templates[name] = template.Must(template.ParseFiles(file))
	}

	return &PagesHandlerImpl{
		logger:    logger,
		service:   service,
		templates: templates,
	}
}

// Рендерит login.html и передает туда user_name
// @router GET /login/:user_name/:telegram_id
func (h *PagesHandlerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: Login] ", "URL: ", r.URL)

	data := map[string]any{
		"user_name":   params.ByName("user_name"),
		"telegram_id": params.ByName("telegram_id"),
	}

	h.renderTemplate(w, "login.html", data)
}

// Рендерит home.html и передает туда id и role
// @router GET /home/:id/:role
func (h *PagesHandlerImpl) Home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: Home] ", "URL: ", r.URL)

	data := map[string]any{
		"id":   params.ByName("id"),
		"role": params.ByName("role"),
	}

	h.renderTemplate(w, "home.html", data)
}

func (h *PagesHandlerImpl) renderTemplate(w http.ResponseWriter, tmplName string, data any) {
	tmpl, ok := h.templates[tmplName]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, data); err != nil {
		h.logger.Error("Failed to render template", "error", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
