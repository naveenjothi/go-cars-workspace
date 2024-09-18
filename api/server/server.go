package server

import (
	"fmt"
	"libs/database"

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
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go-cars",
			AppName:      "go-cars",
		}),
		client: client,
	}

	return server
}
