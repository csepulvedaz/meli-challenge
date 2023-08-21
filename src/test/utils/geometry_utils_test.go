package utils_test

import (
	"testing"

	"github.com/csepulvedaz/meli-challenge/src/models"
	"github.com/csepulvedaz/meli-challenge/src/utils"
)

func TestCanTriangulate(t *testing.T) {
	// Basic test
	pos1 := models.Position{X: 0, Y: 0}
	pos2 := models.Position{X: 0, Y: 3}
	pos3 := models.Position{X: 4, Y: 0}

	satellites := []models.Satellite{
		{Name: "kenobi", Distance: 5},
		{Name: "skywalker", Distance: 4},
		{Name: "sato", Distance: 6},
	}

	canTriangulate := utils.CanTriangulate(2, 2, pos1, pos2, pos3, satellites)
	expectedCanTriangulate := true
	if canTriangulate != expectedCanTriangulate {
		t.Errorf("CanTriangulate() return %t, was expected %t", canTriangulate, expectedCanTriangulate)
	}

	// Test with out-of-range position
	canTriangulate = utils.CanTriangulate(10, 10, pos1, pos2, pos3, satellites)
	expectedCanTriangulate = false
	if canTriangulate != expectedCanTriangulate {
		t.Errorf("CanTriangulate() returns %t, was expected %t", canTriangulate, expectedCanTriangulate)
	}

	// Test with exact distance
	pos1 = models.Position{X: 0, Y: 0}
	pos2 = models.Position{X: 0, Y: 3}
	pos3 = models.Position{X: 4, Y: 0}

	satellites = []models.Satellite{
		{Name: "kenobi", Distance: 5},
		{Name: "skywalker", Distance: 4},
		{Name: "sato", Distance: 5},
	}

	canTriangulate = utils.CanTriangulate(2, 2, pos1, pos2, pos3, satellites)
	expectedCanTriangulate = true
	if canTriangulate != expectedCanTriangulate {
		t.Errorf("CanTriangulate() returns %t, was expected %t", canTriangulate, expectedCanTriangulate)
	}
}
