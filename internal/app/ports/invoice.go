package ports

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
)

type InvoicePort interface {
    Create(ctx context.Context, invoice *domain.Invoice) error
    FindByID(ctx context.Context, id uint) (*domain.Invoice, error)
    Update(ctx context.Context, invoice *domain.Invoice) error
    Delete(ctx context.Context, id uint) error
    FindAll(ctx context.Context) ([]domain.Invoice, error)
}
