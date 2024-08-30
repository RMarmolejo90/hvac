package services

import (
	"context"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/ports"
)

type ServiceService struct {
	repo ports.ServicePort
}

func NewServiceService(repo ports.ServicePort) *ServiceService {
	return &ServiceService{repo: repo}
}

func (s *ServiceService) CreateService(ctx context.Context, service *domain.Service) error {
	return s.repo.Create(ctx, service)
}

func (s *ServiceService) GetServiceByID(ctx context.Context, id uint) (*domain.Service, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ServiceService) UpdateService(ctx context.Context, service *domain.Service) error {
	return s.repo.Update(ctx, service)
}

func (s *ServiceService) DeleteService(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *ServiceService) ListServices(ctx context.Context) ([]domain.Service, error) {
	return s.repo.FindAll(ctx)
}
