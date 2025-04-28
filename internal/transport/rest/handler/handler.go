package handler

import (
	"github.com/Oxygenss/linker/internal/services"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
)

type Handler struct {
	Telegram TelegramHandler
	Pages    PagesHandler
	Redirect RedirectHadnler
}

func NewHandler(service *services.Service, logger *logger.Logger, bot *bot.Bot) *Handler {
	return &Handler{
		Telegram: NewTelegramHandler(bot, logger),
		Pages:    NewPagesHandler(service, logger),
		Redirect: NewRedirectHandler(service, logger),
	}
}
