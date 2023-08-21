package utils_test

import (
	"reflect"
	"testing"

	"github.com/csepulvedaz/meli-challenge/src/models"
	"github.com/csepulvedaz/meli-challenge/src/utils"
)

func TestSatelliteToString(t *testing.T) {
	// Create a Satellite object to convert to string
	satellite := models.Satellite{
		Name:     "kenobi",
		Distance: 100.0,
		Message:  []string{"este", "", "", "mensaje", ""},
	}

	// Convert the Satellite to string
	str, err := utils.SatelliteToString(satellite)

	// Verify that there are no errors
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// Verify that the string is the expected one
	expectedStr := `{"name":"kenobi","distance":100,"message":["este","","","mensaje",""]}`
	if str != expectedStr {
		t.Errorf("TestSatelliteToString() returns %+v, was expected %+v", expectedStr, str)
	}
}

func TestStringToSatellite(t *testing.T) {
	// Create a string to convert to Satellite
	str := `{"name":"kenobi","distance":100,"message":["este","","","mensaje",""]}`

	// Convert the string to Satellite
	satellite, err := utils.StringToSatellite(str)

	// Verify that there are no errors
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// Verify that the Satellite is the expected one
	expectedSatellite := models.Satellite{
		Name:     "kenobi",
		Distance: 100.0,
		Message:  []string{"este", "", "", "mensaje", ""},
	}
	if !reflect.DeepEqual(satellite, expectedSatellite) {
		t.Errorf("TestStringToSatellite() returns %+v, was expected %+v", satellite, expectedSatellite)
	}
}
