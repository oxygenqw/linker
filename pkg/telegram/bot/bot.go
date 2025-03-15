package bot

import (
	"fmt"
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

type Bot struct {
	*gotgbot.Bot
}

func NewBot(token string) *Bot {
	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		log.Fatalf("Telegram Bot API initialization error: %v", err)
	}
	return &Bot{Bot: bot}
}

func (b *Bot) SetWebhook(url string) error {
	opts := &gotgbot.SetWebhookOpts{
		DropPendingUpdates: true,
	}
	success, err := b.Bot.SetWebhook(url, opts)
	if err != nil {
		log.Printf("Error setting webhook: %v", err)
		return err
	}
	if !success {
		log.Printf("Telegram rejected webhook URL: %s", url)
		return fmt.Errorf("webhook not set")
	}
	log.Printf("Webhook successfully set to: %s", url)
	return nil
}
