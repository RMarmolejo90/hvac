package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	TechnicianID uint       // Foreign key to the Technician model
	Technician   Technician `gorm:"foreignkey:TechnicianID"` // Association to Technician
	Jobs         []Job      // One-to-many relationship with Job
	Date         time.Time  `gorm:"not null"` // Date of the schedule
}
