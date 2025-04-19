package bot

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

type Bot struct {
	*gotgbot.Bot
}

// New creates a new bot instance and sets up the webhook
// token - bot token provided by BotFather
// webhookURL - full URL for receiving updates (e.g., "https://example.com/bot")
func NewTelegramBot(token string, webhookURL string) (*Bot, error) {
	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		return nil, err
	}

	b := &Bot{Bot: bot}

	if err := b.setWebhook(webhookURL); err != nil {
		return nil, fmt.Errorf("set webhook failed: %w", err)
	}

	return b, nil

}

// setWebhook private method for webhook setup
func (b *Bot) setWebhook(url string) error {
	// Очищаем историю сообщений, пока бот не работал
	opts := &gotgbot.SetWebhookOpts{
		DropPendingUpdates: true,
	}

	success, err := b.Bot.SetWebhook(url, opts)
	if err != nil {
		return fmt.Errorf("telegram API error: %w", err)
	}

	if !success {
		return fmt.Errorf("telegram rejected webhook URL: %s", url)
	}

	return nil
}
