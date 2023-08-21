package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/csepulvedaz/meli-challenge/src/controllers"
)

// SecretRoutes func for describe group of routes.
func SecretRoutes(a *fiber.App) {
	// Top secret
	a.Post("/topsecret", controllers.GenerateTopSecret)

	// Top secret split
	a.Get("/topsecret_split", controllers.GenerateSplitTopSecret)
	a.Post("/topsecret_split/:satellite_name", controllers.AddSplitSatellite)
	a.Delete("/topsecret_split/:satellite_name", controllers.DeleteSplitSatellite)
}
