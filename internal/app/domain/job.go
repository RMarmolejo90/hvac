package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	CustomerID     uint
	LocationID     uint
	TechnicianID   uint
	ScheduleID     uint // Foreign key to the Schedule model
	ScheduledDate  time.Time
	CompletionDate *time.Time
	Status         string `gorm:"type:varchar(20);not null;check:status IN ('scheduled', 'working', 'paused', 'delayed', 'rescheduled', 'cancelled', 'needs quote', 'quote sent', 'quote approved', 'awaiting parts', 'ready to schedule', 'requires return', 'completed')"`
	Notes          string
	MaterialsCost  float64
	LaborCost      float64
	StatusHistory  []JobStatusHistory
	Services       []Service `gorm:"many2many:job_services;"`
	Tags           []Tag     `gorm:"many2many:job_tags;"`
	Invoices       []Invoice
	Quotes         []Quote
	Schedules      []Schedule // Relationship to Schedule
}
