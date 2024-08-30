package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/postgresDB"
	"gorm.io/gorm"
)

type HourlyRateRepository struct {
	db *gorm.DB
}

func NewHourlyRateRepository(db *gorm.DB) *HourlyRateRepository {
	return &HourlyRateRepository{db: postgresDB.DB}
}

func (r *HourlyRateRepository) Create(ctx context.Context, hourlyRate *domain.HourlyRate) error {
	return r.db.WithContext(ctx).Create(hourlyRate).Error
}

func (r *HourlyRateRepository) FindByID(ctx context.Context, id uint) (*domain.HourlyRate, error) {
	var hourlyRate domain.HourlyRate
	if err := r.db.WithContext(ctx).First(&hourlyRate, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &hourlyRate, nil
}

func (r *HourlyRateRepository) Update(ctx context.Context, hourlyRate *domain.HourlyRate) error {
	return r.db.WithContext(ctx).Save(hourlyRate).Error
}

func (r *HourlyRateRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.HourlyRate{}, id).Error
}

func (r *HourlyRateRepository) FindAll(ctx context.Context) ([]domain.HourlyRate, error) {
	var hourlyRates []domain.HourlyRate
	if err := r.db.WithContext(ctx).Find(&hourlyRates).Error; err != nil {
		return nil, err
	}
	return hourlyRates, nil
}
