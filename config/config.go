package config

import (
	"github.com/joho/godotenv"
)

func LoadConfig() map[string]string {
	config, _ := godotenv.Read(".env")
	return config
}
