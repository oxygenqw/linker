package service

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
		StudentService: NewStudentService(repository.StudentRepositiry),
		TeacherService: NewTeacherService(repository.TeacherRepository),
		UserService:    NewUserService(repository.UserRepository),
	}
}
