package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/csepulvedaz/meli-challenge/core/config"
	"github.com/csepulvedaz/meli-challenge/core/middleware"
)

func main() {
	// Create new fiber app
	app := fiber.New(config.FiberConfig())

	// Middlewares
	middleware.FiberMiddleware(app)

	// Routes
	config.SetupRoutes(app)

	// Start server
	config.StartServer(app)
}
