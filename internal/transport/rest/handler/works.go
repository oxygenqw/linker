package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/renderer"
	"github.com/Oxygenss/linker/internal/services"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type WorksHandlerImpl struct {
	logger   *logger.Logger
	service  *services.Service
	renderer *renderer.TemplateRenderer
}

func NewWorksHandler(service *services.Service, renderer *renderer.TemplateRenderer, logger *logger.Logger) *WorksHandlerImpl {
	return &WorksHandlerImpl{
		logger:   logger,
		service:  service,
		renderer: renderer,
	}
}

func (h *WorksHandlerImpl) ToForm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: ToRequestForm] ", "URL: ", r.URL)

	userID := ps.ByName("user_id")
	role, err := h.service.UserService.GetRoleByID(userID)
	if err != nil {

	}

	data := map[string]any{
		"user_id": userID,
		"role":    role,
	}

	h.renderer.RenderTemplate(w, "create_work_form.html", data)
}

func (h *WorksHandlerImpl) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: CreateWork] ", "URL: ", r.URL)

	userIDStr := ps.ByName("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, `{"message":"Некорректный user_id"}`, http.StatusBadRequest)
		return
	}

	var work models.Work
	if err := json.NewDecoder(r.Body).Decode(&work); err != nil {
		http.Error(w, `{"message":"Некорректный формат запроса"}`, http.StatusBadRequest)
		return
	}

	work.UserID = userID
	if work.ID == uuid.Nil {
		work.ID = uuid.New()
	}

	if err := h.service.WorkService.Create(work); err != nil {
		h.logger.Error("Ошибка создания работы: ", err)
		http.Error(w, `{"message":"Ошибка сервера"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(work)
}

func (h *WorksHandlerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	workIDStr := ps.ByName("id")
	h.logger.Info("Delete work request, ID:", workIDStr)

	workID, err := uuid.Parse(workIDStr)
	if err != nil {
		http.Error(w, `{"message":"Некорректный workId"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.WorkService.Delete(workID); err != nil {
		h.logger.Error("Ошибка удаления работы: ", err)
		http.Error(w, `{"message":"Ошибка сервера"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Работа успешно удалена"}`))
}
