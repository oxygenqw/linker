package service

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/google/uuid"
)

type Student interface {
	Create(student models.Student) (uuid.UUID, error)
	GetByTelegramID(telegramID int64) (models.Student, error)
	GetByID(id string) (models.Student, error)
	GetAll() ([]models.Student, error)
	Update(student models.Student) error
}

type StudentService struct {
	repository repository.Student
}

func NewStudentService(repository repository.Student) *StudentService {
	return &StudentService{
		repository: repository,
	}
}

func (s *StudentService) Create(student models.Student) (uuid.UUID, error) {
	return s.repository.Create(student)
}

func (s *StudentService) GetByTelegramID(telegramID int64) (models.Student, error) {
	return s.repository.GetByTelegramID(telegramID)
}

func (s *StudentService) GetByID(id string) (models.Student, error) {
	return s.repository.GetByID(id)
}

func (s *StudentService) GetAll() ([]models.Student, error) {
	return s.repository.GetAll()
}

func (s *StudentService) Update(student models.Student) error {
	return s.repository.Update(student)
}
