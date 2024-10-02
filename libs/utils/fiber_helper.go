package utils

import (
	"encoding/json"
	"fmt"
	"libs/database"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBodyPayload(ctx *fiber.Ctx, dto interface{}) error {
	body := ctx.Body()
	if err := json.Unmarshal(body, dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Unable to parse body"})
	}
	return nil
}

type FiberServer struct {
	app    *fiber.App
	client *mongo.Client
}

// NewFiberServer initializes a new Fiber server with a MongoDB client.
func NewFiberServer(prefix, appName string) *FiberServer {
	client, err := database.InitializeMongoClient(prefix)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB client: %s", err)
	}

	// Create a new Fiber server instance
	return &FiberServer{
		app: fiber.New(fiber.Config{
			ServerHeader: appName,
			AppName:      appName,
		}),
		client: client,
	}
}

// WithClient returns a Fiber handler that provides the MongoDB client to the handler.
func (s *FiberServer) WithClient(handler func(*fiber.Ctx, *mongo.Client) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func(start time.Time) {
			fmt.Printf("%s %s took=%v\n", c.Route().Method, c.Path(), time.Since(start))
		}(time.Now())
		return handler(c, s.client)
	}
}

func (s *FiberServer) Listen(addr string) error {
	return s.app.Listen(addr)
}

// Add any additional methods you want to FiberServer here
func (s *FiberServer) Get(path string, handler fiber.Handler) fiber.Router {
	return s.app.Get(path, handler)
}

// Method to register POST routes
func (s *FiberServer) Post(path string, handler fiber.Handler) fiber.Router {
	return s.app.Post(path, handler)
}
