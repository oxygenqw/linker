package router

import (
	"net/http"

	"github.com/Oxygenss/linker/internal/handler"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	handler *handler.Handler
	appURL  string
}

func New(handler *handler.Handler, appURL string) *Router {
	return &Router{handler: handler, appURL: appURL}
}

func (r *Router) InitRoutes() *httprouter.Router {
	router := httprouter.New()

	// telegram api
	router.HandlerFunc(http.MethodPost, "/bot", r.handler.CreateBotEndpointHandler(r.appURL))

	// redirect
	router.HandlerFunc(http.MethodGet, "/", r.handler.Input)
	router.POST("/users/:telegram_id", r.handler.NewUser)

	// html pages
	router.GET("/login/:user_name/:telegram_id", r.handler.Login)
	router.GET("/home/:id/:role", r.handler.Home)
	router.GET("/students/:id/:role", r.handler.Students)
	router.GET("/teachers/:id/:role", r.handler.Teachers)
	router.GET("/profile/:id/:role", r.handler.Profile)

	// static
	router.ServeFiles("/static/*filepath", http.Dir("./ui/static/"))

	return router
}
