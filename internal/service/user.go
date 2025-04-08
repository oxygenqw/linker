package service

import (
	"fmt"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
)

type UserService struct {
	repository repository.Repository
}

func NewUserService(repository repository.Repository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repository.GetAllUsers()
}

func (s *UserService) Create(user models.User) error {
	if s == nil {
		return fmt.Errorf("UserService is not initialized")
	}

	err := s.repository.AddUser(user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *UserService) GetByTelegramID(telegramID int64) (models.User, error) {
	fmt.Println("SERIVE ID", telegramID)
	user, err := s.repository.GetByTelegramID(telegramID)

	fmt.Println("SERVICE USER", user)
	return user, err
}
