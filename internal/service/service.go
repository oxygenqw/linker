package service

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
)

type User interface {
	AddUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetByTelegramID(telegramID int64) (models.User, error)
}

type Service struct {
	User
}

func NewService(repository repository.Repository) *Service {
	return &Service{User: NewUserService(repository)}
}
