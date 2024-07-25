package storage

import (
	"context"
	"math/rand"
	"time"
	"url-shortner/model"

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
	var url model.URL
	err := s.db.Collection("urls").FindOne(ctx, bson.M{"short_url": shortURL}).Decode(&url)
	if err != nil {
		return "", err
	}
	return url.LongURL, nil
}
