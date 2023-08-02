package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var PORT string

func LoadAppCfg() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
}
