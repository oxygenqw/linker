package main

import (
	"net/http"

	"github.com/Oxygenss/linker/internal/config"
	"github.com/Oxygenss/linker/internal/renderer"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/Oxygenss/linker/internal/services"
	"github.com/Oxygenss/linker/internal/transport/rest/handler"
	"github.com/Oxygenss/linker/internal/transport/rest/router"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
)

func main() {
	cfg := config.MustLoad()
	log := logger.NewLogger()

	log.Info("Initialize telegram bot...")
	bot, err := bot.NewTelegramBot(cfg.Telegram.BotToken, cfg.Telegram.AppURL+"/bot")
	if err != nil {
		log.Fatalf("Error init telegram bot: %v", err)
	}

	log.Info("Initialize repository...")
	repository, err := repository.NewRepository(cfg, &log)
	if err != nil {
		log.Fatalf("Error init storage: %v", err)
	}

	log.Info("Initialize service...")
	service := services.NewService(repository, &log, bot)

	log.Info("Initialize renderer...")
	renderer, err := renderer.NewTemplateRenderer()
	if err != nil {
		log.Fatalf("Error init renderer: %v", err)
	}

	log.Info("Initialize handler...")
	handler := handler.NewHandler(service, renderer, &log, bot)

	log.Info("Initialize router...")
	router := router.NewRouter(handler, cfg.Telegram.AppURL)

	serve := cfg.Server.Host + ":" + cfg.Server.Port

	srv := &http.Server{
		Addr:    serve,
		Handler: router.InitRoutes(),
	}

	log.Info("HTTP Server starting... ", serve)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting HTTP server: %s\n", err)
	}
}
