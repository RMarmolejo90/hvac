package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/postgresDB"
	"gorm.io/gorm"
)

type ConsumablesRepository struct {
	db *gorm.DB
}

func NewConsumablesRepository(db *gorm.DB) *ConsumablesRepository {
	return &ConsumablesRepository{db: postgresDB.DB}
}

func (r *ConsumablesRepository) Create(ctx context.Context, consumable *domain.Consumables) error {
	return r.db.WithContext(ctx).Create(consumable).Error
}

func (r *ConsumablesRepository) FindByID(ctx context.Context, id uint) (*domain.Consumables, error) {
	var consumable domain.Consumables
	if err := r.db.WithContext(ctx).First(&consumable, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &consumable, nil
}

func (r *ConsumablesRepository) Update(ctx context.Context, consumable *domain.Consumables) error {
	return r.db.WithContext(ctx).Save(consumable).Error
}

func (r *ConsumablesRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Consumables{}, id).Error
}

func (r *ConsumablesRepository) FindAllByEquipmentID(ctx context.Context, equipmentID uint) ([]domain.Consumables, error) {
	var consumables []domain.Consumables
	if err := r.db.WithContext(ctx).Where("equipment_id = ?", equipmentID).Find(&consumables).Error; err != nil {
		return nil, err
	}
	return consumables, nil
}
