package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type URL struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	LongURL  string             `bson:"long_url"`
	ShortURL string             `bson:"short_url"`
}
