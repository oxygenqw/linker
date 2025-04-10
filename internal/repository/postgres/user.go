package postgres

import (
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CheckByTelegramID(telegramID int64) (bool, error) {
	if r.db == nil {
		return false, fmt.Errorf("database connection is not initialized")
	}

	var studentFound bool
	studentQuery := `SELECT EXISTS(SELECT 1 FROM students WHERE telegram_id = $1)`
	err := r.db.QueryRow(studentQuery, telegramID).Scan(&studentFound)
	if err != nil {
		return false, fmt.Errorf("failed to query students: %w", err)
	}

	if studentFound {
		return true, nil
	}

	var teacherFound bool
	teacherQuery := `SELECT EXISTS(SELECT 1 FROM teachers WHERE telegram_id = $1)`
	err = r.db.QueryRow(teacherQuery, telegramID).Scan(&teacherFound)
	if err != nil {
		return false, fmt.Errorf("failed to query teachers: %w", err)
	}

	return teacherFound, nil
}
