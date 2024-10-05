package handlers

import (
	"infra/services"
	"libs/constants"
	"libs/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var country_collection_name = "countries"

func GetCountryHandler(ctx *fiber.Ctx) error {
	dbClient := ctx.Locals("dbClient").(*mongo.Client)
	collection := database.GetCollection(dbClient, constants.LOCATIONS_DB_NAME, country_collection_name)

	cityService := services.NewCityService(collection)

	return cityService.FindOneCityByID(ctx)
}
