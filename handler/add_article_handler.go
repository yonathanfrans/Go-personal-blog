package handler

import (
	"example/personal-blog/model"
	"example/personal-blog/storage"
	"html/template"
	"net/http"
	"strings"
	"time"
)

var addArticleTemplate = template.Must(template.ParseFiles("templates/admin/add_article.html"))

type AddArticlePageData struct {
	Title string 
	Header string
	SubHeader string
}

func AddArticleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data := AddArticlePageData {
			Title: "Add a New Article - Admin",
			Header: "My Dashboard Page",
			SubHeader: "Add a New Article Here!",
		}

		err := addArticleTemplate.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return

	case http.MethodPost:
		title := r.FormValue("title")
		content := r.FormValue("content")

		if strings.TrimSpace(title) == "" || strings.TrimSpace(content) == "" {
			http.Error(w, "Title dan Content tidak boleh kosong", http.StatusBadRequest)
			return
		}

		articles, err := storage.LoadArticles()
		if err != nil {
			http.Error(w, "Gagal memuat artikel: "+err.Error(), http.StatusInternalServerError)
			return
		}

		maxID := 0
		for _, article := range articles {
			if article.ID > maxID {
				maxID = article.ID
			}
		}
		nextID := maxID + 1

		newArticle := model.Article {
			ID: nextID,
			Title: title,
			Content: content,
			PublishedAt: time.Now(),
		}

		err = storage.SaveArticle(newArticle)
		if err != nil {
			http.Error(w, "Gagal menyimpan artikel baru: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}