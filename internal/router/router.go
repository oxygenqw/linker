package router

import (
	"net/http"

	handler "github.com/Oxygenss/linker/internal/handlers"
)

type Router struct {
	handler *handler.Handler
	appURL  string
}

func NewRouter(handler *handler.Handler, appURL string) *Router {
	return &Router{handler: handler, appURL: appURL}
}

func (r *Router) InitRoutes() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("/", r.handler.Home)
	router.HandleFunc("/list", r.handler.List)
	router.HandleFunc("/initialize", r.handler.Initialize)
	router.HandleFunc("/bot", r.handler.CreateBotEndpointHandler(r.appURL))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	return router
}
