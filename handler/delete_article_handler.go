package handler

import (
	"errors"
	"example/personal-blog/storage"
	"net/http"
	"os"
	"strconv"
)

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter, must be a number", http.StatusBadRequest)
		return
	}

	err = storage.DeleteArticle(id)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			http.Error(w, "Article Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete article: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}