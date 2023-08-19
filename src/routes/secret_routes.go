package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/csepulvedaz/meli-challenge/src/controllers"
)

// SecretRoutes func for describe group of routes.
func SecretRoutes(a *fiber.App) {
	a.Post("/topsecret", controllers.PostTopSecret)
}
