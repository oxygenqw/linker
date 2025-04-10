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

	router.HandlerFunc(http.MethodPost, "/bot", r.handler.CreateBotEndpointHandler(r.appURL))

	router.HandlerFunc(http.MethodGet, "/", r.handler.Input)

	router.HandlerFunc(http.MethodPost, "/users", r.handler.NewUser)

	router.HandlerFunc(http.MethodGet, "/students", r.handler.Students)
	router.HandlerFunc(http.MethodGet, "/teachers", r.handler.Teachers)

	router.HandlerFunc(http.MethodGet, "/home", r.handler.Home)

	router.HandlerFunc(http.MethodGet, "/profile", r.handler.Profile)

	router.ServeFiles("/static/*filepath", http.Dir("./ui/static/"))

	return router
}
