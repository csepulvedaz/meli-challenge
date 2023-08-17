package config

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// getPort func for getting the port from the environment
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":4000"
	} else {
		port = ":" + port
	}
	return port
}

// Start func for starting a simple server
func StartServer(a *fiber.App) {
	// Run server
	if err := a.Listen(getPort()); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
