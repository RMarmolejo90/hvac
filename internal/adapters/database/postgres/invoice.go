package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/postgresDB"
	"gorm.io/gorm"
)

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{db: postgresDB.DB}
}

func (r *InvoiceRepository) Create(ctx context.Context, invoice *domain.Invoice) error {
	return r.db.WithContext(ctx).Create(invoice).Error
}

func (r *InvoiceRepository) FindByID(ctx context.Context, id uint) (*domain.Invoice, error) {
	var invoice domain.Invoice
	if err := r.db.WithContext(ctx).Preload("Payments").First(&invoice, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &invoice, nil
}

func (r *InvoiceRepository) Update(ctx context.Context, invoice *domain.Invoice) error {
	return r.db.WithContext(ctx).Save(invoice).Error
}

func (r *InvoiceRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Invoice{}, id).Error
}

func (r *InvoiceRepository) FindAll(ctx context.Context) ([]domain.Invoice, error) {
	var invoices []domain.Invoice
	if err := r.db.WithContext(ctx).Preload("Payments").Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}
