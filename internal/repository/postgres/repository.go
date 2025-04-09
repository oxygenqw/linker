package postgres

import (
	"database/sql"
)

type Repository struct {
	Student *StudentRepository
	Teacher *TeacherRepository
}

func NewPostgresRepository(db *sql.DB) *Repository {
	return &Repository{
		Student: NewStudentRepository(db),
		Teacher: NewTeacherRepository(db),
	}
}
