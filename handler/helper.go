package handler

import (
	"errors"
	"example/personal-blog/model"
	"example/personal-blog/storage"
	"net/http"
	"strconv"
)

type BasePageData struct {
	Title string
	Header string
	SubHeader string
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