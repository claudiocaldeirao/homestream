package database

import (
	"context"
	"log"
	"time"

	"github.com/claudiocaldeirao/homestream/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect(cfg *config.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(cfg.MongoURI)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ MongoDB ping failed: %v", err)
	}

	Client = client
	log.Println("✅ Connected to MongoDB")
}

func GetDatabase(cfg *config.Config) *mongo.Database {
	if Client == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return Client.Database(cfg.MongoDatabase)
}

func DropDatabase(cfg *config.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := Client.Database(cfg.MongoDatabase).Drop(ctx)

	if err != nil {
		log.Fatalf("Failed to drop database: %v", err)
	}
}
