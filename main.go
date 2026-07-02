package main

import (
	"example/personal-blog/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}