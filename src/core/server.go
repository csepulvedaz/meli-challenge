package core

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/csepulvedaz/meli-challenge/src/middlewares"
	"github.com/csepulvedaz/meli-challenge/src/routes"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber/
func fiberConfig() fiber.Config {
	// Return Fiber configuration.
	return fiber.Config{
		AppName: "Top Secret Api",
	}
}

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

// StartServer func for starting the server
func StartServer() {
	// Create new fiber app
	app := fiber.New(fiberConfig())

	// Registry middlewares
	middlewares.FiberMiddleware(app)

	// Registry routes
	routes.SetupRoutes(app)

	// Run server
	if err := app.Listen(getPort()); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
