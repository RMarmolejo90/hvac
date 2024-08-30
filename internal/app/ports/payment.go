package ports

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
)

type PaymentPort interface {
    Create(ctx context.Context, payment *domain.Payment) error
    FindByID(ctx context.Context, id uint) (*domain.Payment, error)
    Update(ctx context.Context, payment *domain.Payment) error
    Delete(ctx context.Context, id uint) error
    FindAllByInvoiceID(ctx context.Context, invoiceID uint) ([]domain.Payment, error)
}
