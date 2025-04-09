package repository

import (
	"fmt"

	"github.com/Oxygenss/linker/internal/config"
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository/postgres"
)

type Repository struct {
	Student Student
	Teacher Teacher
}

func NewRepository(config *config.Config) (*Repository, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
	)

	db, err := postgres.NewPostgresConnection(dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	repo := postgres.NewPostgresRepository(db)

	return &Repository{
		Student: repo.Student,
		Teacher: repo.Teacher,
	}, nil
}

type Student interface {
	GetByTelegramID(telegramID int64) (models.Student, error)
	GetAll() ([]models.Student, error)
	Create(student models.Student) error
}

type Teacher interface {
	GetByTelegramID(telegramID int64) (models.Teacher, error)
	GetAll() ([]models.Teacher, error)
	Create(teacher models.Teacher) error
}
