package domain
import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	JobID           uint
	InvoiceDate     time.Time
	TotalAmount     float64 `gorm:"not null"`
	PaymentStatus   string  `gorm:"type:varchar(20);not null;check:payment_status IN ('paid', 'unpaid', 'refunded')"`
	TechnicianNotes string
	OfficeNotes     string
	Payments        []Payment
}
