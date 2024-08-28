package db
import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=EST"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})