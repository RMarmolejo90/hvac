package models

import "gorm.io/gorm"

type Consumables struct {
	gorm.Model
	EquipmentID uint
	Type        string `gorm:"type:varchar(50);not null"`
	Size        string `gorm:"type:varchar(30);not null"`
	Quantity    int    `gorm:"not null;check:quantity > 0"`
}
