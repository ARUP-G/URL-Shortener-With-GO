package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/ARUP-G/URL-Shortener-With-GO/storage"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenURL(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ShortenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("Error decoding request: %v", err)
			http.Error(w, "Invalid request payload"+err.Error(), http.StatusBadRequest)
			return
		}

		// Validate the URL
		_, err := url.ParseRequestURI(req.URL)
		if err != nil {
			http.Error(w, "Invalid URL format", http.StatusBadRequest)
			return
		}

		shortURL, err := store.SaveURL(context.Background(), req.URL)
		if err != nil {
			log.Printf("Error saving YRL: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := ShortenResponse{ShortURL: shortURL}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func Redirect(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.URL.Path[1:]
		log.Printf("Received request for short URL: %s", shortURL)

		longURL, err := store.GetURL(context.Background(), shortURL)
		if err != nil {
			log.Printf("Error retrieving long URL for %s: %v", shortURL, err)
			http.NotFound(w, r)
			return
		}

		log.Printf("Retrieved long URL: %s", longURL)

		log.Printf("Redirecting to: %s", longURL)
		http.Redirect(w, r, longURL, http.StatusFound)
	}
}
