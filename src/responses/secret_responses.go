package responses

import (
	"github.com/csepulvedaz/meli-challenge/src/models"
)

type SecretResponse struct {
	Position models.Position `json:"position"`
	Message  string          `json:"message"`
}
