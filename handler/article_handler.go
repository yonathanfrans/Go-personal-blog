package handler

import (
	"errors"
	"example/personal-blog/model"
	"html/template"
	"net/http"
	"os"
)

var articleTemplate = template.Must(template.ParseFiles("templates/article.html"))

type ArticlePageData struct {
	BasePageData
	Article model.Article
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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

	// tmpl, err := template.ParseFiles("templates/article.html")
	// if err != nil {
	// 	http.Error(w, "Template tidak ditemukan: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	data := ArticlePageData {
		BasePageData: BasePageData{
			Header: "My Personal Blog",
			SubHeader: "Welcome to My Blog!",
		},
		Article: article,
	}

	err = articleTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}