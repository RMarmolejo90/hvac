package ports

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
)

type QuotePort interface {
    Create(ctx context.Context, quote *domain.Quote) error
    FindByID(ctx context.Context, id uint) (*domain.Quote, error)
    Update(ctx context.Context, quote *domain.Quote) error
    Delete(ctx context.Context, id uint) error
    FindAll(ctx context.Context) ([]domain.Quote, error)
}
