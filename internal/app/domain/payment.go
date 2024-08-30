package domain
import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	InvoiceID     uint
	PaymentDate   time.Time
	Amount        float64 `gorm:"not null;check:amount >= 0"`
	PaymentMethod string  `gorm:"type:varchar(50);not null"`
}
