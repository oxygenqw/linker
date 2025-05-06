package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/services"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

type RequestsHandlerImpl struct {
	logger    *logger.Logger
	service   *services.Service
	templates map[string]*template.Template
}

func NewRequestsHandler(service *services.Service, logger *logger.Logger) *RequestsHandlerImpl {
	files := []string{
		"./web/pages/request_form.html",
	}

	templates := make(map[string]*template.Template)
	for _, file := range files {
		name := filepath.Base(file)
		templates[name] = template.Must(template.ParseFiles(file))
	}

	return &RequestsHandlerImpl{
		logger:    logger,
		service:   service,
		templates: templates,
	}
}

func (h *RequestsHandlerImpl) ToRequestForm(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: ToRequestForm] ", "URL: ", r.URL)

	data := map[string]any{
		"id":           params.ByName("id"),
		"role":         params.ByName("role"),
		"recipient_id": params.ByName("recipient_id"),
	}

	h.renderTemplate(w, "request_form.html", data)
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

func (h *RequestsHandlerImpl) renderTemplate(w http.ResponseWriter, tmplName string, data any) {
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
