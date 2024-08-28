package db

import (
	"os"

	"github.com/rmarmolejo90/hvac/config"
	"github.com/rmarmolejo90/hvac/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	config.LoadEnv()
	dsn := os.Getenv("DB_DSN")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Infof("\n\n*** Connected To The Database ***\n\n")

	// Need to add all models for migration
	err = DB.AutoMigrate()
	if err != nil {
		log.Errorf("Error Migrating the database: " + err.Error())
	}
}
