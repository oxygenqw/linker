package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/google/uuid"
)

type TeacherRepository struct {
	logger *logger.Logger
	db     *sql.DB
}

func NewTeacherRepository(db *sql.DB, logger *logger.Logger) *TeacherRepository {
	return &TeacherRepository{
		db:     db,
		logger: logger,
	}
}

func (r *TeacherRepository) GetByID(id string) (models.Teacher, error) {
	if r.db == nil {
		return models.Teacher{}, fmt.Errorf("database connection is not initialized")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return models.Teacher{}, fmt.Errorf("invalid UUID format: %w", err)
	}

	query := `SELECT id, telegram_id, first_name, middle_name, last_name FROM teachers WHERE id = $1`

	var teacher models.Teacher
	err = r.db.QueryRow(query, id).Scan(
		&teacher.ID,
		&teacher.TelegramID,
		&teacher.FirstName,
		&teacher.LastName,
		&teacher.MiddleName,
	)

	switch {
	case err == nil:
		return teacher, nil
	case errors.Is(err, sql.ErrNoRows):
		return models.Teacher{}, fmt.Errorf("teacher not found: %w", err)
	default:
		return models.Teacher{}, fmt.Errorf("failed to get teacher: %w", err)
	}
}

func (r *TeacherRepository) GetByTelegramID(telegramID int64) (models.Teacher, error) {
	r.logger.Info("[R: GetByTelegramID]")
	if r.db == nil {
		return models.Teacher{}, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, telegram_id, first_name, middle_name, last_name FROM teachers WHERE telegram_id = $1`

	var teacher models.Teacher
	err := r.db.QueryRow(query, telegramID).Scan(
		&teacher.ID,
		&teacher.TelegramID,
		&teacher.FirstName,
		&teacher.LastName,
		&teacher.MiddleName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Teacher{}, fmt.Errorf("teacher not found")
		}
		r.logger.Errorf("E: GetByTelegramID: %v", err)
		return models.Teacher{}, fmt.Errorf("failed to retrieve student: %w", err)
	}

	return teacher, nil
}

func (r *TeacherRepository) GetAll() ([]models.Teacher, error) {
	query := `SELECT id, telegram_id, first_name, middle_name, last_name FROM teachers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve teachers: %w", err)
	}
	defer rows.Close()

	var teachers []models.Teacher
	for rows.Next() {
		var teacher models.Teacher
		if err := rows.Scan(&teacher.ID, &teacher.TelegramID, &teacher.FirstName, &teacher.LastName, &teacher.MiddleName); err != nil {
			return nil, fmt.Errorf("failed to scan teacher: %w", err)
		}
		teachers = append(teachers, teacher)
	}

	return teachers, nil
}

func (r *TeacherRepository) Create(teacher models.Teacher) (uuid.UUID, error) {
	if r.db == nil {
		return uuid.Nil, fmt.Errorf("database connection is not initialized")
	}

	teacher.ID = uuid.New()

	query := `INSERT INTO teachers (id, telegram_id, first_name, middle_name, last_name) 
              VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query,
		teacher.ID,
		teacher.TelegramID,
		teacher.FirstName,
		teacher.MiddleName,
		teacher.LastName)

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert teacher: %w", err)
	}

	return teacher.ID, nil
}
