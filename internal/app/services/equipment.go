package services

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
    "github.com/rmarmolejo90/hvac/internal/app/ports"
)

type EquipmentService struct {
    repo ports.EquipmentPort
}

func NewEquipmentService(repo ports.EquipmentPort) *EquipmentService {
    return &EquipmentService{repo: repo}
}

func (s *EquipmentService) CreateEquipment(ctx context.Context, equipment *domain.Equipment) error {
    return s.repo.Create(ctx, equipment)
}

func (s *EquipmentService) GetEquipmentByID(ctx context.Context, id uint) (*domain.Equipment, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *EquipmentService) UpdateEquipment(ctx context.Context, equipment *domain.Equipment) error {
    return s.repo.Update(ctx, equipment)
}

func (s *EquipmentService) DeleteEquipment(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *EquipmentService) ListEquipmentByLocationID(ctx context.Context, locationID uint) ([]domain.Equipment, error) {
    return s.repo.FindAllByLocationID(ctx, locationID)
}
