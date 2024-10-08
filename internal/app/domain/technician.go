package domain

import (
	"time"

	"gorm.io/gorm"
)

type Technician struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
	Email     string `gorm:"type:varchar(100);not null;unique"`
	Phone     string `gorm:"type:varchar(20);not null"`
	HireDate  time.Time
	Status    string     `gorm:"type:varchar(20);not null"`
	Jobs      []Job      // One-to-many relationship with Jobs
	Schedules []Schedule // One-to-many relationship with Schedules
}
