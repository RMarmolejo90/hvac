package ports

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
)

type HourlyRatePort interface {
    Create(ctx context.Context, hourlyRate *domain.HourlyRate) error
    FindByID(ctx context.Context, id uint) (*domain.HourlyRate, error)
    Update(ctx context.Context, hourlyRate *domain.HourlyRate) error
    Delete(ctx context.Context, id uint) error
    FindAll(ctx context.Context) ([]domain.HourlyRate, error)
}
