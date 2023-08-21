package utils_test

import (
	"testing"

	"github.com/csepulvedaz/meli-challenge/src/utils"
)

func TestRoundFloat(t *testing.T) {
	// Basic test
	val := utils.RoundFloat(3.14159, 2)
	expectedVal := float32(3.14)
	if val != expectedVal {
		t.Errorf("RoundFloat() returns %f, was expected %f", val, expectedVal)
	}

	// Test with precision 0
	val = utils.RoundFloat(3.14159, 0)
	expectedVal = float32(3)
	if val != expectedVal {
		t.Errorf("RoundFloat() returns %f, was expected %f", val, expectedVal)
	}

	// Test with negative precision
	val = utils.RoundFloat(3.14159, -1)
	expectedVal = float32(3.14159)
	if val != expectedVal {
		t.Errorf("RoundFloat() returns %f, was expected %f", val, expectedVal)
	}

	// Test with negative number
	val = utils.RoundFloat(-3.14159, 2)
	expectedVal = float32(-3.14)
	if val != expectedVal {
		t.Errorf("RoundFloat() returns %f, was expected %f", val, expectedVal)
	}
}
