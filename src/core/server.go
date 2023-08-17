package core

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/csepulvedaz/meli-challenge/src/middlewares"
	"github.com/csepulvedaz/meli-challenge/src/routes"
	"github.com/csepulvedaz/meli-challenge/src/utils"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber/
func fiberConfig() fiber.Config {
	return fiber.Config{
		AppName: fmt.Sprintf("Top Secret Api (%s)", utils.GetEnv()),
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
	port := utils.GetPort()
	if err := app.Listen(port); err != nil {
		log.Fatalf("Oops... Server is not running! Reason: %v", err)
	}
}
