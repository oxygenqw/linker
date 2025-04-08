package service

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
)

type Users interface {
	Create(user models.User) error
	GetAll() ([]models.User, error)
	GetByTelegramID(telegramID int64) (models.User, error)
}

type Service struct {
	Users
}

func NewService(repository repository.Repository) *Service {
	return &Service{Users: NewUserService(repository)}
}
