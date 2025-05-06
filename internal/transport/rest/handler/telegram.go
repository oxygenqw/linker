package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
		body, err := io.ReadAll(r.Body)
		if err != nil {
			h.logger.Error("Error reading request body", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var update gotgbot.Update
		if err := json.Unmarshal(body, &update); err != nil {
			h.logger.Error("Error decoding update", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if update.Message == nil {
			h.logger.Warn("Received update without message")
			http.Error(w, "Bot update didn't include a message", http.StatusBadRequest)
			return
		}

		user := update.Message.From
		telegramID := user.Id

		params := url.Values{}
		params.Add("first_name", user.FirstName)
		params.Add("last_name", user.LastName)
		params.Add("user_name", user.Username)
		params.Add("telegram_id", strconv.FormatInt(telegramID, 10))

		appURLWithParams := fmt.Sprintf("%s?%s", appURL, params.Encode())

		message := "Welcome to the Linker"
		opts := &gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Open Linker", WebApp: &gotgbot.WebAppInfo{Url: appURLWithParams}},
				}},
			},
		}

		if _, err := h.bot.SendMessage(telegramID, message, opts); err != nil {
			h.logger.Error("Error sending message", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
