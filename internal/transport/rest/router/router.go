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

	// users
	router.GET("/", r.handler.Users.TelegramAuth)

	router.POST("/users/student", r.handler.Users.CreateStudent)
	router.POST("/users/teacher", r.handler.Users.CreateTeacher)

	router.GET("/users/student/:id", r.handler.Users.StudentProfile)
	router.GET("/users/teacher/:id", r.handler.Users.TeacherProfile)

	router.GET("/student/edit/:id", r.handler.Users.EditStudentProfile)
	router.GET("/teacher/edit/:id", r.handler.Users.EditTeacherProfile)

	router.PATCH("/users/student", r.handler.Users.StudentUpdate)
	router.PATCH("/users/teacher", r.handler.Users.TeacherUpdate)

	router.DELETE("/users/student/:id", r.handler.Users.StudentDelete)
	router.DELETE("/users/teacher/:id", r.handler.Users.TeacherDelete)

	// html pages
	router.GET("/login/:user_name/:telegram_id", r.handler.Pages.Login)
	router.GET("/home/:id/:role", r.handler.Pages.Home)

	router.GET("/students/:id/:role", r.handler.Pages.Students)
	router.GET("/teachers/:id/:role", r.handler.Pages.Teachers)

	router.GET("/student/profile/:id/:role/:student_id", r.handler.Pages.StudentProfile)
	router.GET("/teacher/profile/:id/:role/:teacher_id", r.handler.Pages.TeacherProfile)

	// static
	router.ServeFiles("/static/*filepath", http.Dir("./web/static/"))

	return router
}
