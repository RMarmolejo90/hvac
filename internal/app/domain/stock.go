package domain

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	TechnicianID uint
	Name         string `gorm:"type:varchar(100);not null"`
	PartNumber   string `gorm:"type:varchar(50);not null"`
	Description  string
	Cost         float64
	Quantity     int
	Jobs         []Job     `gorm:"many2many:job_stock_items;foreignKey:ID;joinForeignKey:StockID;References:ID;joinReferences:JobID"`          // Stock used in jobs
	Services     []Service `gorm:"many2many:services_using_stock;foreignKey:ID;joinForeignKey:StockID;References:ID;joinReferences:ServiceID"` // Stock used in services
}
