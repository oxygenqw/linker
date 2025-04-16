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
	router.HandlerFunc(http.MethodPost, "/bot", r.handler.Telegram.CreateBotEndpointHandler(r.appURL))

	// redirect
	router.HandlerFunc(http.MethodGet, "/", r.handler.Redirect.Input)
	router.POST("/users/:telegram_id", r.handler.Redirect.NewUser)

	// html pages
	router.GET("/login/:user_name/:telegram_id", r.handler.Pages.Login)
	router.GET("/home/:id/:role", r.handler.Pages.Home)
	router.GET("/students/:id/:role", r.handler.Pages.Students)
	router.GET("/teachers/:id/:role", r.handler.Pages.Teachers)
	router.GET("/profile/:id/:role", r.handler.Pages.Profile)

	// static
	router.ServeFiles("/static/*filepath", http.Dir("./ui/static/"))

	return router
}
