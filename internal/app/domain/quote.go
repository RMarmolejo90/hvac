package domain

import (
	"time"

	"gorm.io/gorm"
)

type Quote struct {
	gorm.Model
	JobID       uint
	Name        string `gorm:"type:varchar(100);not null"`
	QuoteDate   time.Time
	Price       float64 `gorm:"not null"`
	Description string  `gorm:"type:text;not null"`
	Status      string  `gorm:"type:varchar(20);not null;check:status IN ('not sent', 'awaiting response', 'approved', 'declined')"`
}
