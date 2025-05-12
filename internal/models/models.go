package models

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID         uuid.UUID `json:"id"`
	TelegramID int64     `json:"telegram_id"`
	UserName   string    `json:"user_name"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	MiddleName string    `json:"middle_name"`
	University string    `json:"university"`
	Faculty    string    `json:"faculty"`
	Specialty  string    `json:"specialty"`
	Idea       string    `json:"idea"`
	About      string    `json:"about"`
	GitHub     string    `json:"github"`
	Job        string    `json:"job"`
	Course     string    `json:"course"`
	Education  string    `json:"education"`
}

type Teacher struct {
	ID         uuid.UUID `json:"id"`
	TelegramID int64     `json:"telegram_id"`
	UserName   string    `json:"user_name"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	MiddleName string    `json:"middle_name"`
	University string    `json:"university"`
	Faculty    string    `json:"faculty"`
	Idea       string    `json:"idea"`
	About      string    `json:"about"`
	Degree     string    `json:"degree"`
	Position   string    `json:"position"`
	IsFree     bool      `json:"is_free"`
}

type Request struct {
	ID          uuid.UUID `json:"id"`
	SenderID    uuid.UUID `json:"sender_id"`
	RecipientID uuid.UUID `json:"recipient_id"`
	Message     string    `json:"message"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
}

type Work struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
}

type Message struct {
	Message string `json:"message"`
}
