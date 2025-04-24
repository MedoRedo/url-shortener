package main

import (
	"fmt"
	"net/http"
	"url-shortener/handler"
)

func main() {
	fmt.Println("Application started ...")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", handler.ShortenURL)
	mux.HandleFunc("/", handler.Redirect)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
