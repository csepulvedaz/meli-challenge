package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/csepulvedaz/meli-challenge/src/models"
	"github.com/csepulvedaz/meli-challenge/src/responses"
	"github.com/csepulvedaz/meli-challenge/src/utils"
)

func PostTopSecret(c *fiber.Ctx) error {
	var data struct {
		Satellites []models.Satellite `json:"satellites"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Se necesitan 3 satélites",
		})
	}

	// Validar que se hayan recibido exactamente tres satélites
	if len(data.Satellites) != 3 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Se necesitan 3 satélites",
		})
	}

	satellites := data.Satellites
	kenobi, skywalker, sato := satellites[0], satellites[1], satellites[2]

	message := utils.GetMessage(kenobi.Message, skywalker.Message, sato.Message)
	x, y := utils.GetLocation(kenobi.Distance, skywalker.Distance, sato.Distance)

	response := responses.SecretResponse{
		Position: models.Position{
			X: utils.RoundFloat(x, 1),
			Y: utils.RoundFloat(y, 1),
		},
		Message: message,
	}

	return c.Status(http.StatusOK).JSON(response)
}
