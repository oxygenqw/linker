package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/Oxygenss/linker/pkg/telegram/bot"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

type UserInfo struct {
	FirstName string
	LastName  string
	UserName  string
}

type Handler struct {
	bot *bot.Bot
}

func NewHandler(bot *bot.Bot) *Handler {
	return &Handler{bot: bot}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home/home.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

// func (h *Handler) Welcome(w http.ResponseWriter, r *http.Request) {
// 	userInfo := UserInfo{
// 		FirstName: r.URL.Query().Get("first_name"),
// 		LastName:  r.URL.Query().Get("last_name"),
// 		UserName:  r.URL.Query().Get("user_name"),
// 	}

// 	tmpl, err := template.ParseFiles("templates/index.html")
// 	if err != nil {
// 		http.Error(w, "Error loading template", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "text/html")
// 	if err := tmpl.Execute(w, userInfo); err != nil {
// 		http.Error(w, "Error executing template", http.StatusInternalServerError)
// 		return
// 	}
// }

func (h *Handler) CreateBotEndpointHandler(appURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("CreateBotEndpointHandler called")
		log.Printf("Serving %s route", r.URL.Path)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Request body: %s", string(body))

		var update gotgbot.Update
		if err := json.Unmarshal(body, &update); err != nil {
			log.Printf("Error decoding update: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Received update: %+v", update)

		if update.Message == nil {
			log.Println("Received update without message")
			http.Error(w, "Bot update didn't include a message", http.StatusBadRequest)
			return
		}

		firstName := update.Message.From.FirstName
		lastName := update.Message.From.LastName
		userName := update.Message.From.Username
		log.Printf("Received message: %s", update.Message.Text)

		appURLWithParams := fmt.Sprintf("%s?first_name=%s&last_name=%s&user_name=%s", appURL, firstName, lastName, userName)
		log.Printf("WebApp URL: %s", appURLWithParams)

		message := "Welcome to the Telegram Mini App Template Bot"
		opts := &gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Open mini app", WebApp: &gotgbot.WebAppInfo{Url: appURLWithParams}},
				}, {}},
			},
		}

		log.Printf("Sending message to chat ID: %d", update.Message.Chat.Id)
		if _, err := h.bot.SendMessage(update.Message.Chat.Id, message, opts); err != nil {
			log.Printf("Error sending message: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Message sent to user: %s (ID: %d)", userName, update.Message.From.Id)
		w.WriteHeader(http.StatusOK)
	}
}
