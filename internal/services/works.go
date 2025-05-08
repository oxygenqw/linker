package services

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/google/uuid"
)

type WorkServiceImpl struct {
	repository repository.WorkRepository
}

func NewWorkService(repository repository.WorkRepository) *WorkServiceImpl {
	return &WorkServiceImpl{
		repository: repository,
	}
}

func (s *WorkServiceImpl) Create(work models.Work) error {
	return s.repository.Create(work)
}

func (s *WorkServiceImpl) GetAll(userID uuid.UUID) ([]models.Work, error) {
	return s.repository.GetAll(userID)
}

func (s *WorkServiceImpl) Delete(id uuid.UUID) error {
	return s.repository.Delete(id)
}
