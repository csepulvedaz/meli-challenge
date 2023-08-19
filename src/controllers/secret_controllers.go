package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/csepulvedaz/meli-challenge/src/models"
	"github.com/csepulvedaz/meli-challenge/src/services"
	"github.com/csepulvedaz/meli-challenge/src/utils"
)

func PostTopSecret(c *fiber.Ctx) error {
	var data struct {
		Satellites []models.Satellite `json:"satellites"`
	}

	// Parse the body into the data struct
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Check if the satellites are 3
	if len(data.Satellites) != 3 {
		return c.Status(fiber.StatusNotFound).JSON(utils.FormatSecretError("Se necesitan 3 satélites"))
	}

	// Get the secret and check if there is an error
	secret, error := services.GetSecret(data.Satellites)
	if error != nil {
		return c.Status(fiber.StatusNotFound).JSON(error)
	}

	return c.Status(http.StatusOK).JSON(secret)
}
