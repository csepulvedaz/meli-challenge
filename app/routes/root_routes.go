package routes

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// healthCheck func for api health check
func healthCheck(c *fiber.Ctx) error {
	env := os.Getenv("ENV")
	log.Printf("ENV: %v", env)
	return c.JSON(fiber.Map{
		"status": "OK",
		"env":    env,
	})
}

// RootRoutes func for describe group of routes.
func RootRoutes(a *fiber.App) {
	route := a.Group("/")
	route.Get("/", healthCheck)
}
