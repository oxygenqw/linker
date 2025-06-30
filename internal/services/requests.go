package services

import (
	"fmt"
	"time"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/google/uuid"
)

type RequestServiceImpl struct {
	bot               *bot.Bot
	logger            logger.Logger
	userRepository    repository.UserRepository
	studentRepository repository.StudentRepository
	teacherRepository repository.TeacherRepository
	requestRepository repository.RequestRepository
}

func NewRequestService(
	userRepository repository.UserRepository,
	studentRepository repository.StudentRepository,
	teacherRepository repository.TeacherRepository,
	requestRepository repository.RequestRepository,
	bot *bot.Bot,
	logger logger.Logger) *RequestServiceImpl {
	return &RequestServiceImpl{
		bot:               bot,
		logger:            logger,
		userRepository:    userRepository,
		studentRepository: studentRepository,
		teacherRepository: teacherRepository,
		requestRepository: requestRepository,
	}
}

func (s *RequestServiceImpl) SendRequest(senderID string, recipientID string, request models.Message) error {
	s.logger.Info("[S: RequestServiceImpl: SendRequest]")

	senderUUID, err := uuid.Parse(senderID)
	if err != nil {
		return fmt.Errorf("invalid sender ID: %w", err)
	}

	recipientUUID, err := uuid.Parse(recipientID)
	if err != nil {
		return fmt.Errorf("invalid recipient ID: %w", err)
	}

	senderRole, err := s.userRepository.GetRoleByID(senderID)
	if err != nil {
		return fmt.Errorf("failed to get sender role: %w", err)
	}

	var senderUsername string
	var profileURL string

	switch senderRole {
	case "student":
		student, err := s.studentRepository.GetByID(senderID)
		if err != nil {
			return fmt.Errorf("failed to get sender student data: %w", err)
		}
		senderUsername = student.UserName

		profileURL = fmt.Sprintf(
			"https://linker.loca.lt/student/profile/%s/%s/%s",
			senderID,
			"student",
			student.ID,
		)

	case "teacher":
		teacher, err := s.teacherRepository.GetByID(senderID)
		if err != nil {
			return fmt.Errorf("failed to get sender teacher data: %w", err)
		}
		senderUsername = teacher.UserName

		profileURL = fmt.Sprintf(
			"https://linker.loca.lt/teacher/profile/%s/%s/%s",
			senderID,
			"teacher",
			teacher.ID,
		)
	default:
		return fmt.Errorf("unknown sender role: %s", senderRole)
	}

	recipientRole, err := s.userRepository.GetRoleByID(recipientID)
	if err != nil {
		return fmt.Errorf("failed to get recipient role: %w", err)
	}
	if recipientRole == "" {
		return fmt.Errorf("recipient not found")
	}

	var telegramID int64

	switch recipientRole {
	case "student":
		student, err := s.studentRepository.GetByID(recipientID)
		if err != nil {
			return fmt.Errorf("failed to get student recipient: %w", err)
		}
		telegramID = student.TelegramID
	case "teacher":
		teacher, err := s.teacherRepository.GetByID(recipientID)
		if err != nil {
			return fmt.Errorf("failed to get teacher recipient: %w", err)
		}
		telegramID = teacher.TelegramID
	default:
		return fmt.Errorf("unknown recipient role: %s", recipientRole)
	}

	if telegramID == 0 {
		return fmt.Errorf("recipient has no Telegram ID")
	}

	req := models.Request{
		ID:          uuid.New(),
		SenderID:    senderUUID,
		RecipientID: recipientUUID,
		Message:     request.Message,
		Status:      "pending",
		CreatedAt:   time.Now(),
	}

	if err := s.requestRepository.Create(req); err != nil {
		return fmt.Errorf("failed to save request to database: %w", err)
	}

	messageText := fmt.Sprintf(
		"✉️ Новый запрос от @%s: \n%s\n",
		senderUsername,
		request.Message,
	)

	callbackData := fmt.Sprintf("reject_request:%s", req.ID.String())

	opts := &gotgbot.SendMessageOpts{
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
				{
					{
						Text: "Перейти в профиль отправителя",
						Url:  profileURL,
					},
				},
				{
					{
						Text:         "Отклонить",
						CallbackData: callbackData,
					},
				},
			},
		},
	}

	if _, err := s.bot.SendMessage(telegramID, messageText, opts); err != nil {
		s.logger.Error("Failed to send Telegram message", "error", err)
		return fmt.Errorf("failed to send Telegram message: %w", err)
	}

	s.logger.Info("Message sent successfully",
		"sender", senderUsername,
		"recipient", telegramID,
		"requestID", req.ID)

	return nil
}
