package ports

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
)

type EquipmentPort interface {
    Create(ctx context.Context, equipment *domain.Equipment) error
    FindByID(ctx context.Context, id uint) (*domain.Equipment, error)
    Update(ctx context.Context, equipment *domain.Equipment) error
    Delete(ctx context.Context, id uint) error
    FindAllByLocationID(ctx context.Context, locationID uint) ([]domain.Equipment, error)
}
