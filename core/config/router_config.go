package config

import (
	"github.com/csepulvedaz/meli-challenge/app/routes"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func for registry all api routes.
func SetupRoutes(a *fiber.App) {
	routes.RootRoutes(a)
}
