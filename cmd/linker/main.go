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
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
)

func main() {
	config := config.MustLoad()

	logger := logger.GetLogger()


	bot := bot.NewBot(config.Telegram.BotToken)
	bot.SetWebhook(config.Telegram.AppURL + "/bot")

	storage, err := repository.NewRepository(config)
	if err != nil {
		logger.Fatalf("error init storage: %v", err)
	}

	service := service.NewService(storage)

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
		logger.Info("Serve start:", serve)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	logger.Info("Server stopping...")

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		logger.Fatalf("server shutdown failed: %v", err)
	}

	logger.Println("Server stopped")
}
