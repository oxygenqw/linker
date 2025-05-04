package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/services"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

type CardsHandlerImpl struct {
	logger    *logger.Logger
	service   *services.Service
	templates map[string]*template.Template
}

func NewCardsHandler(service *services.Service, logger *logger.Logger) *CardsHandlerImpl {
	files := []string{
		"./web/pages/students.html",
		"./web/pages/teachers.html",
		"./web/pages/student_user_profile.html",
		"./web/pages/teacher_user_profile.html",
	}

	templates := make(map[string]*template.Template)
	for _, file := range files {
		name := filepath.Base(file)
		templates[name] = template.Must(template.ParseFiles(file))
	}

	return &CardsHandlerImpl{
		logger:    logger,
		service:   service,
		templates: templates,
	}
}

func (h *CardsHandlerImpl) StudentProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: StudentProfile] ", "URL: ", r.URL)

	id := params.ByName("id")
	role := params.ByName("role")
	student_id := params.ByName("student_id")

	student, err := h.service.StudentService.GetByID(student_id)
	if err != nil {

	}

	data := map[string]any{
		"student": student,
		"id":      id,
		"role":    role,
	}

	h.renderTemplate(w, "student_user_profile.html", data)
}

func (h *CardsHandlerImpl) TeacherProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: TeacherProfile] ", "URL: ", r.URL)

	id := params.ByName("id")
	role := params.ByName("role")
	teacher_id := params.ByName("teacher_id")

	teacher, err := h.service.TeacherService.GetByID(teacher_id)
	if err != nil {

	}

	data := map[string]any{
		"teacher": teacher,
		"id":      id,
		"role":    role,
	}

	h.renderTemplate(w, "teacher_user_profile.html", data)
}

// Рендерит students.html и передает туда список студентов
// @router GET /students/:id/:role
func (h *CardsHandlerImpl) Students(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: Students] ", "URL: ", r.URL)

	id := params.ByName("id")
	role := params.ByName("role")
	search := r.URL.Query().Get("search")

	var students []models.Student
	var err error

	if search != "" {
		students, err = h.service.StudentService.Search(search)
	} else {
		students, err = h.service.StudentService.GetAll()
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при получении списка студентов, %s", err), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"students": students,
		"id":       id,
		"role":     role,
	}

	h.renderTemplate(w, "students.html", data)
}

func (h *CardsHandlerImpl) Teachers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: Teachers] ", "URL: ", r.URL)

	id := params.ByName("id")
	role := params.ByName("role")
	search := r.URL.Query().Get("search")

	var teachers []models.Teacher
	var err error

	if search != "" {
		teachers, err = h.service.TeacherService.Search(search)
	} else {
		teachers, err = h.service.TeacherService.GetAll()
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при получении списка преподавателей, %s", err), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"teachers": teachers,
		"id":       id,
		"role":     role,
	}

	h.renderTemplate(w, "teachers.html", data)
}

func (h *CardsHandlerImpl) renderTemplate(w http.ResponseWriter, tmplName string, data any) {
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
