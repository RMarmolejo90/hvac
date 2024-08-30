package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name      string     `gorm:"type:varchar(50);not null"`
	HexColor  string     `gorm:"type:varchar(7);not null"`
	Jobs      []Job      `gorm:"many2many:job_tags;"`
	Locations []Location `gorm:"many2many:location_tags;"`
}
