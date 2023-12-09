package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Err loading .env file")
	}
	return os.Getenv(key)
}
