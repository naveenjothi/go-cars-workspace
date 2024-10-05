package routes

import (
	"infra/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterFiberRoutes(app *fiber.App) {
	app.Get("/city/:id", handlers.GetCityHandler)
	app.Get("/country/:id", handlers.GetCountryHandler)
}
