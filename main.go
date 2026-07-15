package main

import (
	"example/personal-blog/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/article", handler.ArticleHandler)

	http.HandleFunc("/admin", handler.DashboardHandler)
	http.HandleFunc("/admin/add", handler.AddArticleHandler)
	http.HandleFunc("/admin/edit", handler.EditArticleHandler)
	http.HandleFunc("/admin/delete", handler.DeleteArticleHandler)

	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/logout", handler.LogoutHandler)
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}