package services

import (
	consts "github.com/csepulvedaz/meli-challenge/src/constants"
	"github.com/csepulvedaz/meli-challenge/src/models"
	"github.com/csepulvedaz/meli-challenge/src/utils"
)

// GetSecret for getting the secret from the satellites
func GetSecret(satellites []models.Satellite) (models.Secret, error) {
	kenobiPos, skywalkerPos, satoPos := consts.KenobiPos, consts.SkywalkerPos, consts.SatoPos
	kenobi, skywalker, sato := satellites[0], satellites[1], satellites[2]

	// Get the message and the current position of the emitter
	message := utils.GetMessage(kenobi.Message, skywalker.Message, sato.Message)
	x, y := utils.GetLocation(kenobi.Distance, skywalker.Distance, sato.Distance)

	// Calculate if the intersection is valid
	canTriangulate := utils.CanTriangulate(x, y, kenobiPos, skywalkerPos, satoPos, satellites)
	if !canTriangulate {
		return models.Secret{}, utils.FormatSecretError("No se puede calcular la posición del emisor")
	}

	// Round the position and format the secret
	roundedX := utils.RoundFloat(x, 1)
	roundedY := utils.RoundFloat(y, 1)
	secret := utils.FormatSecret(roundedX, roundedY, message)

	return secret, nil
}
