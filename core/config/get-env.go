package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	DEV   = "dev.env.yaml"
	STAGE = "stage.env.yaml"
	TEST  = "test.env.yaml"
)

// GetEnvVal pass cone of the constants in the method
func GetEnvVal(ENV string, key string) string {

	// load .env file
	// TODO use environment info from Orchestrator to select env file
	err := godotenv.Load(ENV)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
