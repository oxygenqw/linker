package repository

import (
	"fmt"

	"github.com/Oxygenss/linker/internal/config"
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository/postgres"
)

type Repository interface {
	AddUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetByTelegramID(telegramID int64) (models.User, error)
}

func New(config *config.Config) (Repository, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
	)

	db, err := postgres.NewPostgresDB(dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	repo := postgres.New(db)
	return repo, nil
}
