package handler

import (
	"html/template"
	"net/http"
)

var loginTemplate = template.Must(template.ParseFiles("templates/admin/login.html"))

type LoginPageData struct {
	BasePageData
}

const (
	adminUsername = "admin"
	adminPassword = "admin123"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if IsAuthenticated(r) {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		data := LoginPageData{
			BasePageData: BasePageData{
				Title: "Login Page",
				Header: "Welcome to Login Page!",
				SubHeader: "Please fill out the login form",
			},
		}

		err := loginTemplate.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return

	case http.MethodPost:
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == adminUsername && password == adminPassword {
			http.SetCookie(w, &http.Cookie{
				Name: "session",
				Value: "authenticated",
				Path: "/",
				HttpOnly: true,
			})

			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		http.Error(w, "Invalid username/password", http.StatusUnauthorized)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}