package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/service"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type Redirect interface {
	NewUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Input(w http.ResponseWriter, r *http.Request)
}

type RedirectHandler struct {
	logger  *logger.Logger
	service *service.Service
}

func NewRedirectHandler(service *service.Service, logger *logger.Logger) *RedirectHandler {
	return &RedirectHandler{
		logger:  logger,
		service: service,
	}
}

// ...
// @router GET /
func (h *RedirectHandler) Input(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("[H: Input] ", "URL: ", r.URL)

	telegramIDStr := r.URL.Query().Get("telegram_id")
	telegramID, err := strconv.ParseInt(telegramIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Telegram ID", http.StatusBadRequest)
		return
	}

	role, err := h.service.User.GetRole(telegramID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	switch role {
	case "":
		userName := r.URL.Query().Get("user_name")
		http.Redirect(w, r, fmt.Sprintf("/login/%s/%s", userName, telegramIDStr), http.StatusFound)
	case "student":
		student, err := h.service.Student.GetByTelegramID(telegramID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/home/%s/%s", student.ID, role), http.StatusFound)
		return
	case "teacher":
		teacher, err := h.service.Student.GetByTelegramID(telegramID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/home/%s/%s", teacher.ID, role), http.StatusFound)
		return
	}
}

// После заполнения первоначальных данных парсим форму и переходим в профиль
// @router POST /users/:telegram_id
func (h *RedirectHandler) NewUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: NewUser] ", "URL: ", r.URL)

	telegramIDStr := ps.ByName("telegram_id")
	telegramID, err := strconv.ParseInt(telegramIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Telegram ID", http.StatusBadRequest)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	role := r.FormValue("role")

	var id uuid.UUID

	switch role {
	case "student":
		student := models.Student{
			TelegramID: telegramID,
			FirstName:  r.FormValue("first_name"),
			LastName:   r.FormValue("last_name"),
			MiddleName: r.FormValue("middle_name"),
		}

		id, err = h.service.Student.Create(student)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add user: %v", err), http.StatusInternalServerError)
		}
	case "teacher":
		teacher := models.Teacher{
			TelegramID: telegramID,
			FirstName:  r.FormValue("first_name"),
			LastName:   r.FormValue("last_name"),
			MiddleName: r.FormValue("middle_name"),
		}
		id, err = h.service.Teacher.Create(teacher)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add user: %v", err), http.StatusInternalServerError)
		}
	}

	http.Redirect(w, r, fmt.Sprintf("/home/%s/%s", id, role), http.StatusFound)
}
