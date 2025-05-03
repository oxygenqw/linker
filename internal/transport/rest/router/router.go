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

	// me
	router.PATCH("/students", r.handler.Redirect.UpdateStudent)
	router.PATCH("/teachers", r.handler.Redirect.UpdateTeacher)

	router.DELETE("/students/:id", r.handler.Redirect.DeleteStudent)
	router.DELETE("/teachers/:id", r.handler.Redirect.DeleteTeacher)

	router.GET("/profile/:id/:role", r.handler.Pages.Profile)

	// redirects
	router.HandlerFunc(http.MethodGet, "/", r.handler.Redirect.Input)
	router.POST("/users/:telegram_id", r.handler.Redirect.Create)

	// html pages
	router.GET("/login/:user_name/:telegram_id", r.handler.Pages.Login)
	router.GET("/home/:id/:role", r.handler.Pages.Home)

	router.GET("/student/edit/:id", r.handler.Pages.EditStudentProfile)
	router.GET("/teacher/edit/:id", r.handler.Pages.EditTeacherProfile)

	router.GET("/students/:id/:role", r.handler.Pages.Students)
	router.GET("/teachers/:id/:role", r.handler.Pages.Teachers)

	router.GET("/student/profile/:id/:role/:student_id", r.handler.Pages.StudentProfile)
	router.GET("/teacher/profile/:id/:role/:teacher_id", r.handler.Pages.TeacherProfile)

	// static
	router.ServeFiles("/static/*filepath", http.Dir("./web/static/"))

	return router
}
