package services

import (
	"github.com/Oxygenss/linker/internal/repository"
)

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (s *UserServiceImpl) GetRoleByID(id string) (string, error) {
	return s.repository.GetRoleByID(id)
}

func (s *UserServiceImpl) GetRoleByTelegramID(telegramID int64) (string, error) {
	return s.repository.GetRoleByTelegramID(telegramID)
}
