package domain

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	CustomerID  uint
	Address     string `gorm:"type:varchar(255);not null"`
	City        string `gorm:"type:varchar(50);not null"`
	State       string `gorm:"type:varchar(50);not null"`
	ZipCode     string `gorm:"type:varchar(10);not null"`
	Equipment   []Equipment
	Consumables []Consumables
	Jobs        []Job
	Tags        []Tag `gorm:"many2many:location_tags;"`
}
