package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/csepulvedaz/meli-challenge/src/controllers"
)

// RootRoutes func for describe group of routes.
func RootRoutes(a *fiber.App) {
	route := a.Group("/")
	route.Get("/", controllers.HealthCheck)
}
