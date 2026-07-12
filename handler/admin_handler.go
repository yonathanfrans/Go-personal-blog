package handler

import (
	"example/personal-blog/model"
	"example/personal-blog/storage"
	"html/template"
	"net/http"
)

var dashboardTemplate = template.Must(template.ParseFiles("templates/admin/dashboard.html"))

type DashboardPageData struct {
	Title string
	Header string
	SubHeader string
	Articles []model.Article
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	articles, err := storage.LoadArticles()
	if err != nil {
		http.Error(w, "Failed to Load Articles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := DashboardPageData{
		Title: "Dashboard Page",
		Header: "My Dashboard Page",
		SubHeader: "Welcome Admin!",
		Articles: articles,
	}

	err = dashboardTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}