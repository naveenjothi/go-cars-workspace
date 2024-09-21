package main

import (
	"api/routes"
	app "libs"
	"libs/utils"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app.Listen("API", func(server *utils.FiberServer) {
		routes.RegisterFiberRoutes(server)
	})
}
