package config

import (
	"github.com/rmarmolejo90/hvac/internal/config/log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Unable to load environment variables -->> " + err.Error())
	}
	log.Infof("Successfully oaded environment variables!")
}
