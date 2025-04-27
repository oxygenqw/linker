package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/Oxygenss/linker/internal/service"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

type PagesHandler interface {
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Home(w http.ResponseWriter, r *http.Request, params httprouter.Params)

	Profile(w http.ResponseWriter, r *http.Request, params httprouter.Params)

	Students(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Teachers(w http.ResponseWriter, r *http.Request, params httprouter.Params)

	EditStudentProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	EditTeacherProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params)

	StudentProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	TeacherProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type PagesHandlerImpl struct {
	logger    *logger.Logger
	service   *service.Service
	templates map[string]*template.Template
}

func NewPagesHandler(service *service.Service, logger *logger.Logger) *PagesHandlerImpl {
	files, err := filepath.Glob("./ui/pages/*.html")
	if err != nil {
		logger.Fatal("Failed to find template files", "error", err)
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

// Рендерит profile.html и передает туда информацию о пользователе
// @router GET /profile/:id/:role
func (h *PagesHandlerImpl) Profile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: Profile]", " URL: ", r.URL)

	id := params.ByName("id")
	role := params.ByName("role")

	var data map[string]any

	switch role {
	case "student":
		student, err := h.service.StudentService.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data = map[string]any{
			"user": student,
		}

		h.renderTemplate(w, "student_profile.html", data)

	case "teacher":
		teacher, err := h.service.TeacherService.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data = map[string]any{
			"user": teacher,
		}

		h.renderTemplate(w, "teacher_profile.html", data)
	default:
		http.Error(w, "Invalid role specified", http.StatusBadRequest)
		return
	}
}

func (h *PagesHandlerImpl) EditStudentProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: EditStudentProfile]", " URL: ", r.URL)

	id := params.ByName("id")

	var data map[string]any

	student, err := h.service.StudentService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data = map[string]any{
		"user": student,
	}

	h.renderTemplate(w, "student_editor.html", data)
}

func (h *PagesHandlerImpl) EditTeacherProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: EditTeacherProfile]", " URL: ", r.URL)

	id := params.ByName("id")

	var data map[string]any

	teacher, err := h.service.TeacherService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data = map[string]any{
		"user": teacher,
	}

	h.renderTemplate(w, "teacher_editor.html", data)
}

func (h *PagesHandlerImpl) StudentProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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

func (h *PagesHandlerImpl) TeacherProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
func (h *PagesHandlerImpl) Students(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: Students] ", "URL: ", r.URL)

	id := params.ByName("id")
	role := params.ByName("role")

	students, err := h.service.StudentService.GetAll()
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

// Рендерит teachers.html и передает туда список преподавателей
// @router GET /teachers/:id/:role
func (h *PagesHandlerImpl) Teachers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: Teachers] ", "URL: ", r.URL)

	id := params.ByName("id")
	role := params.ByName("role")

	teachers, err := h.service.TeacherService.GetAll()
	if err != nil {
		http.Error(w, "Ошибка при получении списка преподавателей", http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"teachers": teachers,
		"id":       id,
		"role":     role,
	}

	h.renderTemplate(w, "teachers.html", data)
}
