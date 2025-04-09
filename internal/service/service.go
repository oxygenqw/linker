package service

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
)

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

type Service struct {
	Student
	Teacher
}

func NewService(storage *repository.Repository) *Service {
	return &Service{
		Student: NewStudentService(storage.Student),
		Teacher: NewTeacherService(storage.Teacher),
	}
}
