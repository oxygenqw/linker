package router

import (
	"net/http"

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

	router.Get("/", r.handler.Home)
	router.Get("/list", r.handler.List)
	router.Post("/initialize", r.handler.Initialize)



	router.Post("/bot", r.handler.CreateBotEndpointHandler(r.appURL))
	

	fs := http.FileServer(http.Dir("./templates/home"))
	router.Handle("/*", http.StripPrefix("/", fs))

	return router
}
