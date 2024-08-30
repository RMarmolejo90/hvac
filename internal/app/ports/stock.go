package ports

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
)

type StockPort interface {
    Create(ctx context.Context, stock *domain.Stock) error
    FindByID(ctx context.Context, id uint) (*domain.Stock, error)
    Update(ctx context.Context, stock *domain.Stock) error
    Delete(ctx context.Context, id uint) error
    FindAllByTechnicianID(ctx context.Context, technicianID uint) ([]domain.Stock, error)
}
