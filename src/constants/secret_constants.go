package constants

import (
	"github.com/csepulvedaz/meli-challenge/src/models"
)

var KenobiPos = models.Position{X: -500, Y: -200}
var SkywalkerPos = models.Position{X: 100, Y: -100}
var SatoPos = models.Position{X: 500, Y: 100}

var SatelliteKeys = [...]string{"kenobi", "skywalker", "sato"}

const InsuficientInfo = "no hay suficiente información"
const EmitterCannotBeLocated = "no se puede calcular la posición del emisor"
