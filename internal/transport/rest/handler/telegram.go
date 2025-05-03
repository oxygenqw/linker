package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

type TelegramHandlerImpl struct {
	logger *logger.Logger
	bot    *bot.Bot
}

func NewTelegramHandler(bot *bot.Bot, logger *logger.Logger) *TelegramHandlerImpl {
	return &TelegramHandlerImpl{
		logger: logger,
		bot:    bot,
	}
}

func (h *TelegramHandlerImpl) CreateBotEndpointHandler(appURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//log.Println("CreateBotEndpointHandler called")
		//log.Printf("Serving %s route", r.URL.Path)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//log.Printf("Request body: %s", string(body))

		var update gotgbot.Update
		if err := json.Unmarshal(body, &update); err != nil {
			log.Printf("Error decoding update: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//log.Printf("Received update: %+v", update)

		if update.Message == nil {
			log.Println("Received update without message")
			http.Error(w, "Bot update didn't include a message", http.StatusBadRequest)
			return
		}

		firstName := update.Message.From.FirstName
		lastName := update.Message.From.LastName
		userName := update.Message.From.Username
		telegramID := update.Message.From.Id
		//log.Printf("Received message: %s", update.Message.Text)

		log.Println("INFO", firstName, lastName, userName, telegramID)

		appURLWithParams := fmt.Sprintf("%s?first_name=%s&last_name=%s&user_name=%s&telegram_id=%s", appURL, firstName, lastName, userName, strconv.FormatInt(telegramID, 10))

		log.Printf("WebApp URL: %s", appURLWithParams)

		message := "Welcome to the Linker"
		opts := &gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Open Linker", WebApp: &gotgbot.WebAppInfo{Url: appURLWithParams}},
				}, {}},
			},
		}

		//log.Printf("Sending message to chat ID: %d", update.Message.Chat.Id)
		if _, err := h.bot.SendMessage(update.Message.Chat.Id, message, opts); err != nil {
			log.Printf("Error sending message: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//log.Printf("Message sent to user: %s (ID: %d)", userName, update.Message.From.Id)
		w.WriteHeader(http.StatusOK)
	}
}
