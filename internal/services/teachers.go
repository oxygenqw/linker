package services

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/google/uuid"
)

type TeacherService interface {
	Create(teacher models.Teacher) (uuid.UUID, error)
	GetByTelegramID(telegramID int64) (models.Teacher, error)
	GetByID(id string) (models.Teacher, error)
	GetAll() ([]models.Teacher, error)
	Update(teacher models.Teacher) error
}

type TeacherServiceImpl struct {
	repository repository.TeacherRepository
}

func NewTeacherService(repository repository.TeacherRepository) *TeacherServiceImpl {
	return &TeacherServiceImpl{
		repository: repository,
	}
}

func (s *TeacherServiceImpl) Create(teacher models.Teacher) (uuid.UUID, error) {
	return s.repository.Create(teacher)
}

func (s *TeacherServiceImpl) GetByTelegramID(telegramID int64) (models.Teacher, error) {
	return s.repository.GetByTelegramID(telegramID)
}

func (s *TeacherServiceImpl) GetByID(id string) (models.Teacher, error) {
	return s.repository.GetByID(id)
}

func (s *TeacherServiceImpl) GetAll() ([]models.Teacher, error) {
	return s.repository.GetAll()
}

func (s *TeacherServiceImpl) Update(teacher models.Teacher) error {
	return s.repository.Update(teacher)
}
