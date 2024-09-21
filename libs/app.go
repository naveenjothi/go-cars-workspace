package app

import (
	"fmt"
	"libs/utils"
	"log"
	"os"
	"strconv"
)

// Listen starts the Fiber server and listens on the specified port.
func Listen(prefix string, beforeAppListen func(server *utils.FiberServer)) {
	appName := os.Getenv(fmt.Sprintf("%s_APP_NAME", prefix))
	appPort := os.Getenv(fmt.Sprintf("%s_APP_PORT", prefix))

	// Create a new Fiber server instance
	server := utils.NewFiberServer(prefix, appName)

	// Convert port to integer and handle error
	port, err := strconv.Atoi(appPort)
	if err != nil {
		log.Fatalf("Invalid port number: %s", appPort)
	}

	// Call the provided handler function before starting the server
	beforeAppListen(server)

	// Start listening on the specified port
	if err := server.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Unable to start the %s: %s", appName, err)
	}
	log.Printf("%s app is running on port %s", appName, appPort)
}
