package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Oxygenss/linker/internal/models"
	"github.com/Oxygenss/linker/internal/service"
	"github.com/Oxygenss/linker/pkg/telegram/bot"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/google/uuid"
)

type UserInfo struct {
	FirstName string
	LastName  string
	UserName  string
}

type Handler struct {
	service service.Service
	bot     *bot.Bot
}

func NewHandler(service service.Service, bot *bot.Bot) *Handler {
	return &Handler{service: service, bot: bot}
}

type TemplateData struct {
	UserName   string
	TelegramID string
}

type TemplateDataUsers struct {
	Users []models.User
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	telegramIDStr := r.URL.Query().Get("telegram_id")
	fmt.Println("telegramIDSTR", telegramIDStr)
	telegramID, err := strconv.ParseInt(telegramIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Telegram ID", http.StatusBadRequest)
		return
	}

	fmt.Println("TelegramINT:", telegramID)

	user, err := h.service.GetByTelegramID(telegramID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	if user.ID == uuid.Nil || user.TelegramID == 0 {
		log.Println("Нет пользователя")
		userName := r.URL.Query().Get("user_name")
		tmpl, err := template.ParseFiles("./ui/pages/login.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		data := TemplateData{
			UserName:   userName,
			TelegramID: telegramIDStr,
		}

		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	} else {
		log.Println("Есть пользователя")
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

func (h *Handler) Initialize(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	telegramIDStr := r.FormValue("telegram_id")
	telegramID, err := strconv.ParseInt(telegramIDStr, 10, 64)
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

	if err := h.service.AddUser(user); err != nil {
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

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	telegramIDStr := r.FormValue("telegram_id")
	users, err := h.service.GetAllUsers()
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

func (h *Handler) CreateBotEndpointHandler(appURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//log.Println("CreateBotEndpointHandler called")
		//log.Printf("Serving %s route", r.URL.Path)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//log.Printf("Request body: %s", string(body))

		var update gotgbot.Update
		if err := json.Unmarshal(body, &update); err != nil {
			log.Printf("Error decoding update: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//log.Printf("Received update: %+v", update)

		if update.Message == nil {
			log.Println("Received update without message")
			http.Error(w, "Bot update didn't include a message", http.StatusBadRequest)
			return
		}

		firstName := update.Message.From.FirstName
		lastName := update.Message.From.LastName
		userName := update.Message.From.Username
		telegramID := update.Message.From.Id
		//log.Printf("Received message: %s", update.Message.Text)

		log.Println("INFO", firstName, lastName, userName, telegramID)

		appURLWithParams := fmt.Sprintf("%s?first_name=%s&last_name=%s&user_name=%s&telegram_id=%s", appURL, firstName, lastName, userName, strconv.FormatInt(telegramID, 10))

		//log.Printf("WebApp URL: %s", appURLWithParams)

		message := "Welcome to the Telegram Mini App Template Bot"
		opts := &gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Open mini app", WebApp: &gotgbot.WebAppInfo{Url: appURLWithParams}},
				}, {}},
			},
		}

		//log.Printf("Sending message to chat ID: %d", update.Message.Chat.Id)
		if _, err := h.bot.SendMessage(update.Message.Chat.Id, message, opts); err != nil {
			log.Printf("Error sending message: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//log.Printf("Message sent to user: %s (ID: %d)", userName, update.Message.From.Id)
		w.WriteHeader(http.StatusOK)
	}
}
