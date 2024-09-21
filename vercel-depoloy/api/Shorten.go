package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ARUP-G/URL-Shortener-With-GO/handler"
	"github.com/ARUP-G/URL-Shortener-With-GO/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var urlStore *storage.MongoStorage

func init() {
	// Setup MongoDB connection
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	log.Println("Connected to MongoDB")
	db := client.Database("urlshortener")
	urlStore = storage.NewMongoStorage(db)
}

// ShortenURL is the Vercel serverless function for shortening URLs
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	handler.ShortenURL(urlStore)(w, r)
}

// Vercel requires that the main function name matches the file name without the extension
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	ShortenURL(w, r)
}
