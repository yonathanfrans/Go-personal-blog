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
	BasePageData
}

func AddArticleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data := AddArticlePageData {
			BasePageData: BasePageData{
				Title: "Add a New Article - Admin",
				Header: "My Dashboard Page",
				SubHeader: "Add a New Article Here!",
			},
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
			http.Error(w, "Title and content cannot be emptyg", http.StatusBadRequest)
			return
		}

		articles, err := storage.LoadArticles()
		if err != nil {
			http.Error(w, "Failed to Load Articles: "+err.Error(), http.StatusInternalServerError)
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
			http.Error(w, "Failed to Save a New Article: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}