package database

import (
	"context"
	"log"
	"time"

	"skema/app/model"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitialConnect(dbName string, mongoURI string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatalf("Error while connecting to mongo: %v\n", err)
	}

	return client.Database(dbName)
}

func InitConnection() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "admin:admin@/skema?charset=utf8&parseTime=True&loc=Local")
	return db, err
}

func StubConnection() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", ":memory:?cache=shared")
	db.AutoMigrate(
		model.User{},
		model.Schema{},
		model.Comment{},
	)
	return db, err
}
