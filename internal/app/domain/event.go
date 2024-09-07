package domain

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ScheduleID     uint      `gorm:"not null"` // Foreign key to Schedule
	Schedule       Schedule  `gorm:"foreignkey:ScheduleID"`
	Title          string    `gorm:"type:varchar(100);not null"` // e.g., "Day Off", "Training Session"
	Description    string    `gorm:"type:text"`                  // Detailed description
	EventType      string    `gorm:"type:varchar(50);not null"`  // e.g., "PTO", "Training", "Meeting"
	StartTime      time.Time `gorm:"not null"`                   // When the event starts
	EndTime        time.Time `gorm:"not null"`                   // When the event ends
	IsAllDay       bool      `gorm:"default:false"`              // Indicates if the event spans the entire day
	Recurring      bool      `gorm:"default:false"`              // Indicates if the event is recurring
	RecurrenceRule string    `gorm:"type:varchar(255)"`          // e.g., "FREQ=WEEKLY;BYDAY=MO,WE,FR"
}
