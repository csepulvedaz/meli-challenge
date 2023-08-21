package utils_test

import (
	"reflect"
	"testing"

	"github.com/csepulvedaz/meli-challenge/src/models"
	"github.com/csepulvedaz/meli-challenge/src/utils"
)

func TestGetLocation(t *testing.T) {
	// Basic test with correct distances
	x, y := utils.GetLocation(600, 650, 880)
	expectedX, expectedY := -345.6875, 411.6250
	if x != float32(expectedX) || y != float32(expectedY) {
		t.Errorf("GetLocation() returns (%f, %f), was expected (%f, %f)", x, y, expectedX, expectedY)
	}

	// It will work with incorrect distances, but the result will be wrong
}

func TestGetMessage(t *testing.T) {
	// Basic test
	msg := utils.GetMessage(
		[]string{"este", "", "", "mensaje", ""},
		[]string{"", "es", "", "", "secreto"},
		[]string{"este", "", "un", "", ""},
	)
	expectedMsg := "este es un mensaje secreto"
	if msg != expectedMsg {
		t.Errorf("GetMessage() returns %q, was expected %q", msg, expectedMsg)
	}

	// Test with empty messages
	msg = utils.GetMessage(
		[]string{"", "", "", "", ""},
		[]string{"", "", "", "", ""},
		[]string{"", "", "", "", ""},
	)
	expectedMsg = ""
	if msg != expectedMsg {
		t.Errorf("GetMessage() returns %q, was expected %q", msg, expectedMsg)
	}

	// Test with extra words
	msg = utils.GetMessage(
		[]string{"este", "", "", "mensaje", ""},
		[]string{"", "es", "", "", "secreto", "extra"},
		[]string{"este", "", "un", "", ""},
	)
	expectedMsg = "este es un mensaje secreto extra"
	if msg != expectedMsg {
		t.Errorf("GetMessage() returns %q, was expected %q", msg, expectedMsg)
	}
}

func TestFindSatelliteByName(t *testing.T) {
	// Basic test
	satellites := []models.Satellite{
		{Name: "kenobi", Distance: 100},
		{Name: "skywalker", Distance: 115.5},
		{Name: "sato", Distance: 142.7},
	}

	// Test with found name
	satellite := utils.FindSatelliteByName(satellites, "skywalker")
	expectedSatellite := models.Satellite{Name: "skywalker", Distance: 115.5}
	if !reflect.DeepEqual(satellite, expectedSatellite) {
		t.Errorf("FindSatelliteByName() returns %+v, was expected %+v", satellite, expectedSatellite)
	}

	// Test with not found name
	satellite = utils.FindSatelliteByName(satellites, "notexist")
	expectedSatellite = models.Satellite{}
	if !reflect.DeepEqual(satellite, expectedSatellite) {
		t.Errorf("FindSatelliteByName() returns %+v, was expected %+v", satellite, expectedSatellite)
	}
}
