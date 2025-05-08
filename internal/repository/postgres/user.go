package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Oxygenss/linker/pkg/logger"
)

type UserRepositoryImpl struct {
	logger *logger.Logger
	db     *sql.DB
}

func NewUserRepository(db *sql.DB, logger *logger.Logger) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepositoryImpl) GetRoleByTelegramID(telegramID int64) (string, error) {
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

func (r *UserRepositoryImpl) GetRoleByID(id string) (string, error) {
	r.logger.Info("[R: GetRoleByID]")
	if r.db == nil {
		return "", fmt.Errorf("database connection is not initialized")
	}

	var studentFound bool
	studentQuery := `SELECT EXISTS(SELECT 1 FROM students WHERE id = $1)`
	err := r.db.QueryRow(studentQuery, id).Scan(&studentFound)
	if err != nil {
		r.logger.Error("Failed to query students by ID", "error", err)
		return "", fmt.Errorf("failed to query students: %w", err)
	}

	if studentFound {
		return "student", nil
	}

	var teacherFound bool
	teacherQuery := `SELECT EXISTS(SELECT 1 FROM teachers WHERE id = $1)`
	err = r.db.QueryRow(teacherQuery, id).Scan(&teacherFound)
	if err != nil {
		r.logger.Error("Failed to query teachers by ID", "error", err)
		return "", fmt.Errorf("failed to query teachers: %w", err)
	}

	if teacherFound {
		return "teacher", nil
	}

	return "", nil
}
