package models

import (
	"time"

	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model
	LocationID    uint
	Brand         string
	Type          string
	InstallDate   time.Time
	InstalledByUs bool
	ModelNumber   string
	Serial        string
	Name          string `gorm:"type:varchar(100);not null"`
	Location      string
	Notes         string
	Consumables   []Consumables
}
