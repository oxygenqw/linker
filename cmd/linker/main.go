package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Oxygenss/linker/internal/config"
	"github.com/Oxygenss/linker/internal/handler"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/Oxygenss/linker/internal/router"
	"github.com/Oxygenss/linker/internal/service"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
)

func main() {
	config := config.MustLoad()
	logger := logger.GetLogger()

	bot, err := bot.New(config.Telegram.BotToken, config.Telegram.AppURL+"/bot")
	if err != nil {
		logger.Fatalf("error init telegram bot: %v", err)
	}

	storage, err := repository.New(config, &logger)
	if err != nil {
		logger.Fatalf("error init storage: %v", err)
	}

	service := service.New(storage)
	handler := handler.New(service, &logger, bot)
	router := router.New(handler, config.Telegram.AppURL)

	serve := config.Server.Host + ":" + config.Server.Port

	srv := &http.Server{
		Addr:    serve,
		Handler: router.InitRoutes(),
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		logger.Info("Server start:", serve)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
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
