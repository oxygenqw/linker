package models

import "github.com/google/uuid"

type Student struct {
	ID         uuid.UUID
	TelegramID int64
	FirstName  string
	LastName   string
	MiddleName string
}

type Teacher struct {
	ID         uuid.UUID
	TelegramID int64
	FirstName  string
	LastName   string
	MiddleName string
}
