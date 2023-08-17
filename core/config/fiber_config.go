package config

import (
	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber/
func FiberConfig() fiber.Config {
	// Return Fiber configuration.
	return fiber.Config{
		AppName: "Top Secret Api",
	}
}
