package services

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
    "github.com/rmarmolejo90/hvac/internal/app/ports"
)

type ConsumablesService struct {
    repo ports.ConsumablesPort
}

func NewConsumablesService(repo ports.ConsumablesPort) *ConsumablesService {
    return &ConsumablesService{repo: repo}
}

func (s *ConsumablesService) CreateConsumable(ctx context.Context, consumable *domain.Consumables) error {
    return s.repo.Create(ctx, consumable)
}

func (s *ConsumablesService) GetConsumableByID(ctx context.Context, id uint) (*domain.Consumables, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *ConsumablesService) UpdateConsumable(ctx context.Context, consumable *domain.Consumables) error {
    return s.repo.Update(ctx, consumable)
}

func (s *ConsumablesService) DeleteConsumable(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *ConsumablesService) ListConsumablesByEquipmentID(ctx context.Context, equipmentID uint) ([]domain.Consumables, error) {
    return s.repo.FindAllByEquipmentID(ctx, equipmentID)
}
