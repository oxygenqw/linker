package service

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
)

type Service struct {
	Student
	Teacher
	User
}

type Student interface {
	GetByTelegramID(telegramID int64) (models.Student, error)
	GetAll() ([]models.Student, error)
	Create(student models.Student) error
}

type Teacher interface {
	GetByTelegramID(telegramID int64) (models.Teacher, error)
	GetAll() ([]models.Teacher, error)
	Create(teacher models.Teacher) error
}

type User interface {
	CheckByTelegramID(telegramID int64) (bool, error)
}

func New(repository *repository.Repository) *Service {
	return &Service{
		Student: NewStudentService(repository.Student),
		Teacher: NewTeacherService(repository.Teacher),
		User:    NewUserService(repository.User),
	}
}
