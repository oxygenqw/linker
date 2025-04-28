package services

import (
	"github.com/Oxygenss/linker/internal/repository"
)

type UserService interface {
	GetRole(telegramID int64) (string, error)
}

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (s *UserServiceImpl) GetRole(telegramID int64) (string, error) {
	return s.repository.GetRole(telegramID)
}
