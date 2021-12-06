package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func SetIndexes(collection *mongo.Collection, keys bsonx.Doc) {
	index := mongo.IndexModel{}
	index.Keys = keys
	unique := true
	index.Options = &options.IndexOptions{
		Unique: &unique,
	}
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)

	_, err := collection.Indexes().CreateOne(context.Background(), index, opts)

	if err != nil {
		log.Fatalf("Error while creating indexes: %v", err)
	}
}
