package handlers

import (
	"infra/services"
	"libs/constants"
	"libs/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var city_collection_name = "cities"

func GetCityHandler(ctx *fiber.Ctx) error {
	dbClient := ctx.Locals("dbClient").(*mongo.Client)
	collection := database.GetCollection(dbClient, constants.LOCATIONS_DB_NAME, city_collection_name)

	cityService := services.NewCityService(collection)

	return cityService.FindOneCityByID(ctx)
}
