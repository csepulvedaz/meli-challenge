package core

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/csepulvedaz/meli-challenge/src/middlewares"
	"github.com/csepulvedaz/meli-challenge/src/routes"
	"github.com/csepulvedaz/meli-challenge/src/utils"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber/
func fiberConfig() fiber.Config {
	return fiber.Config{
		AppName: "Top Secret Api",
	}
}

// StartServer func for starting the server
func StartServer() {
	// Create new fiber app
	app := fiber.New(fiberConfig())

	// Registry middlewares
	middlewares.FiberMiddleware(app)

	// Registry routes
	routes.RootRoutes(app)

	// Run server
	if err := app.Listen(utils.GetPort()); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
