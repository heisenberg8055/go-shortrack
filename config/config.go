package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() map[string]string {
	config, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}
	return config
}
