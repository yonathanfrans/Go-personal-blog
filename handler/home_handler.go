package handler

import (
	"example/personal-blog/model"
	"example/personal-blog/storage"
	"html/template"
	"net/http"
)

var homeTemplate = template.Must(template.ParseFiles("templates/home.html"))

type HomePageData struct {
	Title string
	Articles []model.Article
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	articles, err := storage.LoadArticles()
	if err != nil {
		http.Error(w, "Gagal memuat artikel: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := HomePageData{
		Title: "My Personal Blog",
		Articles: articles,
	}
	
	err = homeTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}