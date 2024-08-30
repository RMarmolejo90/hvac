package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/postgresDB"
	"gorm.io/gorm"
)

type StockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) *StockRepository {
	return &StockRepository{db: postgresDB.DB}
}

func (r *StockRepository) Create(ctx context.Context, stock *domain.Stock) error {
	return r.db.WithContext(ctx).Create(stock).Error
}

func (r *StockRepository) FindByID(ctx context.Context, id uint) (*domain.Stock, error) {
	var stock domain.Stock
	if err := r.db.WithContext(ctx).Preload("Jobs").First(&stock, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &stock, nil
}

func (r *StockRepository) Update(ctx context.Context, stock *domain.Stock) error {
	return r.db.WithContext(ctx).Save(stock).Error
}

func (r *StockRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Stock{}, id).Error
}

func (r *StockRepository) FindAllByTechnicianID(ctx context.Context, technicianID uint) ([]domain.Stock, error) {
	var stocks []domain.Stock
	if err := r.db.WithContext(ctx).Where("technician_id = ?", technicianID).Preload("Jobs").Find(&stocks).Error; err != nil {
		return nil, err
	}
	return stocks, nil
}
