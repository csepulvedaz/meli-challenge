package core

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"

	"github.com/csepulvedaz/meli-challenge/src/middlewares"
	"github.com/csepulvedaz/meli-challenge/src/routes"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber/
func fiberConfig() fiber.Config {
	return fiber.Config{
		AppName: fmt.Sprintf("Top Secret Api (%s)", os.Getenv("ENV")),
	}
}

// StartServer func for starting the server
func StartServer() {
	godotenv.Load()

	// Create new fiber app
	app := fiber.New(fiberConfig())

	// Registry middlewares
	middlewares.FiberMiddleware(app)

	// Registry routes
	routes.RootRoutes(app)
	routes.SecretRoutes(app)

	// Run server
	port := ":" + os.Getenv("PORT")
	if err := app.Listen(port); err != nil {
		log.Fatalf("Oops... Server is not running! Reason: %v", err)
	}
}
