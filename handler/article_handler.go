package handler

import (
	"errors"
	"example/personal-blog/model"
	"example/personal-blog/storage"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

var articleTemplate = template.Must(template.ParseFiles("templates/article.html"))

type ArticlePageData struct {
	Header string
	SubHeader string
	Article model.Article
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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

	article, err := storage.LoadArticle(id)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			http.Error(w, "Article Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// tmpl, err := template.ParseFiles("templates/article.html")
	// if err != nil {
	// 	http.Error(w, "Template tidak ditemukan: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	data := ArticlePageData {
		Header: "My Personal Blog",
		SubHeader: "Welcome to My Blog!",
		Article: article,
	}

	err = articleTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}