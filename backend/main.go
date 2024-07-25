package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"url-shortner/handler"
	"url-shortner/storage"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://database:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("urlshortener")
	urlStore := storage.NewMongoStorage(db)

	//http.Handle("/", http.FileServer(http.Dir("/url-shortner/frontend")))
	//http.Handle("/staic/", http.StripPrefix("/static/", http.FileServer(http.Dir("/url-shortner/frontend/static"))))

	// Routes
	http.HandleFunc("/shorten", handler.ShortenURL(urlStore))
	//http.HandleFunc("/", handler.Redirect(urlStore))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "../frontend/index.html")
		} else {
			handler.Redirect(urlStore)(w, r)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8181"
	}

	fmt.Printf("Server started at :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
