package models

import "github.com/google/uuid"

type Student struct {
	ID         uuid.UUID
	TelegramID int64
	FirstName  string
	LastName   string
	MiddleName string
	University string
	Faculty    string
	Idea       string
	About      string

	GitHub    string
	Job       string
	Course    string
	Education string
}

type Teacher struct {
	ID         uuid.UUID
	TelegramID int64
	FirstName  string
	LastName   string
	MiddleName string
	University string
	Faculty    string
	Idea       string
	About      string

	Degree   string
	Position string
	IsFree   bool
}
