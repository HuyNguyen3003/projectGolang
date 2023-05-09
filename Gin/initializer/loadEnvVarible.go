package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVarible() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
