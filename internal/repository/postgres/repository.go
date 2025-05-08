package postgres

import (
	"database/sql"

	"github.com/Oxygenss/linker/pkg/logger"
)

type Repository struct {
	StudentRepository *StudentRepositoryImpl
	TeacherRepository *TeacherRepositoryImpl
	UserRepository    *UserRepositoryImpl
	RequestRepository *RequestRepositoryImpl
}

func NewPostgresRepository(db *sql.DB, logger *logger.Logger) *Repository {
	return &Repository{
		StudentRepository: NewStudentRepository(db, logger),
		TeacherRepository: NewTeacherRepository(db, logger),
		UserRepository:    NewUserRepository(db, logger),
		RequestRepository: NewRequestRepository(db, logger),
	}
}
