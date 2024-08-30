package models

import "gorm.io/gorm"

type HourlyRate struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(100);not null;unique"` // Name of the rate (e.g., "Hotel Rate", "Residential Rate")
	Description string     `gorm:"type:text"`                         // Description
	Rate        float64    `gorm:"not null"`                          // Hourly rate
	Customers   []Customer `gorm:"foreignkey:HourlyRateID"`           // Customers associated with this rate
}
