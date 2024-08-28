package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Unable to load environment variables -->> " + err.Error())
	}
}
