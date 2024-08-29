package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName    string     `gorm:"type:varchar(50);not null"`
	LastName     string     `gorm:"type:varchar(50);not null"`
	Email        string     `gorm:"type:varchar(100);not null;unique"`
	Phone        string     `gorm:"type:varchar(20);not null"`
	Status       string     `gorm:"type:varchar(20);not null;check:status IN ('new', 'active', 'blacklist', 'gold')"`
	HourlyRateID uint       // Foreign key to the HourlyRate model
	HourlyRate   HourlyRate `gorm:"foreignkey:HourlyRateID"` // Association to HourlyRate
	Locations    []Location
	Jobs         []Job
}
