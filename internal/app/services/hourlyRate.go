package services

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
    "github.com/rmarmolejo90/hvac/internal/app/ports"
)

type HourlyRateService struct {
    repo ports.HourlyRatePort
}

func NewHourlyRateService(repo ports.HourlyRatePort) *HourlyRateService {
    return &HourlyRateService{repo: repo}
}

func (s *HourlyRateService) CreateHourlyRate(ctx context.Context, hourlyRate *domain.HourlyRate) error {
    return s.repo.Create(ctx, hourlyRate)
}

func (s *HourlyRateService) GetHourlyRateByID(ctx context.Context, id uint) (*domain.HourlyRate, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *HourlyRateService) UpdateHourlyRate(ctx context.Context, hourlyRate *domain.HourlyRate) error {
    return s.repo.Update(ctx, hourlyRate)
}

func (s *HourlyRateService) DeleteHourlyRate(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *HourlyRateService) ListHourlyRates(ctx context.Context) ([]domain.HourlyRate, error) {
    return s.repo.FindAll(ctx)
}
