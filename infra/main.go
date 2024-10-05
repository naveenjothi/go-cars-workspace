package main

import (
	"infra/routes"
	"libs/database"
	"log"

	"github.com/gofiber/fiber/v2"

	app "libs"
)

var DB_NAME = "locations"

func main() {
	app.Listen("INFRA", func(app *fiber.App) {
		dbClient, err := database.InitializeMongoClient("INFRA")
		if err != nil {
			log.Fatalf("Failed to initialize MongoDB client: %s", err)
		}
		app.Use(func(c *fiber.Ctx) error {
			c.Locals("dbClient", dbClient)
			return c.Next()
		})
		routes.RegisterFiberRoutes(app)
	})
}
