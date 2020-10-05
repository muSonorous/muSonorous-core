package main

import (
	"fmt"

	"github.com/muSonorous-core/core/controllers"
)

func main() {
	fmt.Println("Initializing core service.....")
	controller := controllers.New()

	// Hardcoding port... TODO: Remove hardcoded and use env file
	controller.Start("3000")
}
