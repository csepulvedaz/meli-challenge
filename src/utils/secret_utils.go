package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	consts "github.com/csepulvedaz/meli-challenge/src/constants"
	"github.com/csepulvedaz/meli-challenge/src/models"
)

// GetLocation for getting the location of the emitter
func GetLocation(distances ...float32) (x, y float32) {
	kenobiPos, skywalkerPos, satoPos := consts.KenobiPos, consts.SkywalkerPos, consts.SatoPos

	// Get the current position of the satellites, and the current distance to the emitter
	x1, y1 := kenobiPos.X, kenobiPos.Y
	x2, y2 := skywalkerPos.X, skywalkerPos.Y
	x3, y3 := satoPos.X, satoPos.Y
	r1, r2, r3 := distances[0], distances[1], distances[2]

	// Calculate intersection of 3 circles
	A := 2 * (x2 - x1)
	B := 2 * (y2 - y1)
	C := r1*r1 - r2*r2 - x1*x1 + x2*x2 - y1*y1 + y2*y2
	D := 2 * (x3 - x2)
	E := 2 * (y3 - y2)
	F := r2*r2 - r3*r3 - x2*x2 + x3*x3 - y2*y2 + y3*y3

	x = (C*E - F*B) / (E*A - B*D)
	y = (C*D - A*F) / (B*D - A*E)

	return x, y
}

// GetMessage for getting the message from the satellites
func GetMessage(messages ...[]string) (msg string) {
	merged := make([]string, 0)
	first, second, third := messages[0], messages[1], messages[2]

	// Merge the messages in order and does not enter two of the same words in a row
	for i := 0; i < len(first) || i < len(second) || i < len(third); i++ {
		if i < len(first) && first[i] != "" && (len(merged) == 0 || first[i] != merged[len(merged)-1]) {
			merged = append(merged, first[i])
		}
		if i < len(second) && second[i] != "" && (len(merged) == 0 || second[i] != merged[len(merged)-1]) {
			merged = append(merged, second[i])
		}
		if i < len(third) && third[i] != "" && (len(merged) == 0 || third[i] != merged[len(merged)-1]) {
			merged = append(merged, third[i])
		}
	}

	return strings.Join(merged, " ")
}

func FormatSecret(x, y float32, message string) models.Secret {
	return models.Secret{
		Position: models.Position{
			X: x,
			Y: y,
		},
		Message: message,
	}
}

func FormatSecretError(msg string) error {
	return fiber.NewError(fiber.StatusNotFound, msg)
}
