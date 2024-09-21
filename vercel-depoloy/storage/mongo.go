package storage

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/ARUP-G/URL-Shortener-With-GO/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type Storage interface {
	SaveURL(ctx context.Context, longURL string) (string, error)
	GetURL(ctx context.Context, shortURL string) (string, error)
}

type MongoStorage struct {
	db *mongo.Database
}

func NewMongoStorage(db *mongo.Database) *MongoStorage {
	return &MongoStorage{db: db}
}

func (s *MongoStorage) SaveURL(ctx context.Context, longURL string) (string, error) {
	shortURL := randString(6)
	url := model.URL{LongURL: longURL, ShortURL: shortURL}
	_, err := s.db.Collection("urls").InsertOne(ctx, url)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func (s *MongoStorage) GetURL(ctx context.Context, shortURL string) (string, error) {
	log.Printf("Attempting to retrieve long URL for short URL: %s", shortURL)

	var url model.URL
	err := s.db.Collection("urls").FindOne(ctx, bson.M{"short_url": shortURL}).Decode(&url)
	if err != nil {
		log.Printf("Error retrieving URL from database: %v", err)
		return "", err
	}

	log.Printf("Retrieved long URL from database: %s", url.LongURL)
	return url.LongURL, nil
}
