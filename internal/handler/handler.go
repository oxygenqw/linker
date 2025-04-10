package handler

import (
	"net/http"

	"github.com/Oxygenss/linker/internal/service"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
)

type Handler struct {
	Telegram
	Pages
}

type Telegram interface {
	CreateBotEndpointHandler(appURL string) http.HandlerFunc
}

type Pages interface {
	NewUser(w http.ResponseWriter, r *http.Request)
	Input(w http.ResponseWriter, r *http.Request)
	Home(w http.ResponseWriter, r *http.Request)
	Students(w http.ResponseWriter, r *http.Request)
	Teachers(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
}

func New(service *service.Service, bot *bot.Bot) *Handler {
	return &Handler{
		Telegram: NewTelegramHandler(bot),
		Pages:    NewPagesHandler(service),
	}
}
