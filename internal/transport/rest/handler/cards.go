package handler

import (
	"fmt"
	"net/http"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/renderer"
	"github.com/Oxygenss/linker/internal/services"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

type CardsHandlerImpl struct {
	logger   *logger.Logger
	service  *services.Service
	renderer *renderer.TemplateRenderer
}

func NewCardsHandler(service *services.Service, renderer *renderer.TemplateRenderer, logger *logger.Logger) *CardsHandlerImpl {
	return &CardsHandlerImpl{
		logger:   logger,
		service:  service,
		renderer: renderer,
	}
}

func (h *CardsHandlerImpl) StudentProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: StudentProfile] ", "URL: ", r.URL)

	id := params.ByName("id")
	role := params.ByName("role")
	student_id := params.ByName("student_id")

	student, err := h.service.StudentService.GetByID(student_id)
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
		"id":      id,
		"role":    role,
	}

	h.renderer.RenderTemplate(w, "student_user_profile.html", data)
}

func (h *CardsHandlerImpl) TeacherProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("[H: TeacherProfile] ", "URL: ", r.URL)

	id := params.ByName("id")
	role := params.ByName("role")
	teacher_id := params.ByName("teacher_id")

	teacher, err := h.service.TeacherService.GetByID(teacher_id)
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
		"id":      id,
		"role":    role,
	}

	h.renderer.RenderTemplate(w, "teacher_user_profile.html", data)
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

	h.renderer.RenderTemplate(w, "students.html", data)
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

	h.renderer.RenderTemplate(w, "teachers.html", data)
}
