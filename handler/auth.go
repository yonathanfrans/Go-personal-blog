package handler

import "net/http"

func IsAuthenticated(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err != nil {
		return false
	}

	return cookie.Value == "authenticated"
}