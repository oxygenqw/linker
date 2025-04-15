package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Oxygenss/linker/pkg/logger"
)

type UserRepository struct {
	logger *logger.Logger
	db     *sql.DB
}

func NewUserRepository(db *sql.DB, logger *logger.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepository) GetRole(telegramID int64) (string, error) {
	r.logger.Info("[R: GetRole]")
	if r.db == nil {
		return "", fmt.Errorf("database connection is not initialized")
	}

	var studentFound bool
	studentQuery := `SELECT EXISTS(SELECT 1 FROM students WHERE telegram_id = $1)`
	err := r.db.QueryRow(studentQuery, telegramID).Scan(&studentFound)
	if err != nil {
		fmt.Println("studentQuery", err)
		return "", fmt.Errorf("failed to query students: %w", err)
	}

	if studentFound {
		return "student", nil
	}

	var teacherFound bool
	teacherQuery := `SELECT EXISTS(SELECT 1 FROM teachers WHERE telegram_id = $1)`
	err = r.db.QueryRow(teacherQuery, telegramID).Scan(&teacherFound)
	if err != nil {
		return "", fmt.Errorf("failed to query teachers: %w", err)
	}

	if teacherFound {
		return "teacher", nil
	}
	return "", nil
}
