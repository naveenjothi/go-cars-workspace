package utils

import (
	"context"
	"libs/database"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type DBClient struct {
	Name   string
	Client *mongo.Client
}

type Clients struct {
	DBClients map[string]*DBClient
}

func LoadClients() *Clients {
	dbClients := make(map[string]*DBClient)

	apiClient, err := database.InitializeMongoClient("API")
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB client: %s", err)
	}
	locationClient, err := database.InitializeMongoClient("LOCATION")
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB client: %s", err)
	}
	dbClients["api"] = &DBClient{Name: "api", Client: apiClient}
	dbClients["location"] = &DBClient{Name: "location", Client: locationClient}

	return &Clients{
		DBClients: dbClients,
	}
}

func GracefulShutdown(app *fiber.App, clients *Clients) {
	// Create a channel to listen for OS signals (SIGINT, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-quit

	log.Println("Shutting down server...")

	// Shut down Fiber gracefully with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Fiber forced to shutdown: %v", err)
	}

	// Close all database connections
	log.Println("Closing database connections...")
	for _, dbClient := range clients.DBClients {
		if err := dbClient.Client.Disconnect(ctx); err != nil {
			log.Printf("Failed to close DB client %s: %v", dbClient.Name, err)
		}
	}

	log.Println("Server shutdown complete.")
}
