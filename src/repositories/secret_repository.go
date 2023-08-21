package repositories

import (
	"log"

	consts "github.com/csepulvedaz/meli-challenge/src/constants"
	"github.com/csepulvedaz/meli-challenge/src/core/database"
	"github.com/csepulvedaz/meli-challenge/src/models"
	"github.com/csepulvedaz/meli-challenge/src/utils"
)

// GetSatellite for getting a satellite from database
func GetSatellites() ([]models.Satellite, error) {
	client := database.GetRedisClient()
	keys := consts.SatelliteKeys
	satellites := []models.Satellite{}

	// Get the satellites from the database
	for _, key := range keys {

		// Get the satellite from the database
		res, err := client.Get(key)
		if err != nil {
			log.Println(err)
			return []models.Satellite{}, err
		}

		// Check if the value is empty and continue
		if res == "" {
			continue
		}

		// String to satellite struct
		satellite, err := utils.StringToSatellite(res)
		if err != nil {
			log.Fatalln(err)
			return []models.Satellite{}, err
		}

		satellites = append(satellites, satellite)
	}

	return satellites, nil
}

// UpsertSatellite for upserting a satellite
func UpsertSatellite(name string, satellite string) error {
	client := database.GetRedisClient()

	// Set the satellite in the database
	err := client.Set(name, satellite)
	if err != nil {
		return err
	}

	return nil
}
