package service

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
)

type StudentService struct {
	repository repository.Student
}

func NewStudentService(repository repository.Student) *StudentService {
	return &StudentService{
		repository: repository,
	}
}

func (s *StudentService) GetByTelegramID(telegramID int64) (models.Student, error) {
	return s.repository.GetByTelegramID(telegramID)
}

func (s *StudentService) GetAll() ([]models.Student, error) {
	return s.repository.GetAll()
}

func (s *StudentService) Create(student models.Student) error {
	return s.repository.Create(student)
}
