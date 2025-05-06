package handler

import (
	"net/http"

	"github.com/Oxygenss/linker/internal/renderer"
	"github.com/Oxygenss/linker/internal/services"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
	"github.com/julienschmidt/httprouter"
)

type TelegramHandler interface {
	CreateBotEndpointHandler(appURL string) http.HandlerFunc
}

type UserHandler interface {
	TelegramAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	CreateStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	CreateTeacher(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	StudentProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	TeacherProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	EditStudentProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	EditTeacherProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	StudentUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	TeacherUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	StudentDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	TeacherDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type PagesHandler interface {
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Home(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type CardsHandler interface {
	Students(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Teachers(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	StudentProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	TeacherProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type RequestsHandler interface {
	ToRequestForm(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	RequestToStudent(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type Handler struct {
	Telegram TelegramHandler
	Pages    PagesHandler
	Users    UserHandler
	Cards    CardsHandler
	Requests RequestsHandler
}

func NewHandler(service *services.Service, renderer *renderer.TemplateRenderer, logger *logger.Logger, bot *bot.Bot) *Handler {
	return &Handler{
		Telegram: NewTelegramHandler(bot, logger),
		Pages:    NewPagesHandler(service, renderer, logger),
		Users:    NewUserHandler(service, renderer, logger),
		Cards:    NewCardsHandler(service, renderer, logger),
		Requests: NewRequestsHandler(service, renderer, logger),
	}
}
