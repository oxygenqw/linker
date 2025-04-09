package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/service"
)

type PagesHandler struct {
	service service.Service
}

func NewPagesHandler(service *service.Service) *PagesHandler {
	return &PagesHandler{
		service: *service,
	}
}

func (h *PagesHandler) Home(w http.ResponseWriter, r *http.Request) {
	telegramIDStr := r.URL.Query().Get("telegram_id")
	role := r.URL.Query().Get("role")
	fmt.Println("Home ", telegramIDStr, role)
	// Обработка и отображение информации на основе Telegram ID
}

// При запуске приложения пытаемся найти пользователя или преподавателя, если не нашли то открываем login, иначе редиректим в home
// @router GET /
func (h *PagesHandler) Input(w http.ResponseWriter, r *http.Request) {
	log.Println("Input URL:", r.URL)
	telegramIDStr := r.URL.Query().Get("telegram_id")

	telegramID, err := strconv.ParseInt(telegramIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Telegram ID", http.StatusBadRequest)
		return
	}

	var role string
	var userFound bool
	_, err = h.service.Student.GetByTelegramID(telegramID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	if err != sql.ErrNoRows {
		userFound = true
		role = "student"
	}

	_, err = h.service.Teacher.GetByTelegramID(telegramID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	if err != sql.ErrNoRows {
		userFound = true
		role = "teacher"
	}

	if !userFound {
		userName := r.URL.Query().Get("user_name")
		tmpl, err := template.ParseFiles("./ui/pages/login.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		data := map[string]any{
			"UserName":   userName,
			"TelegramID": telegramIDStr,
		}

		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	} else {
		http.Redirect(w, r, fmt.Sprintf("/home?telegram_id=%s&role=%s", telegramIDStr, role), http.StatusFound)
	}
}

// После заполнения первоначальных данных парсим форму и переходим в профиль
// @router POST /users
func (h *PagesHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	telegramID, err := strconv.ParseInt(r.FormValue("telegram_id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid Telegram ID", http.StatusBadRequest)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	role := r.FormValue("role")

	switch role {
	case "student":
		student := models.Student{
			TelegramID: telegramID,
			FirstName:  r.FormValue("first_name"),
			LastName:   r.FormValue("last_name"),
			MiddleName: r.FormValue("middle_name"),
		}

		err := h.service.Student.Create(student)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add user: %v", err), http.StatusInternalServerError)
		}
	case "teacher":
		teacher := models.Teacher{
			TelegramID: telegramID,
			FirstName:  r.FormValue("first_name"),
			LastName:   r.FormValue("last_name"),
			MiddleName: r.FormValue("middle_name"),
		}
		err := h.service.Teacher.Create(teacher)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add user: %v", err), http.StatusInternalServerError)
		}
	}

	http.Redirect(w, r, fmt.Sprintf("/home?telegram_id=%s", strconv.FormatInt(telegramID, 10)), http.StatusFound)
}

// func (h *PagesHandler) List(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "Failed to parse form", http.StatusBadRequest)
// 		return
// 	}

// 	telegramIDStr := r.FormValue("telegram_id")
// 	users, err := h.service.Users.GetAll()
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to retrieve users: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	tmpl, err := template.ParseFiles("./ui/pages/list.html")
// 	if err != nil {
// 		http.Error(w, "Error loading template", http.StatusInternalServerError)
// 		return
// 	}

// 	data := struct {
// 		Users       []models.User
// 		CurrentUser struct {
// 			TelegramID string
// 		}
// 	}{
// 		Users: users,
// 		CurrentUser: struct {
// 			TelegramID string
// 		}{TelegramID: telegramIDStr},
// 	}

// 	if err := tmpl.Execute(w, data); err != nil {
// 		http.Error(w, "Error executing template", http.StatusInternalServerError)
// 		return
// 	}
// }
