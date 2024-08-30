package services

import (
	"context"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/ports"
)

type LocationService struct {
	repo ports.LocationPort
}

func NewLocationService(repo ports.LocationPort) *LocationService {
	return &LocationService{repo: repo}
}

func (s *LocationService) CreateLocation(ctx context.Context, location *domain.Location) error {
	return s.repo.Create(ctx, location)
}

func (s *LocationService) GetLocationByID(ctx context.Context, id uint) (*domain.Location, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *LocationService) UpdateLocation(ctx context.Context, location *domain.Location) error {
	return s.repo.Update(ctx, location)
}

func (s *LocationService) DeleteLocation(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *LocationService) ListLocationsByCustomerID(ctx context.Context, customerID uint) ([]domain.Location, error) {
	return s.repo.FindAllByCustomerID(ctx, customerID)
}
