package ports

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
)

type ConsumablesPort interface {
    Create(ctx context.Context, consumable *domain.Consumables) error
    FindByID(ctx context.Context, id uint) (*domain.Consumables, error)
    Update(ctx context.Context, consumable *domain.Consumables) error
    Delete(ctx context.Context, id uint) error
    FindAllByEquipmentID(ctx context.Context, equipmentID uint) ([]domain.Consumables, error)
}
