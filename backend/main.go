package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ARUP-G/URL-Shortener-With-GO/handler"
	"github.com/ARUP-G/URL-Shortener-With-GO/storage"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/handlers"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	// // Get MongoDB URI from environment variable 
	// MONGO_URI := os.Getenv("MONGO_URI")
	// if MONGO_URI == "" {
	// 	log.Fatal("MONGO_URI not set in environment")
	// }

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI('mongodb+srv://ard:Mgo66@app-data-1.1chgr.mongodb.net/?retryWrites=true&w=majority&appName=App-data-1')
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	log.Printf("Connected to database.!")

	db := client.Database("urlshortener")
	urlStore := storage.NewMongoStorage(db)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/static"))))

	// Routes
	http.HandleFunc("/shorten", handler.ShortenURL(urlStore))

	// Root route handling
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "../frontend/static/index.html")
		} else {
			handler.Redirect(urlStore)(w, r)
		}
	})

	// Setup CORS for frontend URL
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"https://url-shortener-with-go-l2go.vercel.app"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(http.DefaultServeMux)

	// Get port from environment or default to 8181
	port := os.Getenv("PORT")
	if port == "" {
		port = "8181"
	}

	// Start the server
	fmt.Printf("Server started at :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
