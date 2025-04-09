package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/google/uuid"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) GetByTelegramID(telegramID int64) (models.Student, error) {
	if r.db == nil {
		return models.Student{}, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, telegram_id, first_name, last_name, middle_name FROM students WHERE telegram_id = $1`
	var student models.Student
	err := r.db.QueryRow(query, telegramID).Scan(&student.ID, &student.TelegramID, &student.FirstName, &student.LastName, &student.MiddleName)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Студент не найден")
			return models.Student{}, err // Возвращаем пустого пользователя без ошибки
		}
		fmt.Println("Ошибка при получении пользователя:", err)
		return models.Student{}, fmt.Errorf("failed to retrieve student: %w", err)
	}

	return student, nil
}

func (r *StudentRepository) GetAll() ([]models.Student, error) {
	query := `SELECT id, telegram_id, first_name, last_name, middle_name FROM students`
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

func (r *StudentRepository) Create(student models.Student) error {
	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	student.ID = uuid.New()

	query := `INSERT INTO students (id, telegram_id, first_name, last_name, middle_name) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, student.ID, student.TelegramID, student.FirstName, student.LastName, student.MiddleName)
	if err != nil {
		return fmt.Errorf("failed to insert student: %w", err)
	}

	return nil
}
