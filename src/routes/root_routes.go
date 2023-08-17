package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func getEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	return env
}

// healthCheck func for api health check
func healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "OK",
		"env":    getEnv(),
	})
}

// rootRoutes func for describe group of routes.
func rootRoutes(a *fiber.App) {
	route := a.Group("/")
	route.Get("/", healthCheck)
}
