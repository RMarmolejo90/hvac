package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"gorm.io/gorm"
)

type EquipmentRepository struct {
	db *gorm.DB
}

func NewEquipmentRepository(db *gorm.DB) *EquipmentRepository {
	return &EquipmentRepository{db: db}
}

func (r *EquipmentRepository) Create(ctx context.Context, equipment *domain.Equipment) error {
	return r.db.WithContext(ctx).Create(equipment).Error
}

func (r *EquipmentRepository) FindByID(ctx context.Context, id uint) (*domain.Equipment, error) {
	var equipment domain.Equipment
	if err := r.db.WithContext(ctx).Preload("Consumables").First(&equipment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &equipment, nil
}

func (r *EquipmentRepository) Update(ctx context.Context, equipment *domain.Equipment) error {
	return r.db.WithContext(ctx).Save(equipment).Error
}

func (r *EquipmentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Equipment{}, id).Error
}

func (r *EquipmentRepository) FindAllByLocationID(ctx context.Context, locationID uint) ([]domain.Equipment, error) {
	var equipment []domain.Equipment
	if err := r.db.WithContext(ctx).Where("location_id = ?", locationID).Preload("Consumables").Find(&equipment).Error; err != nil {
		return nil, err
	}
	return equipment, nil
}
