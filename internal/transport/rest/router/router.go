package router

import (
	"net/http"

	"github.com/Oxygenss/linker/internal/transport/rest/handler"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	handler *handler.Handler
	appURL  string
}

func NewRouter(handler *handler.Handler, appURL string) *Router {
	return &Router{handler: handler, appURL: appURL}
}

func (r *Router) InitRoutes() *httprouter.Router {
	router := httprouter.New()

	// telegram api
	router.HandlerFunc(http.MethodPost, "/bot", r.handler.Telegram.CreateBotEndpointHandler(r.appURL))

	// redirects
	router.HandlerFunc(http.MethodGet, "/", r.handler.Redirect.Input)
	router.POST("/users/new/:telegram_id", r.handler.Redirect.CreateUser)

	router.POST("/student/update/:id", r.handler.Redirect.UpdateStudent)
	router.POST("/teacher/update/:id", r.handler.Redirect.UpdateTeacher)

	// html pages
	router.GET("/login/:user_name/:telegram_id", r.handler.Pages.Login)
	router.GET("/home/:id/:role", r.handler.Pages.Home)

	router.GET("/profile/:id/:role", r.handler.Pages.Profile)

	router.GET("/student/edit/:id", r.handler.Pages.EditStudentProfile)
	router.GET("/teacher/edit/:id", r.handler.Pages.EditTeacherProfile)

	router.GET("/students/:id/:role", r.handler.Pages.Students)
	router.GET("/teachers/:id/:role", r.handler.Pages.Teachers)

	router.GET("/student/profile/:id/:role/:student_id", r.handler.Pages.StudentProfile)
	router.GET("/teacher/profile/:id/:role/:teacher_id", r.handler.Pages.TeacherProfile)

	// static
	router.ServeFiles("/static/*filepath", http.Dir("./ui/static/"))

	return router
}
