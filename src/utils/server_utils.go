package utils

import (
	"os"
)

// GetEnv func for getting the environment label
func GetEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	return env
}

// getPort func for getting the port from the environment
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":4000"
	} else {
		port = ":" + port
	}
	return port
}
