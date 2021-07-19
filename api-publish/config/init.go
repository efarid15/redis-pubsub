package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadFile(file string) {
	err := godotenv.Load(file)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
