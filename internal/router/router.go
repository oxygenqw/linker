package router

import (
	"net/http"

	handler "github.com/Oxygenss/linker/internal/handlers"
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

	router.GET("/", wrapHandler(r.handler.Initialize))
	router.GET("/list", wrapHandler(r.handler.List))
	router.POST("/users", wrapHandler(r.handler.NewUser))
	router.POST("/bot", wrapHandler(r.handler.CreateBotEndpointHandler(r.appURL)))


	router.ServeFiles("/static/*filepath", http.Dir("./ui/static/"))

	return router
}


func wrapHandler(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h(w, r)
	}
}
