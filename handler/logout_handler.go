package handler

import "net/http"

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "session",
		Value: "",
		Path: "/",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}