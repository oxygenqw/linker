package router

import (
	"github.com/Oxygenss/linker/internal/handler"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	handler *handler.Handler
	appURL  string
}

func NewRouter(handler *handler.Handler, appURL string) *Router {
	return &Router{handler: handler, appURL: appURL}
}

func (r *Router) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", r.handler.Welcome)
	router.Post("/bot", r.handler.CreateBotEndpointHandler(r.appURL))

	return router
}
