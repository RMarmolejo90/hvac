package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}

func (r *ScheduleRepository) Create(ctx context.Context, schedule *domain.Schedule) error {
	return r.db.WithContext(ctx).Create(schedule).Error
}

func (r *ScheduleRepository) FindByID(ctx context.Context, id uint) (*domain.Schedule, error) {
	var schedule domain.Schedule
	if err := r.db.WithContext(ctx).Preload("Jobs").First(&schedule, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &schedule, nil
}

func (r *ScheduleRepository) Update(ctx context.Context, schedule *domain.Schedule) error {
	return r.db.WithContext(ctx).Save(schedule).Error
}

func (r *ScheduleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Schedule{}, id).Error
}

func (r *ScheduleRepository) FindAllByTechnicianID(ctx context.Context, technicianID uint) ([]domain.Schedule, error) {
	var schedules []domain.Schedule
	if err := r.db.WithContext(ctx).Where("technician_id = ?", technicianID).Preload("Jobs").Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}
