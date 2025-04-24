package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"url-shortener/storage"
	"url-shortener/utils"
)

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	ShortURL string `json:"short_url"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	code := utils.GenerateShortCode()
	storage.SaveURL(code, req.URL)

	res := Response{ShortURL: "http://" + r.Host + "/" + code}
	json.NewEncoder(w).Encode(res)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	original, found := storage.GetURL(code)
	if !found {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, original, http.StatusFound)
}
