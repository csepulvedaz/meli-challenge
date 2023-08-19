package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/csepulvedaz/meli-challenge/src/utils"
)

// GetHealthCheck func for api health check
func GetHealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"env":     utils.GetEnv(),
		"message": "Pong!",
	})
}
