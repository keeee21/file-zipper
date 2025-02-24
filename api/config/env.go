package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnv returns the value of the environment variable key
func GetEnv(key string) string {
	return os.Getenv(key)
}
