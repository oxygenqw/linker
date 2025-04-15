package postgres

import (
	"database/sql"

	"github.com/Oxygenss/linker/pkg/logger"
)

type Repository struct {
	Student *StudentRepository
	Teacher *TeacherRepository
	User    *UserRepository
}

func NewPostgresRepository(db *sql.DB, logger *logger.Logger) *Repository {
	return &Repository{
		Student: NewStudentRepository(db, logger),
		Teacher: NewTeacherRepository(db, logger),
		User:    NewUserRepository(db, logger),
	}
}
