package main

import (
	"net/http"

	"github.com/Oxygenss/linker/internal/config"
	"github.com/Oxygenss/linker/internal/handler"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/Oxygenss/linker/internal/router"
	"github.com/Oxygenss/linker/internal/service"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
)

func main() {
	cfg := config.MustLoad()
	log := logger.GetLogger()

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
	service := service.NewService(repository)

	log.Info("Initialize handler...")
	handler := handler.NewHandler(service, &log, bot)

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
