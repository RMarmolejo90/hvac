package domain
import (
	"time"

	"gorm.io/gorm"
)

type JobStatusHistory struct {
	gorm.Model
	JobID     uint
	Status    string
	ChangedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
