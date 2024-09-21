package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeMongoClient(prefix string) (*mongo.Client, error) {
	dbURI := os.Getenv(fmt.Sprintf("%s_APP_URI", prefix))
	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB")
	return client, nil
}

func getCollection(client *mongo.Client, dbName, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

func GetCollection(client *mongo.Client, collName string) *mongo.Collection {
	dbName := os.Getenv("DB_NAME")
	return getCollection(client, dbName, collName)
}
