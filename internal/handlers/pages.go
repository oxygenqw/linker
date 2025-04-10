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

func (h *PagesHandler) Profile(w http.ResponseWriter, r *http.Request) {
    telegramIDStr := r.URL.Query().Get("telegram_id")
    telegramID, err := strconv.ParseInt(telegramIDStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid Telegram ID", http.StatusBadRequest)
        return
    }

    var userFound bool
    var data map[string]any

    student, err := h.service.Student.GetByTelegramID(telegramID)
    if err == nil {
        userFound = true
        data = map[string]any{
            "user": student,
        }
    } else if err != sql.ErrNoRows {
        http.Error(w, "Error retrieving user", http.StatusInternalServerError)
        return
    }

    if !userFound {
        teacher, err := h.service.Teacher.GetByTelegramID(telegramID)
        if err != nil && err != sql.ErrNoRows {
            http.Error(w, "Error retrieving user", http.StatusInternalServerError)
            return
        }

        if err == nil {
            data = map[string]any{
                "user": teacher,
            }
        }
    }

    if data != nil {
        tmpl, err := template.ParseFiles("./ui/pages/profile.html")
        if err != nil {
            http.Error(w, "Error loading template", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "text/html")
        if err := tmpl.Execute(w, data); err != nil {
            http.Error(w, "Error executing template", http.StatusInternalServerError)
            return
        }
    } else {
        http.Error(w, "User not found", http.StatusNotFound)
    }
}


func (h *PagesHandler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	telegramIDStr := r.URL.Query().Get("telegram_id")
	role := r.URL.Query().Get("role")
	fmt.Println("Home ", telegramIDStr, role)
	tmpl, err := template.ParseFiles("./ui/pages/home.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"telegram_id": telegramIDStr,
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

// При запуске приложения пытаемся найти пользователя или преподавателя, если не нашли то открываем login, иначе редиректим в home
// @router GET /
func (h *PagesHandler) Input(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Input")
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
	fmt.Println("NewUser")
	telegramID, err := strconv.ParseInt(r.FormValue("telegram_id"), 10, 64)
	if err != nil {
		fmt.Println("TYTYYT2")
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

func (h *PagesHandler) Students(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Students")
	tmpl, err := template.ParseFiles("./ui/pages/students.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	students, err := h.service.Student.GetAll()
	if err != nil {
		http.Error(w, "Ошибка при получении списка студентов", http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"students": students,
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func (h *PagesHandler) Teachers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Teachers")
	tmpl, err := template.ParseFiles("./ui/pages/teachers.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	teachers, err := h.service.Teacher.GetAll()
	if err != nil {
		http.Error(w, "Ошибка при получении списка преподавателей", http.StatusInternalServerError)
		return
	}

	fmt.Println("LEN TEA", len(teachers))

	data := map[string]any{
		"teachers": teachers,
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
