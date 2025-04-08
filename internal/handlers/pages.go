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
	"github.com/google/uuid"
)

type PagesHandler struct {
	service service.Service
}

func NewPagesHandler(service *service.Service) *PagesHandler {
	return &PagesHandler{
		service: *service,
	}
}

func (h *PagesHandler) Initialize(w http.ResponseWriter, r *http.Request) {
	log.Println("Initialize URL:", r.URL)
	telegramIDStr := r.URL.Query().Get("telegram_id")

	telegramID, err := strconv.ParseInt(telegramIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Telegram ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetByTelegramID(telegramID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	if user.ID == uuid.Nil || user.TelegramID == 0 {
		userName := r.URL.Query().Get("user_name")
		tmpl, err := template.ParseFiles("./ui/pages/login.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		data := models.TemplateData{
			UserName:   userName,
			TelegramID: telegramIDStr,
		}

		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	} else {
		tmpl, err := template.ParseFiles("./ui/pages/profile.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, user); err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	}
}

// После заполнения первоначальных данных парсим форму и переходим в профиль
// @router POST /users
func (h *PagesHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	log.Println("SaveInfo URL", r.URL.Path)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	telegramID, err := strconv.ParseInt(r.FormValue("telegram_id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid Telegram ID", http.StatusBadRequest)
		return
	}

	user := models.User{
		TelegramID: telegramID,
		FirstName:  r.FormValue("first_name"),
		LastName:   r.FormValue("last_name"),
		SureName:   r.FormValue("surename"),
	}

	if err := h.service.Users.Create(user); err != nil {
		http.Error(w, fmt.Sprintf("Failed to add user: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./ui/pages/profile.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, user); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func (h *PagesHandler) List(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	telegramIDStr := r.FormValue("telegram_id")
	users, err := h.service.Users.GetAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve users: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./ui/pages/list.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Users       []models.User
		CurrentUser struct {
			TelegramID string
		}
	}{
		Users: users,
		CurrentUser: struct {
			TelegramID string
		}{TelegramID: telegramIDStr},
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
