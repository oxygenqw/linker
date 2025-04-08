package handlers

import (
	"net/http"

	"github.com/Oxygenss/linker/internal/service"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
)

type Telegram interface {
	CreateBotEndpointHandler(appURL string) http.HandlerFunc
}

type Pages interface {
	Home(w http.ResponseWriter, r *http.Request)
	Initialize(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	Telegram
	Pages
}

func NewHandler(service *service.Service, bot *bot.Bot) *Handler {
	return &Handler{
		Telegram: NewTelegramHandler(bot),
		Pages:    NewPagesHandler(service),
	}
}