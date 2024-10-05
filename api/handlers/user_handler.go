package handlers

import (
	"api/services"
	"libs/database"
	"libs/utils"

	"github.com/gofiber/fiber/v2"
)

var user_collection_name = "users"

func CreateUserHandler(ctx *fiber.Ctx) error {
	cl := ctx.Locals("clients").(*utils.Clients)
	apiClient := cl.DBClients["api"].Client
	collection := database.GetCollection(apiClient, user_collection_name)

	userService := services.NewUserService(collection)

	return userService.CreateUser(ctx)
}

func GetUserHandler(ctx *fiber.Ctx) error {
	cl := ctx.Locals("clients").(*utils.Clients)
	apiClient := cl.DBClients["api"].Client
	collection := database.GetCollection(apiClient, user_collection_name)

	userService := services.NewUserService(collection)

	return userService.FindOneUserByID(ctx)
}

func UpdateUserHandler(ctx *fiber.Ctx) error {
	cl := ctx.Locals("clients").(*utils.Clients)
	apiClient := cl.DBClients["api"].Client
	collection := database.GetCollection(apiClient, user_collection_name)

	userService := services.NewUserService(collection)

	return userService.UpdateUser(ctx)
}
