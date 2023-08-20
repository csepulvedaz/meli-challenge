package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

// GetHealthCheck func for api health check
func GetHealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"env":     os.Getenv("ENV"),
		"message": "Pong!",
	})
}
