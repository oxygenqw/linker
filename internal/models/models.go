package models

import "github.com/google/uuid"

type Student struct {
	ID         uuid.UUID
	TelegramID int64
	FirstName  string
	LastName   string
	MiddleName string
	GitHub     string
	Job        string
	Idea       string
	About      string
}

type Teacher struct {
	ID         uuid.UUID
	TelegramID int64
	FirstName  string
	LastName   string
	MiddleName string
	Degree     string
	Position   string
	Department string
	IsFree     bool
	Idea       string
	About      string
}
