package bot

import (
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

type Bot struct {
	token string
}

func NewBot(token string) *gotgbot.Bot {
	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		log.Fatalf("Telegram Bot API initialization error: %v", err)
	}

	return bot
}
