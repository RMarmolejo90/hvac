package domain

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null"`
	Description string
	Stock       []Stock `gorm:"many2many:stock_used"`
	Cost        float64
	Price       float64
	Jobs        []Job `gorm:"many2many:job_services;"`
}
