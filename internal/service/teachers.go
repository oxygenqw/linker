package service

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
)

type TeacherService struct {
	repository repository.Teacher
}

func NewTeacherService(repository repository.Teacher) *TeacherService {
	return &TeacherService{
		repository: repository,
	}
}

func (s *TeacherService) GetByTelegramID(telegramID int64) (models.Teacher, error) {
	return s.repository.GetByTelegramID(telegramID)
}

func (s *TeacherService) GetAll() ([]models.Teacher, error) {
	return s.repository.GetAll()
}

func (s *TeacherService) Create(teacher models.Teacher) error {
	return s.repository.Create(teacher)
}
