package domain

import "gorm.io/gorm"

type ContactMethod struct {
	gorm.Model
	CustomerID  uint     `gorm:"not null"` // Foreign key to Customer
	Customer    Customer `gorm:"foreignkey:CustomerID"`
	MethodType  string   `gorm:"type:varchar(50);not null"`  // e.g., "Phone", "Email", "Fax"
	MethodValue string   `gorm:"type:varchar(100);not null"` // e.g., phone number, email address
	IsPreferred bool     `gorm:"default:false"`              // Indicates if this is the preferred contact method
}
