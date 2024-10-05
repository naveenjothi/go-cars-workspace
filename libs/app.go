package app

import (
	"fmt"
	"libs/utils"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Listen starts the Fiber server and listens on the specified port.
func Listen(prefix string, beforeAppListen func(app *fiber.App)) {
	appName := os.Getenv(fmt.Sprintf("%s_APP_NAME", prefix))
	appPort := os.Getenv(fmt.Sprintf("%s_APP_PORT", prefix))

	// Create a new Fiber server instance
	clients := utils.LoadClients()
	app := utils.NewFiberServer(prefix, appName)
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("clients", clients)
		return c.Next()
	})

	// Convert port to integer and handle error
	port, err := strconv.Atoi(appPort)
	if err != nil {
		log.Fatalf("Invalid port number: %s", appPort)
	}

	// Call the provided handler function before starting the server
	beforeAppListen(app)

	// Start listening on the specified
	log.Printf("Starting %s on port %s", appName, appPort)
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
			log.Fatalf("Error starting Fiber server: %v", err)
		}
	}()

	// Call graceful shutdown when SIGINT or SIGTERM is received
	utils.GracefulShutdown(app, clients)
}
