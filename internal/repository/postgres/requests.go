package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/pkg/logger"
	"github.com/google/uuid"
)

type RequestRepositoryImpl struct {
	logger *logger.Logger
	db     *sql.DB
}

func NewRequestRepository(db *sql.DB, logger *logger.Logger) *RequestRepositoryImpl {
	return &RequestRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *RequestRepositoryImpl) Create(request models.Request) error {
	r.logger.Info("[R: RequestRepositoryImpl: Create]")

	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	if request.ID == uuid.Nil {
		request.ID = uuid.New()
	}

	if request.CreatedAt.IsZero() {
		request.CreatedAt = time.Now()
	}

	query := `INSERT INTO requests 
    (id, sender_id, recipient_id, message, status, created_at)
    VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.Exec(query,
		request.ID,
		request.SenderID,
		request.RecipientID,
		request.Message,
		request.Status,
		request.CreatedAt)

	if err != nil {
		r.logger.Error("Failed to create request", "error", err, "requestID", request.ID)
		return fmt.Errorf("failed to insert request: %w", err)
	}

	return nil
}
