package main

import (
	"fmt"
	"net/http"
	"url-shortener/handler"
	"url-shortener/storage"
)

func main() {
	fmt.Println("Application started ...")
	// Initialize Redis connection
	storage.InitRedis()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", handler.ShortenURL)
	mux.HandleFunc("/", handler.Redirect)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
