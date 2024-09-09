package services

import (
	"context"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/ports"
)

type EventService struct {
	eventRepo ports.EventPort
}

func NewEventService(eventRepo ports.EventPort) *EventService {
	return &EventService{eventRepo: eventRepo}
}

func (s *EventService) CreateEvent(ctx context.Context, event *domain.Event) error {
	return s.eventRepo.Create(ctx, event)
}

func (s *EventService) FindEventByID(ctx context.Context, id uint) (*domain.Event, error) {
	return s.eventRepo.FindByID(ctx, id)
}

func (s *EventService) UpdateEvent(ctx context.Context, event *domain.Event) error {
	return s.eventRepo.Update(ctx, event)
}

func (s *EventService) DeleteEvent(ctx context.Context, id uint) error {
	return s.eventRepo.Delete(ctx, id)
}

func (s *EventService) ListEventsByScheduleID(ctx context.Context, scheduleID uint) ([]domain.Event, error) {
	return s.eventRepo.FindAllByScheduleID(ctx, scheduleID)
}
