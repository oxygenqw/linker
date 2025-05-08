package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/google/uuid"
)

type TeacherRepositoryImpl struct {
	logger *logger.Logger
	db     *sql.DB
}

func NewTeacherRepository(db *sql.DB, logger *logger.Logger) *TeacherRepositoryImpl {
	return &TeacherRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *TeacherRepositoryImpl) GetByID(id string) (models.Teacher, error) {
	r.logger.Info("[TeacherRepository: GetByID]", "id", id)

	if r.db == nil {
		return models.Teacher{}, fmt.Errorf("database connection is not initialized")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return models.Teacher{}, fmt.Errorf("invalid UUID format: %w", err)
	}

	query := `SELECT id, telegram_id, user_name, first_name, middle_name, last_name, 
                     degree, position, university, faculty, is_free, idea, about 
              FROM teachers WHERE id = $1`

	var teacher models.Teacher
	err = r.db.QueryRow(query, id).Scan(
		&teacher.ID,
		&teacher.TelegramID,
		&teacher.UserName,
		&teacher.FirstName,
		&teacher.MiddleName,
		&teacher.LastName,
		&teacher.Degree,
		&teacher.Position,
		&teacher.University,
		&teacher.Faculty,
		&teacher.IsFree,
		&teacher.Idea,
		&teacher.About,
	)

	switch {
	case err == nil:
		return teacher, nil
	case errors.Is(err, sql.ErrNoRows):
		return models.Teacher{}, fmt.Errorf("teacher not found: %w", err)
	default:
		r.logger.Error("Failed to get teacher by ID", "error", err, "id", id)
		return models.Teacher{}, fmt.Errorf("failed to get teacher: %w", err)
	}
}

func (r *TeacherRepositoryImpl) GetByTelegramID(telegramID int64) (models.Teacher, error) {
	r.logger.Info("[TeacherRepository: GetByTelegramID]", "telegramID", telegramID)

	if r.db == nil {
		return models.Teacher{}, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, telegram_id, user_name, first_name, middle_name, last_name,
                     degree, position, university, faculty, is_free, idea, about
              FROM teachers WHERE telegram_id = $1`

	var teacher models.Teacher
	err := r.db.QueryRow(query, telegramID).Scan(
		&teacher.ID,
		&teacher.TelegramID,
		&teacher.UserName,
		&teacher.FirstName,
		&teacher.MiddleName,
		&teacher.LastName,
		&teacher.Degree,
		&teacher.Position,
		&teacher.University,
		&teacher.Faculty,
		&teacher.IsFree,
		&teacher.Idea,
		&teacher.About,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Teacher{}, fmt.Errorf("teacher not found")
		}
		r.logger.Error("Failed to get teacher by Telegram ID", "error", err, "telegramID", telegramID)
		return models.Teacher{}, fmt.Errorf("failed to retrieve teacher: %w", err)
	}

	return teacher, nil
}

func (r *TeacherRepositoryImpl) GetAll() ([]models.Teacher, error) {
	r.logger.Info("[TeacherRepository: GetAll]")

	query := `SELECT id, telegram_id, user_name, first_name, middle_name, last_name,
                     degree, position, university, faculty, is_free, idea, about
              FROM teachers`

	rows, err := r.db.Query(query)
	if err != nil {
		r.logger.Error("Failed to retrieve teachers", "error", err)
		return nil, fmt.Errorf("failed to retrieve teachers: %w", err)
	}
	defer rows.Close()

	var teachers []models.Teacher
	for rows.Next() {
		var teacher models.Teacher
		err := rows.Scan(
			&teacher.ID,
			&teacher.TelegramID,
			&teacher.UserName,
			&teacher.FirstName,
			&teacher.MiddleName,
			&teacher.LastName,
			&teacher.Degree,
			&teacher.Position,
			&teacher.University,
			&teacher.Faculty,
			&teacher.IsFree,
			&teacher.Idea,
			&teacher.About,
		)
		if err != nil {
			r.logger.Error("Failed to scan teacher", "error", err)
			return nil, fmt.Errorf("failed to scan teacher: %w", err)
		}
		teachers = append(teachers, teacher)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("Error after scanning teachers", "error", err)
		return nil, fmt.Errorf("error after scanning teachers: %w", err)
	}

	return teachers, nil
}

func (r *TeacherRepositoryImpl) Search(query string) ([]models.Teacher, error) {
	r.logger.Info("[R: TeacherRepository: Search]")

	if r.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	sqlQuery := `
		SELECT id, first_name, middle_name, last_name, degree, university, faculty, is_free, idea
		FROM teachers
		WHERE
			last_name ILIKE '%' || $1 || '%' OR
			first_name ILIKE '%' || $1 || '%' OR
			middle_name ILIKE '%' || $1 || '%' OR
			degree ILIKE '%' || $1 || '%' OR
			university ILIKE '%' || $1 || '%' OR
			faculty ILIKE '%' || $1 || '%' OR
			idea ILIKE '%' || $1 || '%'
	`

	rows, err := r.db.Query(sqlQuery, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve teachers: %w", err)
	}
	defer rows.Close()

	var teachers []models.Teacher
	for rows.Next() {
		var teacher models.Teacher
		err := rows.Scan(
			&teacher.ID,
			&teacher.FirstName,
			&teacher.MiddleName,
			&teacher.LastName,
			&teacher.Degree,
			&teacher.University,
			&teacher.Faculty,
			&teacher.IsFree,
			&teacher.Idea,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan teacher: %w", err)
		}
		teachers = append(teachers, teacher)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after scanning rows: %w", err)
	}

	return teachers, nil
}

func (r *TeacherRepositoryImpl) Create(teacher models.Teacher) (uuid.UUID, error) {
	r.logger.Info("[TeacherRepository: Create]")

	if r.db == nil {
		return uuid.Nil, fmt.Errorf("database connection is not initialized")
	}

	if teacher.ID == uuid.Nil {
		teacher.ID = uuid.New()
	}

	query := `INSERT INTO teachers 
              (id, telegram_id, user_name, first_name, middle_name, last_name, 
               degree, position, university, faculty, is_free, idea, about) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err := r.db.Exec(query,
		teacher.ID,
		teacher.TelegramID,
		teacher.UserName,
		teacher.FirstName,
		teacher.MiddleName,
		teacher.LastName,
		teacher.Degree,
		teacher.Position,
		teacher.University,
		teacher.Faculty,
		teacher.IsFree,
		teacher.Idea,
		teacher.About,
	)

	if err != nil {
		r.logger.Error("Failed to create teacher", "error", err, "teacherID", teacher.ID)
		return uuid.Nil, fmt.Errorf("failed to insert teacher: %w", err)
	}

	return teacher.ID, nil
}

func (r *TeacherRepositoryImpl) Update(teacher models.Teacher) error {
	r.logger.Info("[TeacherRepository: Update]", "teacherID", teacher.ID)

	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	query := `
        UPDATE teachers 
        SET 
            first_name = $1,
            middle_name = $2,
            last_name = $3,
            degree = $4,
            position = $5,
            university = $6,
            faculty = $7,
            is_free = $8,
            idea = $9,
            about = $10
        WHERE id = $11
    `

	result, err := r.db.Exec(query,
		teacher.FirstName,
		teacher.MiddleName,
		teacher.LastName,
		teacher.Degree,
		teacher.Position,
		teacher.University,
		teacher.Faculty,
		teacher.IsFree,
		teacher.Idea,
		teacher.About,
		teacher.ID,
	)

	if err != nil {
		r.logger.Error("Failed to update teacher", "error", err, "teacherID", teacher.ID)
		return fmt.Errorf("failed to update teacher: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Error("Failed to get rows affected", "error", err)
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		r.logger.Warn("Teacher not found for update", "teacherID", teacher.ID)
		return fmt.Errorf("teacher with ID %s not found", teacher.ID)
	}

	return nil
}

func (r *TeacherRepositoryImpl) Delete(id string) error {
	r.logger.Info("[TeacherRepository: Delete]", id)

	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	query := `DELETE FROM teachers WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete teacher with id %s: %w", id, err)
	}

	return nil
}
