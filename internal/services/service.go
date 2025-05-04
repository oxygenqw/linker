package services

import (
	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/repository"
	"github.com/google/uuid"
)

type StudentService interface {
	Create(student models.Student) (uuid.UUID, error)
	GetByTelegramID(telegramID int64) (models.Student, error)
	GetByID(id string) (models.Student, error)
	GetAll() ([]models.Student, error)
	Search(search string) ([]models.Student, error)
	Update(student models.Student) error
	Delete(id string) error
}

type TeacherService interface {
	Create(teacher models.Teacher) (uuid.UUID, error)
	GetByTelegramID(telegramID int64) (models.Teacher, error)
	GetByID(id string) (models.Teacher, error)
	GetAll() ([]models.Teacher, error)
	Search(search string) ([]models.Teacher, error)
	Update(teacher models.Teacher) error
	Delete(id string) error
}

type Service struct {
	StudentService StudentService
	TeacherService TeacherService
	UserService    UserService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		StudentService: NewStudentService(repository.StudentRepository),
		TeacherService: NewTeacherService(repository.TeacherRepository),
		UserService:    NewUserService(repository.UserRepository),
	}
}
