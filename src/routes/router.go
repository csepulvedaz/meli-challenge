package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func for registry all api routes.
func SetupRoutes(a *fiber.App) {
	rootRoutes(a)
}
