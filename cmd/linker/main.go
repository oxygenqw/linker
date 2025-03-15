package main

import (
	"log"
	"net/http"

	"github.com/Oxygenss/linker/internal/config"
	"github.com/Oxygenss/linker/internal/handler"
	"github.com/Oxygenss/linker/internal/router"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
)

func main() {

	config := config.MustLoad()

	bot := bot.NewBot(config.Telegram.BotToken)
	bot.SetWebhook(config.Telegram.AppURL + "/bot")

	handler := handler.NewHandler(bot)

	router := router.NewRouter(handler, config.Telegram.AppURL)

	serve := config.Server.Host + ":" + config.Server.Port

	log.Println("Serve start:", serve)

	log.Fatal(http.ListenAndServe(serve, router.InitRoutes()))
}
