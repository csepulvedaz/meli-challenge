package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv func for loading environment variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
}
