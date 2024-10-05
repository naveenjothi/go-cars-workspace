package routes

import (
	"api/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterFiberRoutes(app *fiber.App) {
	app.Post("/user", handlers.CreateUserHandler)
	app.Get("/user/:id", handlers.GetUserHandler)
	app.Post("/user/:id", handlers.UpdateUserHandler)
	app.Post("/product", handlers.CreateProductHandler)
	app.Get("/product/:id", handlers.GetProductHandler)
	app.Post("/product/:id", handlers.UpdateProductHandler)
}
