package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/postgresDB"
	"gorm.io/gorm"
)

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{db: postgresDB.DB}
}

func (r *LocationRepository) Create(ctx context.Context, location *domain.Location) error {
	return r.db.WithContext(ctx).Create(location).Error
}

func (r *LocationRepository) FindByID(ctx context.Context, id uint) (*domain.Location, error) {
	var location domain.Location
	if err := r.db.WithContext(ctx).Preload("Equipment").Preload("Jobs").Preload("Tags").First(&location, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &location, nil
}

func (r *LocationRepository) Update(ctx context.Context, location *domain.Location) error {
	return r.db.WithContext(ctx).Save(location).Error
}

func (r *LocationRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Location{}, id).Error
}

func (r *LocationRepository) FindAllByCustomerID(ctx context.Context, customerID uint) ([]domain.Location, error) {
	var locations []domain.Location
	if err := r.db.WithContext(ctx).Where("customer_id = ?", customerID).Preload("Equipment").Preload("Jobs").Preload("Tags").Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}
