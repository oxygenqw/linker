package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/google/uuid"
)

type StudentRepositoryImpl struct {
	logger *logger.Logger
	db     *sql.DB
}

func NewStudentRepository(db *sql.DB, logger *logger.Logger) *StudentRepositoryImpl {
	return &StudentRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *StudentRepositoryImpl) GetByID(id string) (models.Student, error) {
	r.logger.Info("[StudentRepository: GetByID]")

	if r.db == nil {
		return models.Student{}, fmt.Errorf("database connection is not initialized")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return models.Student{}, fmt.Errorf("invalid UUID format: %w", err)
	}

	query := `SELECT id, telegram_id, user_name, first_name, middle_name, last_name, github, job, idea, about, 
              university, faculty, course, education FROM students WHERE id = $1`

	var student models.Student
	err = r.db.QueryRow(query, id).Scan(
		&student.ID,
		&student.TelegramID,
		&student.UserName,
		&student.FirstName,
		&student.MiddleName,
		&student.LastName,
		&student.GitHub,
		&student.Job,
		&student.Idea,
		&student.About,
		&student.University,
		&student.Faculty,
		&student.Course,
		&student.Education,
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

func (r *StudentRepositoryImpl) GetByTelegramID(telegramID int64) (models.Student, error) {
	r.logger.Info("[StudentRepository: GetByTelegramID]")

	if r.db == nil {
		return models.Student{}, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, telegram_id, user_name, first_name, middle_name, last_name, github, job, idea, about,
              university, faculty, course, education FROM students WHERE telegram_id = $1`

	var student models.Student
	err := r.db.QueryRow(query, telegramID).Scan(
		&student.ID,
		&student.TelegramID,
		&student.UserName,
		&student.FirstName,
		&student.MiddleName,
		&student.LastName,
		&student.GitHub,
		&student.Job,
		&student.Idea,
		&student.About,
		&student.University,
		&student.Faculty,
		&student.Course,
		&student.Education,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Student{}, fmt.Errorf("student not found")
		}
		return models.Student{}, fmt.Errorf("failed to retrieve student: %w", err)
	}

	return student, nil
}

func (r *StudentRepositoryImpl) GetAll() ([]models.Student, error) {
	r.logger.Info("[StudentRepository: GetAll]")

	if r.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, telegram_id, user_name, first_name, middle_name, last_name, github, job, idea, about,
              university, faculty, course, education FROM students`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve students: %w", err)
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		err := rows.Scan(
			&student.ID,
			&student.TelegramID,
			&student.UserName,
			&student.FirstName,
			&student.MiddleName,
			&student.LastName,
			&student.GitHub,
			&student.Job,
			&student.Idea,
			&student.About,
			&student.University,
			&student.Faculty,
			&student.Course,
			&student.Education,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan student: %w", err)
		}

		students = append(students, student)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error after scanning rows: %w", err)
	}

	return students, nil
}

func (r *StudentRepositoryImpl) Create(student models.Student) (uuid.UUID, error) {
	r.logger.Info("[StudentRepository: Create]")

	if r.db == nil {
		return uuid.Nil, fmt.Errorf("database connection is not initialized")
	}

	student.ID = uuid.New()

	query := `INSERT INTO students 
              (id, telegram_id, user_name, first_name, middle_name, last_name, github, job, idea, about,
               university, faculty, course, education) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	_, err := r.db.Exec(query,
		student.ID,
		student.TelegramID,
		student.UserName,
		student.FirstName,
		student.MiddleName,
		student.LastName,
		student.GitHub,
		student.Job,
		student.Idea,
		student.About,
		student.University,
		student.Faculty,
		student.Course,
		student.Education,
	)

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert student: %w", err)
	}

	return student.ID, nil
}

func (r *StudentRepositoryImpl) Update(student models.Student) error {
	r.logger.Info("[StudentRepository: Update]", "studentID", student.ID)

	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	query := `
	UPDATE students 
	SET 
		first_name = $1,
		middle_name = $2,
		last_name = $3,
		github = $4,
		job = $5,
		idea = $6,
		about = $7,
		university = $8,
		faculty = $9,
		course = $10,
		education = $11
	WHERE id = $12
	`

	result, err := r.db.Exec(query,
		student.FirstName,
		student.MiddleName,
		student.LastName,
		student.GitHub,
		student.Job,
		student.Idea,
		student.About,
		student.University,
		student.Faculty,
		student.Course,
		student.Education,
		student.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update student: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("student with ID %s not found", student.ID)
	}

	return nil
}

func (r *StudentRepositoryImpl) Delete(id string) error {
	r.logger.Info("[StudentRepository: Delete]", id)

	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	query := `DELETE FROM students WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete student with id %s: %w", id, err)
	}

	return nil
}