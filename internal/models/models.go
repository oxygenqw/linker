package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID
	TelegramID int64
	FirstName  string
	LastName   string
	SureName   string
}
