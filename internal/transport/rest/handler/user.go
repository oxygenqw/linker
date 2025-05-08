package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/renderer"
	"github.com/Oxygenss/linker/internal/services"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type UserHandlerImpl struct {
	logger   *logger.Logger
	service  *services.Service
	renderer *renderer.TemplateRenderer
}

func NewUserHandler(service *services.Service, renderer *renderer.TemplateRenderer, logger *logger.Logger) *UserHandlerImpl {
	return &UserHandlerImpl{
		logger:   logger,
		service:  service,
		renderer: renderer,
	}
}

func (h *UserHandlerImpl) TelegramAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: TelegramAuth] ", "URL: ", r.URL)

	telegramIDStr := r.URL.Query().Get("telegram_id")
	telegramID, err := strconv.ParseInt(telegramIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Telegram ID", http.StatusBadRequest)
		return
	}

	role, err := h.service.UserService.GetRoleByTelegramID(telegramID)
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

func (h *UserHandlerImpl) CreateStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: CreateStudent] ", "URL: ", r.URL)

	var student models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
		return
	}

	if student.TelegramID == 0 {
		http.Error(w, "Missing telegram_id", http.StatusBadRequest)
		return
	}

	id, err := h.service.StudentService.Create(student)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add student: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"id":"%s"}`, id)
}

func (h *UserHandlerImpl) CreateTeacher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: CreateTeacher] ", "URL: ", r.URL)

	var teacher models.Teacher
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&teacher); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if teacher.TelegramID == 0 {
		http.Error(w, "Missing telegram_id", http.StatusBadRequest)
		return
	}

	id, err := h.service.TeacherService.Create(teacher)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add teacher: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"id":"%s"}`, id)
}

func (h *UserHandlerImpl) StudentProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: StudentProfile]", " URL: ", r.URL)

	id := ps.ByName("id")

	student, err := h.service.StudentService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	works, err := h.service.WorkService.GetAll(student.ID)
	if err != nil {
		http.Error(w, "Ошибка получения работ: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"student": student,
		"works":   works,
	}

	h.renderer.RenderTemplate(w, "student_profile.html", data)
}

func (h *UserHandlerImpl) TeacherProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: TeacherProfile]", " URL: ", r.URL)

	id := ps.ByName("id")

	teacher, err := h.service.TeacherService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	works, err := h.service.WorkService.GetAll(teacher.ID)
	if err != nil {
		http.Error(w, "Ошибка получения работ: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"teacher": teacher,
		"works":   works,
	}

	h.renderer.RenderTemplate(w, "teacher_profile.html", data)
}

func (h *UserHandlerImpl) EditStudentProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: EditStudentProfile]", " URL: ", r.URL)

	id := ps.ByName("id")

	var data map[string]any

	student, err := h.service.StudentService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data = map[string]any{
		"student": student,
	}

	h.renderer.RenderTemplate(w, "student_editor.html", data)
}

func (h *UserHandlerImpl) EditTeacherProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: EditTeacherProfile]", " URL: ", r.URL)

	id := ps.ByName("id")

	var data map[string]any

	teacher, err := h.service.TeacherService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data = map[string]any{
		"teacher": teacher,
	}

	h.renderer.RenderTemplate(w, "teacher_editor.html", data)
}

func (h *UserHandlerImpl) StudentUpdate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.logger.Info("[H: StudentUpdate] ", "URL: ", r.URL)

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
}

func (h *UserHandlerImpl) TeacherUpdate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.logger.Info("[H: TeacherUpdate] ", "URL: ", r.URL)

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
}

func (h *UserHandlerImpl) StudentDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: StudentDelete] ", "URL: ", r.URL)

	id := ps.ByName("id")

	err := h.service.StudentService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandlerImpl) TeacherDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: TeacherDelete] ", "URL: ", r.URL)

	id := ps.ByName("id")

	err := h.service.TeacherService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
