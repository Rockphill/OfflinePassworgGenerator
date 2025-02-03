package handlers

import (
	"back/utils"
	"html/template"
	"net/http"
	"strconv"
)

// HomeHandler обрабатывает главную страницу
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// PasswordHandler обрабатывает запросы на генерацию пароля
func PasswordHandler(w http.ResponseWriter, r *http.Request) {
	lengthStr := r.URL.Query().Get("length")
	includeSpecialStr := r.URL.Query().Get("special")

	length := 8
	includeSpecial := false

	if lengthStr != "" {
		if l, err := strconv.Atoi(lengthStr); err == nil && l >= 8 && l <= 12 {
			length = l
		}
	}

	if includeSpecialStr == "true" {
		includeSpecial = true
	}

	password := utils.GeneratePassword(length, includeSpecial)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(password))
}
