package repository

import (
	"fmt"

	"github.com/Oxygenss/linker/internal/config"
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository/postgres"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/google/uuid"
)

type Repository struct {
	logger  logger.Logger
	Student Student
	Teacher Teacher
	User    User
}

type Student interface {
	GetByTelegramID(telegramID int64) (models.Student, error)
	GetByID(id string) (models.Student, error)
	Create(student models.Student) (uuid.UUID, error)
	Update(student models.Student) error
	GetAll() ([]models.Student, error)
}

type Teacher interface {
	GetByTelegramID(telegramID int64) (models.Teacher, error)
	GetByID(id string) (models.Teacher, error)
	Update(teacher models.Teacher) error
	Create(teacher models.Teacher) (uuid.UUID, error)
	GetAll() ([]models.Teacher, error)
}

type User interface {
	GetRole(telegramID int64) (string, error)
}

func NewRepository(config *config.Config, logger *logger.Logger) (*Repository, error) {
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

	repository := postgres.NewPostgresRepository(db, logger)

	return &Repository{
		Student: repository.Student,
		Teacher: repository.Teacher,
		User:    repository.User,
	}, nil
}
