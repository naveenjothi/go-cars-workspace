package migrations

import (
	"libs/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var cars_collection_name = "cars"

func MigrateCars(client *mongo.Client) error {
	collection := database.GetCollection(client, cars_collection_name)

	// userService := services.NewUserService(collection)

	// return userService.CreateUser(ctx)
	return nil
}
