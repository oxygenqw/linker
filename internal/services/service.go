package services

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
	"github.com/google/uuid"
)

type UserService interface {
	GetRoleByTelegramID(telegramID int64) (string, error)
	GetRoleByID(id string) (string, error)
}

type StudentService interface {
	Create(student models.Student) (uuid.UUID, error)
	GetByTelegramID(telegramID int64) (models.Student, error)
	GetByID(id string) (models.Student, error)
	GetAll() ([]models.Student, error)
	Search(search string) ([]models.Student, error)
	Update(student models.Student) error
	Delete(id string) error
}

type TeacherService interface {
	Create(teacher models.Teacher) (uuid.UUID, error)
	GetByTelegramID(telegramID int64) (models.Teacher, error)
	GetByID(id string) (models.Teacher, error)
	GetAll() ([]models.Teacher, error)
	Search(search string) ([]models.Teacher, error)
	Update(teacher models.Teacher) error
	Delete(id string) error
}

type RequestService interface {
	SendRequest(senderID string, recipientID string, message models.Message) error
}

type WorkService interface {
	Create(models.Work) error
	GetAll(userID uuid.UUID) ([]models.Work, error)
	Delete(id uuid.UUID) error
}

type Service struct {
	StudentService StudentService
	TeacherService TeacherService
	UserService    UserService
	RequestService RequestService
	WorkService    WorkService
}

func NewService(repository *repository.Repository, logger *logger.Logger, bot *bot.Bot) *Service {
	return &Service{
		StudentService: NewStudentService(repository.StudentRepository),
		TeacherService: NewTeacherService(repository.TeacherRepository),
		UserService:    NewUserService(repository.UserRepository),
		RequestService: NewRequestService(
			repository.UserRepository,
			repository.StudentRepository,
			repository.TeacherRepository,
			repository.RequestRepository,
			bot,
			*logger),
		WorkService: NewWorkService(repository.WorkRepository),
	}
}
