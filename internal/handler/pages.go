package handler

import (
	"html/template"
	"net/http"

	"github.com/Oxygenss/linker/internal/service"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

type Pages interface {
	Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Students(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Teachers(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type PagesHandler struct {
	logger  *logger.Logger
	service *service.Service
}

func NewPagesHandler(service *service.Service, logger *logger.Logger) *PagesHandler {
	return &PagesHandler{
		logger:  logger,
		service: service,
	}
}

// Рендерит login.html и передает туда user_name
// @router GET /login/:user_name/:telegram_id
func (h *PagesHandler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: Login] ", "URL: ", r.URL)

	user_name := ps.ByName("user_name")
	telegramID := ps.ByName("telegram_id")

	data := map[string]any{
		"user_name":   user_name,
		"telegram_id": telegramID,
	}

	tmpl, err := template.ParseFiles("./ui/pages/login.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

// Рендерит home.html и передает туда id и role
// @router GET /home/:id/:role
func (h *PagesHandler) Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: Home] ", "URL: ", r.URL)

	id := ps.ByName("id")
	role := ps.ByName("role")

	data := map[string]any{
		"id":   id,
		"role": role,
	}

	tmpl, err := template.ParseFiles("./ui/pages/home.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

// Рендерит profile.html и передает туда информацию о пользователе
// @router GET /profile/:id/:role
func (h *PagesHandler) Profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: Profile]", " URL: ", r.URL)

	id := ps.ByName("id")
	role := ps.ByName("role")

	var data map[string]any
	var err error

	switch role {
	case "student":
		student, err := h.service.Student.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data = map[string]any{
			"user": student,
			"id":   id,
			"role": "student",
		}

	case "teacher":
		teacher, err := h.service.Teacher.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data = map[string]any{
			"user": teacher,
			"id":   id,
			"role": "teacher",
		}
	default:
		http.Error(w, "Invalid role specified", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("./ui/pages/profile.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

// Рендерит students.html и передает туда список студентов
// @router GET /students/:id/:role
func (h *PagesHandler) Students(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: Students] ", "URL: ", r.URL)

	id := ps.ByName("id")
	role := ps.ByName("role")

	students, err := h.service.Student.GetAll()
	if err != nil {
		http.Error(w, "Ошибка при получении списка студентов", http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"students": students,
		"id":       id,
		"role":     role,
	}

	tmpl, err := template.ParseFiles("./ui/pages/students.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

// Рендерит teachers.html и передает туда список преподавателей
// @router GET /teachers/:id/:role
func (h *PagesHandler) Teachers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("[H: Teachers] ", "URL: ", r.URL)

	id := ps.ByName("id")
	role := ps.ByName("role")

	teachers, err := h.service.Teacher.GetAll()
	if err != nil {
		http.Error(w, "Ошибка при получении списка преподавателей", http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"teachers": teachers,
		"id":       id,
		"role":     role,
	}

	tmpl, err := template.ParseFiles("./ui/pages/teachers.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
