package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Oxygenss/linker/internal/handler"
	"github.com/Oxygenss/linker/internal/router"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting API service")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Printf("TELEGRAM_BOT_TOKEN: %s", os.Getenv("TELEGRAM_BOT_TOKEN"))
	log.Printf("TELEGRAM_WEB_APP_URL: %s", os.Getenv("TELEGRAM_WEB_APP_URL"))

	bot := bot.NewBot(os.Getenv("TELEGRAM_BOT_TOKEN"))

	handler := handler.NewHandler(bot)

	router := router.NewRouter(handler, os.Getenv("TELEGRAM_WEB_APP_URL"))

	log.Fatal(http.ListenAndServe(":8080", router.InitRoutes()))
}
