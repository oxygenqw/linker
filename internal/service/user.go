package service

import (
	"github.com/Oxygenss/linker/internal/repository"
)

type User interface {
	GetRole(telegramID int64) (string, error)
}

type UserService struct {
	repository repository.User
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetRole(telegramID int64) (string, error) {
	return s.repository.GetRole(telegramID)
}
