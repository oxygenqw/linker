package service

import (
	"github.com/Oxygenss/linker/internal/repository"
)

type UserService struct {
	repository repository.User
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CheckByTelegramID(telegramID int64) (bool, error) {
	return s.repository.CheckByTelegramID(telegramID)
}
