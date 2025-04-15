package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/google/uuid"
)

type StudentRepository struct {
	logger *logger.Logger
	db     *sql.DB
}

func NewStudentRepository(db *sql.DB, logger *logger.Logger) *StudentRepository {
	return &StudentRepository{
		db:     db,
		logger: logger,
	}
}

func (r *StudentRepository) GetByID(id string) (models.Student, error) {
	if r.db == nil {
		return models.Student{}, fmt.Errorf("database connection is not initialized")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return models.Student{}, fmt.Errorf("invalid UUID format: %w", err)
	}

	query := `SELECT id, telegram_id, first_name, middle_name, last_name FROM students WHERE id = $1`

	var student models.Student
	err = r.db.QueryRow(query, id).Scan(
		&student.ID,
		&student.TelegramID,
		&student.FirstName,
		&student.LastName,
		&student.MiddleName,
	)

	switch {
	case err == nil:
		return student, nil
	case errors.Is(err, sql.ErrNoRows):
		return models.Student{}, fmt.Errorf("student not found: %w", err)
	default:
		return models.Student{}, fmt.Errorf("failed to get student: %w", err)
	}
}

func (r *StudentRepository) GetByTelegramID(telegramID int64) (models.Student, error) {
	r.logger.Info("[R: GetByTelegramID]")
	if r.db == nil {
		return models.Student{}, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, telegram_id, first_name, middle_name, last_name FROM students WHERE telegram_id = $1`

	var student models.Student
	err := r.db.QueryRow(query, telegramID).Scan(
		&student.ID,
		&student.TelegramID,
		&student.FirstName,
		&student.LastName,
		&student.MiddleName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Student{}, fmt.Errorf("student not found")
		}
		r.logger.Errorf("E: GetByTelegramID: %v", err)
		return models.Student{}, fmt.Errorf("failed to retrieve student: %w", err)
	}

	return student, nil
}

func (r *StudentRepository) GetAll() ([]models.Student, error) {
	query := `SELECT id, telegram_id, first_name, middle_name, last_name FROM students`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve students: %w", err)
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.TelegramID, &student.FirstName, &student.LastName, &student.MiddleName); err != nil {
			return nil, fmt.Errorf("failed to scan student: %w", err)
		}
		students = append(students, student)
	}

	return students, nil
}

func (r *StudentRepository) Create(student models.Student) (uuid.UUID, error) {
	if r.db == nil {
		return uuid.Nil, fmt.Errorf("database connection is not initialized")
	}

	student.ID = uuid.New()

	query := `INSERT INTO students (id, telegram_id, first_name, middle_name, last_name) 
              VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query,
		student.ID,
		student.TelegramID,
		student.FirstName,
		student.MiddleName,
		student.LastName)

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert student: %w", err)
	}

	return student.ID, nil
}
