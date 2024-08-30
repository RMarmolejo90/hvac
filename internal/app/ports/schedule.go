package ports

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
)

type SchedulePort interface {
    Create(ctx context.Context, schedule *domain.Schedule) error
    FindByID(ctx context.Context, id uint) (*domain.Schedule, error)
    Update(ctx context.Context, schedule *domain.Schedule) error
    Delete(ctx context.Context, id uint) error
    FindAllByTechnicianID(ctx context.Context, technicianID uint) ([]domain.Schedule, error)
}
