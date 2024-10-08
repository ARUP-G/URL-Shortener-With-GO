package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ARUP-G/URL-Shortener-With-GO/handler"
	"github.com/ARUP-G/URL-Shortener-With-GO/storage"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/handlers"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://database:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to database")

	db := client.Database("urlshortener")
	urlStore := storage.NewMongoStorage(db)

	// http.Handle("/", http.FileServer(http.Dir("../frontend")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/static"))))

	// Routes
	http.HandleFunc("/shorten", handler.ShortenURL(urlStore))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "../frontend/static/index.html")
		} else {
			handler.Redirect(urlStore)(w, r)
		}
	})
	// Setup CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(http.DefaultServeMux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8181"
	}

	fmt.Printf("Server started at :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
