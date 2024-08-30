package services

import (
	"context"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/ports"
)

type ScheduleService struct {
	repo ports.SchedulePort
}

func NewScheduleService(repo ports.SchedulePort) *ScheduleService {
	return &ScheduleService{repo: repo}
}

func (s *ScheduleService) CreateSchedule(ctx context.Context, schedule *domain.Schedule) error {
	return s.repo.Create(ctx, schedule)
}

func (s *ScheduleService) GetScheduleByID(ctx context.Context, id uint) (*domain.Schedule, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ScheduleService) UpdateSchedule(ctx context.Context, schedule *domain.Schedule) error {
	return s.repo.Update(ctx, schedule)
}

func (s *ScheduleService) DeleteSchedule(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *ScheduleService) ListSchedulesByTechnicianID(ctx context.Context, technicianID uint) ([]domain.Schedule, error) {
	return s.repo.FindAllByTechnicianID(ctx, technicianID)
}
