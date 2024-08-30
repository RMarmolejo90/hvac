package services

import (
	"context"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/ports"
)

type TechnicianService struct {
	repo ports.TechnicianPort
}

func NewTechnicianService(repo ports.TechnicianPort) *TechnicianService {
	return &TechnicianService{repo: repo}
}

func (s *TechnicianService) CreateTechnician(ctx context.Context, technician *domain.Technician) error {
	return s.repo.Create(ctx, technician)
}

func (s *TechnicianService) GetTechnicianByID(ctx context.Context, id uint) (*domain.Technician, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *TechnicianService) UpdateTechnician(ctx context.Context, technician *domain.Technician) error {
	return s.repo.Update(ctx, technician)
}

func (s *TechnicianService) DeleteTechnician(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *TechnicianService) ListTechnicians(ctx context.Context) ([]domain.Technician, error) {
	return s.repo.FindAll(ctx)
}
