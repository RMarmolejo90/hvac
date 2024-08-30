package domain
import "gorm.io/gorm"

type Note struct {
	gorm.Model
	CustomerID uint   `gorm:"default:null"`
	LocationID uint   `gorm:"default:null"`
	JobID      uint   `gorm:"default:null"`
	Note       string `gorm:"type:text;not null"`
}
