package handler

import (
	"errors"
	"example/personal-blog/model"
	"example/personal-blog/storage"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var editArticleTemplate = template.Must(template.ParseFiles("templates/admin/edit_article.html"))

type EditArticlePageData struct {
	Title string
	Header string
	SubHeader string
	Article model.Article
}

func getArticleFromRequest(r *http.Request) (model.Article, error) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		return model.Article{}, errors.New("Missing 'id' parameter")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return model.Article{}, errors.New("Invalid 'id' parameter, must be a number")
	}

	return storage.LoadArticle(id)
}

func EditArticleHandler(w http.ResponseWriter, r *http.Request) {
	article, err := getArticleFromRequest(r)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			http.Error(w, "Article Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	switch r.Method {
	case http.MethodGet:
		data := EditArticlePageData {
			Title: "Edit a Article - Admin",
			Header: "My Dashboard Page",
			SubHeader: "Edit Your Article Here",
			Article: article,
		}

		err = editArticleTemplate.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		title := r.FormValue("title")
		content := r.FormValue("content")

		if strings.TrimSpace(title) == "" || strings.TrimSpace(content) == "" {
			http.Error(w, "Title and content cannot be empty", http.StatusBadRequest)
			return
		}

		article.Title = title
		article.Content = content

		err = storage.SaveArticle(article)
		if err != nil {
			http.Error(w, "Gagal menyimpan artikel: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}