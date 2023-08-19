package utils

import (
	"math"
)

func RoundFloat(val float32, precision int) float32 {
	ratio := math.Pow(10, float64(precision))
	return float32(math.Round(float64(val)*ratio) / ratio)
}
