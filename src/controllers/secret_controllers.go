package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	consts "github.com/csepulvedaz/meli-challenge/src/constants"
	"github.com/csepulvedaz/meli-challenge/src/models"
	"github.com/csepulvedaz/meli-challenge/src/services"
	"github.com/csepulvedaz/meli-challenge/src/utils"
)

// PostTopSecret for getting the secret from the satellites
func PostTopSecret(c *fiber.Ctx) error {
	var data struct {
		Satellites []models.Satellite `json:"satellites"`
	}

	// Parse the body into the data struct
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.FormatError(err.Error()))
	}

	// Check if the satellites are 3
	if len(data.Satellites) != 3 {
		return c.Status(fiber.StatusNotFound).JSON(utils.FormatError(consts.InsuficientInfo))
	}

	// Get the secret and check if there is an error
	secret, err := services.GetSecret(data.Satellites)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.FormatError(err.Error()))
	}

	return c.Status(http.StatusOK).JSON(secret)
}

// GetTopSecretSplit for getting the secret from the split satellites
func GetTopSecretSplit(c *fiber.Ctx) error {
	// Get the secret and check if there is an error
	secret, err := services.GetSecretSplit()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.FormatError(err.Error()))
	}

	return c.Status(http.StatusOK).JSON(secret)
}

// PostTopSecretSplit for save split satellites
func PostTopSecretSplit(c *fiber.Ctx) error {
	var data models.Satellite

	// Get the value of the "satellite_name" parameter
	satellite_name := c.Params("satellite_name")

	// Parse the body into the data struct
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.FormatError(err.Error()))
	}

	// Upsert satellite
	data.Name = satellite_name
	err := services.UpsertSatellite(data)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.FormatError(err.Error()))
	}

	return c.Status(http.StatusOK).JSON(data)
}
