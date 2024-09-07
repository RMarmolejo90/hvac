package domain

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	TechnicianID uint       `gorm:"not null"`                // Foreign key to Technician
	Technician   Technician `gorm:"foreignkey:TechnicianID"` // Many-to-one relationship
	Jobs         []Job      `gorm:"foreignkey:ScheduleID"`   // One-to-many relationship with Jobs
	Events       []Event    `gorm:"foreignkey:ScheduleID"`   // One-to-many relationship with Events
	Date         time.Time  `gorm:"not null"`                // Date of the schedule
	StartTime    time.Time  `gorm:"not null"`                // Start time of the schedule
	EndTime      *time.Time `gorm:""`                        // Optional: End time of the schedule
}
