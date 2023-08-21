package utils

import (
	"math"

	"github.com/csepulvedaz/meli-challenge/src/models"
)

// distance for getting the distance between two points
func distance(x1, y1, x2, y2 float32) float32 {
	return float32(math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))))
}

// CanTriangulate for checking if the emitter can be triangulated
func CanTriangulate(x, y float32, pos1, pos2, pos3 models.Position, satellites []models.Satellite) bool {
	// Get the distance from the calculated emitter to the satellites
	d1 := distance(x, y, pos1.X, pos1.Y)
	d2 := distance(x, y, pos2.X, pos2.Y)
	d3 := distance(x, y, pos3.X, pos3.Y)

	// Get the distance from the satellites to the emitter
	r1 := satellites[0].Distance
	r2 := satellites[1].Distance
	r3 := satellites[2].Distance

	/* The formula fails if the distance is greater than the radius because
	the Positions do not match the distances given in order to find the
	position of the emitter.
	*/
	if RoundFloat(d1, 1) > RoundFloat(r1, 1) || RoundFloat(d2, 1) > RoundFloat(r2, 1) || RoundFloat(d3, 1) > RoundFloat(r3, 1) {
		return false
	}

	return true

}
