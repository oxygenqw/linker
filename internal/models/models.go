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

type TemplateData struct {
	UserName   string
	TelegramID string
}

// type TemplateDataUsers struct {
// 	Users []User
// }

type UserInfo struct {
	FirstName string
	LastName  string
	UserName  string
}
