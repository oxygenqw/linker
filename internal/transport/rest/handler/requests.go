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
		"sender_id":    params.ByName("sender_id"),
		"recipient_id": params.ByName("recipient_id"),
	}

	h.renderer.RenderTemplate(w, "request_form.html", data)
}

func (h *RequestsHandlerImpl) RequestToUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: RequestToStudent] ", "URL: ", r.URL)
	senderID := params.ByName("sender_id")
	recipientID := params.ByName("recipient_id")

	var requestData models.Message

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	err := h.service.RequestService.SendRequest(senderID, recipientID, requestData)
	if err != nil {
		http.Error(w, "Ошибка сервера: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Запрос успешно отправлен",
	})
}
