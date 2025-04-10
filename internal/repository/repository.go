package repository

import (
	"fmt"

	"github.com/Oxygenss/linker/internal/config"
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository/postgres"
	"github.com/Oxygenss/linker/pkg/logger"
)

type Repository struct {
	logger  logger.Logger
	Student Student
	Teacher Teacher
	User    User
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

type User interface {
	CheckByTelegramID(telegramID int64) (bool, error)
}

func New(config *config.Config, logger logger.Logger) (*Repository, error) {
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
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	repository := postgres.NewPostgresRepository(db)

	return &Repository{
		logger:  logger,
		Student: repository.Student,
		Teacher: repository.Teacher,
		User:    repository.User,
	}, nil
}
