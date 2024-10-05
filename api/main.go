package main

import (
	"api/routes"

	"github.com/gofiber/fiber/v2"

	app "libs"
)

func main() {
	app.Listen("API", func(app *fiber.App) {
		routes.RegisterFiberRoutes(app)
	})
}
