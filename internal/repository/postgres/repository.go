package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllUsers() ([]models.User, error) {
	query := `SELECT id, telegram_id, first_name, last_name, sure_name FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.TelegramID, &user.FirstName, &user.LastName, &user.SureName); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *Repository) AddUser(user models.User) error {
	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	user.ID = uuid.New()

	query := `INSERT INTO users (id, telegram_id, first_name, last_name, sure_name) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, user.ID, user.TelegramID, user.FirstName, user.LastName, user.SureName)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}
