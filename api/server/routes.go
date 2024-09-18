package server

import (
	"api/handlers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Post("/user", s.withClient(handlers.CreateUserHandler))
	s.App.Get("/user/:id", s.withClient(handlers.GetUserHandler))
	s.App.Post("/user/:id", s.withClient(handlers.UpdateUserHandler))
}

func (s *FiberServer) withClient(handler func(*fiber.Ctx, *mongo.Client) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler(c, s.client)
	}
}
