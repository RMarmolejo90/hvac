package postgres

import (
	"context"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/postgresDB"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: postgresDB.DB}
}

func (r *EventRepository) Create(ctx context.Context, event *domain.Event) error {
	return r.db.WithContext(ctx).Create(event).Error
}

func (r *EventRepository) FindByID(ctx context.Context, id uint) (*domain.Event, error) {
	var event domain.Event
	if err := r.db.WithContext(ctx).First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *EventRepository) Update(ctx context.Context, event *domain.Event) error {
	return r.db.WithContext(ctx).Save(event).Error
}

func (r *EventRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Event{}, id).Error
}

func (r *EventRepository) FindAllByScheduleID(ctx context.Context, scheduleID uint) ([]domain.Event, error) {
	var events []domain.Event
	if err := r.db.WithContext(ctx).Where("schedule_id = ?", scheduleID).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
