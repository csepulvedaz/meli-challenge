package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/csepulvedaz/meli-challenge/src/utils"
)

// healthCheck func for api health check
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "OK",
		"env":    utils.GetEnv(),
	})
}
