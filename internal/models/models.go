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
	University string
	Faculty    string
	Course     string
	Education  string
}

type Teacher struct {
    ID         uuid.UUID
    TelegramID int64
    FirstName  string
    LastName   string
    MiddleName string
    Degree     string
    Position   string
    University string
    Faculty    string
    IsFree     bool      
    Idea       string 
    About      string
}
