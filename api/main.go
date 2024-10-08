package main

import (
	"api/routes"
	"libs/database"
	"log"

	"github.com/gofiber/fiber/v2"

	app "libs"
)

func main() {
	app.Listen("API", func(app *fiber.App) {
		dbClient, err := database.InitializeMongoClient("API")
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
