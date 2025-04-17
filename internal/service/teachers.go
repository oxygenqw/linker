package service

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/google/uuid"
)

type Teacher interface {
	GetByTelegramID(telegramID int64) (models.Teacher, error)
	GetByID(id string) (models.Teacher, error)
	Create(teacher models.Teacher) (uuid.UUID, error)
	Update(teacher models.Teacher) error
	GetAll() ([]models.Teacher, error)
}

type TeacherService struct {
	repository repository.Teacher
}

func NewTeacherService(repository repository.Teacher) *TeacherService {
	return &TeacherService{
		repository: repository,
	}
}

func (s *TeacherService) GetByID(id string) (models.Teacher, error) {
	return s.repository.GetByID(id)
}

func (s *TeacherService) GetByTelegramID(telegramID int64) (models.Teacher, error) {
	return s.repository.GetByTelegramID(telegramID)
}

func (s *TeacherService) GetAll() ([]models.Teacher, error) {
	return s.repository.GetAll()
}

func (s *TeacherService) Create(teacher models.Teacher) (uuid.UUID, error) {
	return s.repository.Create(teacher)
}

func (s *TeacherService) Update(teacher models.Teacher) error {
	return s.repository.Update(teacher)
}
