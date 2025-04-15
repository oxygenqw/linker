package service

import (
	"github.com/Oxygenss/linker/internal/repository"
)

type Service struct {
	Student
	Teacher
	User
}

func New(repository *repository.Repository) *Service {
	return &Service{
		Student: NewStudentService(repository.Student),
		Teacher: NewTeacherService(repository.Teacher),
		User:    NewUserService(repository.User),
	}
}
