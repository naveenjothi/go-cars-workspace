package handlers

import (
	"api/services"
	"libs/constants"
	"libs/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var user_collection_name = "users"

func CreateUserHandler(ctx *fiber.Ctx) error {
	dbClient := ctx.Locals("dbClient").(*mongo.Client)
	collection := database.GetCollection(dbClient, constants.API_DB_NAME, user_collection_name)

	userService := services.NewUserService(collection)

	return userService.CreateUser(ctx)
}

func GetUserHandler(ctx *fiber.Ctx) error {
	dbClient := ctx.Locals("dbClient").(*mongo.Client)
	collection := database.GetCollection(dbClient, constants.API_DB_NAME, user_collection_name)

	userService := services.NewUserService(collection)

	return userService.FindOneUserByID(ctx)
}

func UpdateUserHandler(ctx *fiber.Ctx) error {
	dbClient := ctx.Locals("dbClient").(*mongo.Client)
	collection := database.GetCollection(dbClient, constants.API_DB_NAME, user_collection_name)

	userService := services.NewUserService(collection)

	return userService.UpdateUser(ctx)
}
