package services

import (
	"fmt"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

type RequestServiceImpl struct {
	bot        *bot.Bot
	logger     logger.Logger
	repository repository.StudentRepository
}

func NewRequestService(repository repository.StudentRepository, bot *bot.Bot, logger logger.Logger) *RequestServiceImpl {
	return &RequestServiceImpl{
		bot:        bot,
		logger:     logger,
		repository: repository,
	}
}

func (s *RequestServiceImpl) CreateRequest(senderID string, recipientID string, request models.Request) error {
	s.logger.Info("[S: CreateRequest]")
	fmt.Println(recipientID, senderID, request.Message)

	recipient, err := s.repository.GetByID(recipientID)
	if err != nil {
		return fmt.Errorf("failed to get recipient: %w", err)
	}

	if recipient.TelegramID == 0 {
		return fmt.Errorf("recipient has no Telegram ID")
	}

	fmt.Println("ID: ", recipient.TelegramID)

	messageText := fmt.Sprintf("✉️ Новый запрос:\n\n%s", request.Message)

	opts := &gotgbot.SendMessageOpts{
		ParseMode: "Markdown",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
				{
					{
						Text:         "Ответить",
						CallbackData: fmt.Sprintf("reply_to:%s", senderID),
					},
				},
			},
		},
	}

	_, err = s.bot.SendMessage(recipient.TelegramID, messageText, opts)
	if err != nil {
		s.logger.Error("Failed to send Telegram message", err)
		return fmt.Errorf("failed to send Telegram message: %w", err)
	}

	s.logger.Info(fmt.Sprintf("Message sent to %d: %s", recipient.TelegramID, messageText))
	return nil
}
