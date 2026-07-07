package model

import "time"

type Article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	PublishedAt time.Time `json:"publishedAt"`
}