package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/services"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type RedirectHadnler interface {
	Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Input(w http.ResponseWriter, r *http.Request)
	UserStudentUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UserTeacherUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UserStudentDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UserTeacherDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type RedirectHandlerImpl struct {
	logger  *logger.Logger
	service *services.Service
}

func NewRedirectHandler(service *services.Service, logger *logger.Logger) *RedirectHandlerImpl {
	return &RedirectHandlerImpl{
		logger:  logger,
		service: service,
	}
}

// ...
// @router GET /
func (h *RedirectHandlerImpl) Input(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("[H: Input] ", "URL: ", r.URL)

	telegramIDStr := r.URL.Query().Get("telegram_id")
	telegramID, err := strconv.ParseInt(telegramIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Telegram ID", http.StatusBadRequest)
		return
	}

	role, err := h.service.UserService.GetRole(telegramID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	switch role {
	case "":
		userName := r.URL.Query().Get("user_name")
		http.Redirect(w, r, fmt.Sprintf("/login/%s/%s", userName, telegramIDStr), http.StatusFound)
	case "student":
		student, err := h.service.StudentService.GetByTelegramID(telegramID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/home/%s/%s", student.ID, role), http.StatusFound)
		return
	case "teacher":
		teacher, err := h.service.TeacherService.GetByTelegramID(telegramID)
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
// TODO: Заменить на CreateUser
func (h *RedirectHandlerImpl) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
			UserName:   r.FormValue("user_name"),
			FirstName:  r.FormValue("first_name"),
			LastName:   r.FormValue("last_name"),
			MiddleName: r.FormValue("middle_name"),
		}

		id, err = h.service.StudentService.Create(student)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add user: %v", err), http.StatusInternalServerError)
		}
	case "teacher":
		teacher := models.Teacher{
			TelegramID: telegramID,
			UserName:   r.FormValue("user_name"),
			FirstName:  r.FormValue("first_name"),
			LastName:   r.FormValue("last_name"),
			MiddleName: r.FormValue("middle_name"),
		}
		id, err = h.service.TeacherService.Create(teacher)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add user: %v", err), http.StatusInternalServerError)
		}
	}

	http.Redirect(w, r, fmt.Sprintf("/home/%s/%s", id, role), http.StatusFound)
}

func (h *RedirectHandlerImpl) UserStudentUpdate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.logger.Info("[H: UpdateStudent] ", "URL: ", r.URL)

	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if student.ID == uuid.Nil {
		http.Error(w, "Missing student ID", http.StatusBadRequest)
		return
	}

	err := h.service.StudentService.Update(student)
	if err != nil {
		h.logger.Error("Failed to update student", "error", err, "ID: ", student.ID)
		http.Error(w, "Failed to update student profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (h *RedirectHandlerImpl) UserTeacherUpdate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.logger.Info("[H: UpdateTeacher] ", "URL: ", r.URL)

	var teacher models.Teacher
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if teacher.ID == uuid.Nil {
		http.Error(w, "Missing student ID", http.StatusBadRequest)
		return
	}

	err := h.service.TeacherService.Update(teacher)
	if err != nil {
		h.logger.Error("Failed to update student", "error", err, "ID: ", teacher.ID)
		http.Error(w, "Failed to update student profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (h *RedirectHandlerImpl) UserStudentDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: DeleteStudent] ", "URL: ", r.URL)

	id := ps.ByName("id")

	err := h.service.StudentService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *RedirectHandlerImpl) UserTeacherDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: DeleteTeacher] ", "URL: ", r.URL)

	id := ps.ByName("id")

	err := h.service.TeacherService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
