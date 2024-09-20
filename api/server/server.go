package server

import (
	"fmt"
	"libs/database"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type FiberServer struct {
	*fiber.App

	client *mongo.Client
}

func New() *FiberServer {
	client, err := database.InitializeMongoClient()
	if err != nil {
		fmt.Println(err.Error())
	}
	appName := os.Getenv("API_APP_NAME")
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: appName,
			AppName:      appName,
		}),
		client: client,
	}

	return server
}
