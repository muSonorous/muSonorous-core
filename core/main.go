package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muSonorous-core/core/controllers"
)

// Get Env
func getEnvValue(key string) string {

	// load .env file
	// TODO use environment info from Orchestrator to select env file
	err := godotenv.Load("dev.env.yaml")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	fmt.Println("Initializing core service.....")
	controller := controllers.New()

	// Hardcoding port... TODO: Remove hardcoded and use env file
	controller.Start(getEnvValue("PORT"))
}
