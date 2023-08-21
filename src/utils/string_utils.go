package utils

import (
	"encoding/json"

	"github.com/csepulvedaz/meli-challenge/src/models"
)

// StructToString for converting a struct to string
func SatelliteToString(satellite models.Satellite) (string, error) {
	jsonData, err := json.Marshal(satellite)
	return string(jsonData), err
}

// StringToStruct for converting a string to struct
func StringToSatellite(str string) (models.Satellite, error) {
	var satellite models.Satellite
	err := json.Unmarshal([]byte(str), &satellite)
	return satellite, err
}
