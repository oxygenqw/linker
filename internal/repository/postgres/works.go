package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/google/uuid"
)

type WorksRepositoryImpl struct {
	logger *logger.Logger
	db     *sql.DB
}

func NewWorksRepository(db *sql.DB, logger *logger.Logger) *WorksRepositoryImpl {
	return &WorksRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *WorksRepositoryImpl) Create(work models.Work) error {
	r.logger.Info("[R: WorksRepositoryImpl: Create]")

	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	if work.ID == uuid.Nil {
		work.ID = uuid.New()
	}

	query := `INSERT INTO works (id, user_id, link, description) VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(query,
		work.ID,
		work.UserID,
		work.Link,
		work.Description,
	)

	if err != nil {
		r.logger.Error("Failed to create work", "error", err, "workID", work.ID)
		return fmt.Errorf("failed to insert work: %w", err)
	}

	return nil
}

func (r *WorksRepositoryImpl) GetAll(userID uuid.UUID) ([]models.Work, error) {
	r.logger.Info("[R: WorksRepositoryImpl: GetAll]", "userID", userID)

	if r.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, user_id, link, description FROM works WHERE user_id = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		r.logger.Error("Failed to get works", "error", err, "userID", userID)
		return nil, fmt.Errorf("failed to get works: %w", err)
	}
	defer rows.Close()

	var works []models.Work
	for rows.Next() {
		var work models.Work
		if err := rows.Scan(&work.ID, &work.UserID, &work.Link, &work.Description); err != nil {
			r.logger.Error("Failed to scan work", "error", err)
			continue
		}
		works = append(works, work)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Rows error in GetAll", "error", err)
		return nil, err
	}

	return works, nil
}

func (r *WorksRepositoryImpl) Delete(id uuid.UUID) error {
	r.logger.Info("[R: WorksRepositoryImpl: Delete]", "workID", id)

	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	query := `DELETE FROM works WHERE id = $1`
	res, err := r.db.Exec(query, id)
	if err != nil {
		r.logger.Error("Failed to delete work", "error", err, "workID", id)
		return fmt.Errorf("failed to delete work: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err == nil && rowsAffected == 0 {
		return fmt.Errorf("work with id %s not found", id)
	}

	return err
}
