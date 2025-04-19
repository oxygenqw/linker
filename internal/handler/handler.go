package handler

import (
	"github.com/Oxygenss/linker/internal/service"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
)

type Handler struct {
	Telegram Telegram
	Pages    Pages
	Redirect Redirect
}

func NewHandler(service *service.Service, logger *logger.Logger, bot *bot.Bot) *Handler {
	return &Handler{
		Telegram: NewTelegramHandler(bot, logger),
		Pages:    NewPagesHandler(service, logger),
		Redirect: NewRedirectHandler(service, logger),
	}
}