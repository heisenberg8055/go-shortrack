package config

import (
	"github.com/joho/godotenv"
)

func LoadConfig() {
	_ = godotenv.Load()
}
