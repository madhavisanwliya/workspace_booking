package config

import (
	"log"

	"github.com/joho/godotenv"
)

// init is a special function in go will be called automatically on the startup
// this will run before main and load env variables into the scope
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
