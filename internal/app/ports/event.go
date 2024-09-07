package ports

import (
	"context"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
)

type EventPort interface {
	Create(ctx context.Context, event *domain.Event) error
	FindByID(ctx context.Context, id uint) (*domain.Event, error)
	Update(ctx context.Context, event *domain.Event) error
	Delete(ctx context.Context, id uint) error
	FindAllByScheduleID(ctx context.Context, scheduleID uint) ([]domain.Event, error)
}
