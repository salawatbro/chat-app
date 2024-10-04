package database

import (
	"context"
	"fmt"
	"github.com/salawatbro/chat-app/internal/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func Connect(cfg *config.Config) (*Database, error) {
	// Set client options using the MongoDB URI from .env
	clientOptions := options.Client().ApplyURI(cfg.DB.MongoDBURI)
	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	db := client.Database(cfg.DB.DatabaseName)
	return &Database{
		Client: client,
		DB:     db,
	}, nil
}

func (d *Database) Disconnect() error {
	if err := d.Client.Disconnect(context.Background()); err != nil {
		return err
	}
	fmt.Println("Disconnected from MongoDB")
	return nil
}
