package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/google/uuid"
)

type TeacherRepository struct {
	db *sql.DB
}

func NewTeacherRepository(db *sql.DB) *TeacherRepository {
	return &TeacherRepository{db: db}
}

func (r *TeacherRepository) GetByTelegramID(telegramID int64) (models.Teacher, error) {
	if r.db == nil {
		return models.Teacher{}, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, telegram_id, first_name, middle_name, last_name FROM teachers WHERE telegram_id = $1`
	var Teacher models.Teacher
	err := r.db.QueryRow(query, telegramID).Scan(&Teacher.ID, &Teacher.TelegramID, &Teacher.FirstName, &Teacher.LastName, &Teacher.MiddleName)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Преподаватель не найден")
			return models.Teacher{}, err
		}
		fmt.Println("Ошибка при получении пользователя:", err)
		return models.Teacher{}, fmt.Errorf("failed to retrieve Teacher: %w", err)
	}

	return Teacher, nil
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

func (r *TeacherRepository) Create(teacher models.Teacher) error {
	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	teacher.ID = uuid.New()

	query := `INSERT INTO teachers (id, telegram_id, first_name, middle_name, last_name) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, teacher.ID, teacher.TelegramID, teacher.FirstName, teacher.LastName, teacher.MiddleName)
	if err != nil {
		return fmt.Errorf("failed to insert teacher: %w", err)
	}

	return nil
}
