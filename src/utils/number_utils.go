package utils

import (
	"math"
)

// RoundFloat for rounding a float with precision (decimal places)
func RoundFloat(val float32, precision int) float32 {
	if precision < 0 {
		return val
	}
	ratio := math.Pow(10, float64(precision))
	return float32(math.Round(float64(val)*ratio) / ratio)
}
