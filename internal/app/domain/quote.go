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
}
