package main

import (
	"math"
	"strings"
)

type Position struct {
	X float32
	Y float32
}

var kenobi = Position{-500, -200}
var skywalker = Position{100, -100}
var sato = Position{500, 100}

// GetLocation for getting the location of the emitter
func GetLocation(distances ...float32) (x, y float32) {
	x1, y1 := kenobi.X, kenobi.Y
	x2, y2 := skywalker.X, skywalker.Y
	x3, y3 := sato.X, sato.Y
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

	d1 := distance(x, y, x1, y1)
	d2 := distance(x, y, x2, y2)
	d3 := distance(x, y, x3, y3)

	// Return 0, 0 if the intersection is not valid
	if d1 > r1 || d2 > r2 || d3 > r3 {
		return 0, 0
	}

	return x, y
}

func distance(x1, y1, x2, y2 float32) float32 {
	return float32(math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))))
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
