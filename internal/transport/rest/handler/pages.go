package handler

import (
	"net/http"

	"github.com/Oxygenss/linker/internal/renderer"
	"github.com/Oxygenss/linker/internal/services"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

type PagesHandlerImpl struct {
	logger   *logger.Logger
	service  *services.Service
	renderer *renderer.TemplateRenderer
}

func NewPagesHandler(service *services.Service, renderer *renderer.TemplateRenderer, logger *logger.Logger) *PagesHandlerImpl {
	return &PagesHandlerImpl{
		logger:   logger,
		service:  service,
		renderer: renderer,
	}
}

func (h *PagesHandlerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: Login] ", "URL: ", r.URL)

	data := map[string]any{
		"user_name":   params.ByName("user_name"),
		"telegram_id": params.ByName("telegram_id"),
	}

	h.renderer.RenderTemplate(w, "login.html", data)
}

func (h *PagesHandlerImpl) Home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//h.logger.Info("[H: Home] ", "URL: ", r.URL)

	id := params.ByName("id")
	role := params.ByName("role")
	h.logger.Infof("Home handler: id=%s, role=%s", id, role)
	
	data := map[string]any{
		"id":   params.ByName("id"),
		"role": params.ByName("role"),
	}

	h.renderer.RenderTemplate(w, "home.html", data)
}
