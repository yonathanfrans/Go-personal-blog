package handler

import (
	"errors"
	"example/personal-blog/storage"
	"net/http"
	"os"
)

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	article, err := getArticleFromRequest(r)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			http.Error(w, "Article Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = storage.DeleteArticle(article.ID)
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