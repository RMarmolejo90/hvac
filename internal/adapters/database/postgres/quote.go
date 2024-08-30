package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"gorm.io/gorm"
)

type QuoteRepository struct {
	db *gorm.DB
}

func NewQuoteRepository(db *gorm.DB) *QuoteRepository {
	return &QuoteRepository{db: db}
}

func (r *QuoteRepository) Create(ctx context.Context, quote *domain.Quote) error {
	return r.db.WithContext(ctx).Create(quote).Error
}

func (r *QuoteRepository) FindByID(ctx context.Context, id uint) (*domain.Quote, error) {
	var quote domain.Quote
	if err := r.db.WithContext(ctx).First(&quote, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &quote, nil
}

func (r *QuoteRepository) Update(ctx context.Context, quote *domain.Quote) error {
	return r.db.WithContext(ctx).Save(quote).Error
}

func (r *QuoteRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Quote{}, id).Error
}

func (r *QuoteRepository) FindAll(ctx context.Context) ([]domain.Quote, error) {
	var quotes []domain.Quote
	if err := r.db.WithContext(ctx).Find(&quotes).Error; err != nil {
		return nil, err
	}
	return quotes, nil
}
