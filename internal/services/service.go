package services

import (
	"github.com/Oxygenss/linker/internal/repository"
)

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
