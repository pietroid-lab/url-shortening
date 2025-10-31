package internal

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	MongoClient *mongo.Client
}

func NewRepository(mongoCLient *mongo.Client) *Repository {
	return &Repository{MongoClient: mongoCLient}
}

func (r *Repository) SaveURL(ctx context.Context, originalURL, hash string) {
	coll := r.MongoClient.Database("shortening-url").Collection("shortening-url")

	_, err := coll.InsertOne(ctx, bson.M{"originalURL": originalURL, "key": hash})
	if err != nil {
		log.Fatal("error save URL:", err)
	}

}
