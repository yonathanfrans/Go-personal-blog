package storage

import (
	"encoding/json"
	"example/personal-blog/model"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getArticleFilePath(id int) string {
	return fmt.Sprintf("articles/%d.json", id)
}

func SaveArticle(article model.Article) error {
	err := os.MkdirAll("articles", os.ModePerm)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(article, "", " ")
	if err != nil {
		return err
	}

	filePath := getArticleFilePath(article.ID)

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadArticle(id int) (model.Article, error) {
	filePath := getArticleFilePath(id)

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return model.Article{}, err
	}

	var article model.Article
	err = json.Unmarshal(fileData, &article)
	if err != nil {
		return model.Article{}, err
	}

	return article, nil
}

func LoadArticles() ([]model.Article, error) {
	if _, err := os.Stat("articles"); os.IsNotExist(err) {
		return []model.Article{}, nil
	}

	files, err := os.ReadDir("articles")
	if err != nil {
		return nil, err
	}

	var articles []model.Article

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		if !strings.HasSuffix(fileName, ".json") {
			continue
		}
		
		idStr := strings.TrimSuffix(fileName, ".json")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}

		article, err := LoadArticle(id)
		if err != nil {
			log.Printf("Warning: Failed to read ID article %d: %v", id, err)
			continue
		}

		articles = append(articles, article)
	}
	return articles, nil
}

func DeleteArticle(id int) error {
	filePath := getArticleFilePath(id)

	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}