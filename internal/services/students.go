package services

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/google/uuid"
)

type StudentService interface {
	Create(student models.Student) (uuid.UUID, error)
	GetByTelegramID(telegramID int64) (models.Student, error)
	GetByID(id string) (models.Student, error)
	GetAll() ([]models.Student, error)
	Update(student models.Student) error
}

type StudentServiceImpl struct {
	repository repository.StudentRepositiry
}

func NewStudentService(repository repository.StudentRepositiry) *StudentServiceImpl {
	return &StudentServiceImpl{
		repository: repository,
	}
}

func (s *StudentServiceImpl) Create(student models.Student) (uuid.UUID, error) {
	return s.repository.Create(student)
}

func (s *StudentServiceImpl) GetByTelegramID(telegramID int64) (models.Student, error) {
	return s.repository.GetByTelegramID(telegramID)
}

func (s *StudentServiceImpl) GetByID(id string) (models.Student, error) {
	return s.repository.GetByID(id)
}

func (s *StudentServiceImpl) GetAll() ([]models.Student, error) {
	return s.repository.GetAll()
}

func (s *StudentServiceImpl) Update(student models.Student) error {
	return s.repository.Update(student)
}
