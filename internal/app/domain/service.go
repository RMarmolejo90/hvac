package domain
import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null"`
	Description string
	Cost        float64
	Jobs        []Job `gorm:"many2many:job_services;"`
}
