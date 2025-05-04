package services

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/google/uuid"
)

type StudentServiceImpl struct {
	repository repository.StudentRepository
}

func NewStudentService(repository repository.StudentRepository) StudentService {
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

func (s *StudentServiceImpl) Search(search string) ([]models.Student, error) {
	return s.repository.Search(search)
}

func (s *StudentServiceImpl) Update(student models.Student) error {
	return s.repository.Update(student)
}

func (s *StudentServiceImpl) Delete(id string) error {
	return s.repository.Delete(id)
}
