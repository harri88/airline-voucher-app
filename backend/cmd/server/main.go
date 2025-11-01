package main

import (
	"log"
	"os"
)

// @title Airline Voucher API
// @version 1.0
// @description API for generating airline voucher seats.
// @host localhost:8080
// @BasePath /
func main() {

	app, err := InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server starting on port %s...", port)
	log.Printf("Swagger docs available at http://localhost:%s/swagger/index.html", port)

	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
