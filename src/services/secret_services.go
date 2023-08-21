package services

import (
	"errors"

	consts "github.com/csepulvedaz/meli-challenge/src/constants"
	"github.com/csepulvedaz/meli-challenge/src/models"
	"github.com/csepulvedaz/meli-challenge/src/repositories"
	"github.com/csepulvedaz/meli-challenge/src/utils"
)

// GetSecret for getting the secret from the satellites
func GetSecret(satellites []models.Satellite) (models.Secret, error) {
	kenobiPos, skywalkerPos, satoPos := consts.KenobiPos, consts.SkywalkerPos, consts.SatoPos

	kenobi := utils.FindSatelliteByName(satellites, "kenobi")
	skywalker := utils.FindSatelliteByName(satellites, "skywalker")
	sato := utils.FindSatelliteByName(satellites, "sato")

	// Get the message and the current position of the emitter
	message := utils.GetMessage(kenobi.Message, skywalker.Message, sato.Message)
	x, y := utils.GetLocation(kenobi.Distance, skywalker.Distance, sato.Distance)

	// Calculate if the intersection is valid
	canTriangulate := utils.CanTriangulate(x, y, kenobiPos, skywalkerPos, satoPos, satellites)
	if !canTriangulate {
		return models.Secret{}, errors.New(consts.EmitterCannotBeLocated)
	}

	// Round the position and format the secret
	roundedX := utils.RoundFloat(x, 1)
	roundedY := utils.RoundFloat(y, 1)
	secret := utils.FormatSecret(roundedX, roundedY, message)

	return secret, nil
}

// GetSecretSplit for getting the secret from the satellites
func GetSecretSplit() (models.Secret, error) {
	// Get the satellites from the database
	satellites, err := repositories.GetSatellites()
	if err != nil {
		return models.Secret{}, err
	}

	// Check if there are 3 satellites
	if len(satellites) != 3 {
		return models.Secret{}, errors.New(consts.InsuficientInfo)
	}

	// Get the secret and check if there is an error
	secret, err := GetSecret(satellites)
	if err != nil {
		return models.Secret{}, err
	}

	return secret, nil
}

// UpsertSatellite for upserting a satellite
func UpsertSatellite(satellite models.Satellite) error {
	// Stringify the satellite
	stringSatellite, err := utils.SatelliteToString(satellite)
	if err != nil {
		return err
	}

	// Upsert the satellite
	err = repositories.UpsertSatellite(satellite.Name, stringSatellite)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSatellite for deleting a satellite
func DeleteSatellite(name string) error {
	// Delete the satellite
	err := repositories.DeleteSatellite(name)
	if err != nil {
		return err
	}
	return nil
}
