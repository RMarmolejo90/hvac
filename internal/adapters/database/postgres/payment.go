package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

func (r *PaymentRepository) FindByID(ctx context.Context, id uint) (*domain.Payment, error) {
	var payment domain.Payment
	if err := r.db.WithContext(ctx).First(&payment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &payment, nil
}

func (r *PaymentRepository) Update(ctx context.Context, payment *domain.Payment) error {
	return r.db.WithContext(ctx).Save(payment).Error
}

func (r *PaymentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Payment{}, id).Error
}

func (r *PaymentRepository) FindAllByInvoiceID(ctx context.Context, invoiceID uint) ([]domain.Payment, error) {
	var payments []domain.Payment
	if err := r.db.WithContext(ctx).Where("invoice_id = ?", invoiceID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
