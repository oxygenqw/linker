package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Oxygenss/linker/internal/config"
	"github.com/Oxygenss/linker/internal/handlers"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/Oxygenss/linker/internal/router"
	"github.com/Oxygenss/linker/internal/service"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
)

func main() {
	config := config.MustLoad()

	bot := bot.NewBot(config.Telegram.BotToken)
	bot.SetWebhook(config.Telegram.AppURL + "/bot")

	repository, err := repository.New(config)
	if err != nil {
		log.Fatalf("init error: %v", err)
	}

	service := service.NewService(repository)

	handler := handlers.NewHandler(service, bot)

	router := router.NewRouter(handler, config.Telegram.AppURL)

	serve := config.Server.Host + ":" + config.Server.Port

	srv := &http.Server{
		Addr:    serve,
		Handler: router.InitRoutes(),
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		log.Println("Serve start:", serve)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	log.Println("Server stopping...")

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("Server stopped")
}
