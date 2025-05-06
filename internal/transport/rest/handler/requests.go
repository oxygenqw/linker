package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/renderer"
	"github.com/Oxygenss/linker/internal/services"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

type RequestsHandlerImpl struct {
	logger   *logger.Logger
	service  *services.Service
	renderer *renderer.TemplateRenderer
}

func NewRequestsHandler(service *services.Service, renderer *renderer.TemplateRenderer, logger *logger.Logger) *RequestsHandlerImpl {
	return &RequestsHandlerImpl{
		logger:   logger,
		service:  service,
		renderer: renderer,
	}
}

func (h *RequestsHandlerImpl) ToRequestForm(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: ToRequestForm] ", "URL: ", r.URL)

	data := map[string]any{
		"id":           params.ByName("id"),
		"role":         params.ByName("role"),
		"recipient_id": params.ByName("recipient_id"),
	}

	h.renderer.RenderTemplate(w, "request_form.html", data)
}

func (h *RequestsHandlerImpl) RequestToStudent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: RequestToStudent] ", "URL: ", r.URL)
	id := params.ByName("id")
	recipientID := params.ByName("recipient_id")

	var requestData models.Request

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	if requestData.Message == "" {
		http.Error(w, "Сообщение не может быть пустым", http.StatusBadRequest)
		return
	}

	err := h.service.RequestService.CreateRequest(id, recipientID, requestData)
	if err != nil {
		http.Error(w, "Ошибка сервера: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Запрос успешно отправлен",
	})
}
