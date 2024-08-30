package ports

import (
	"context"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
)

type TechnicianPort interface {
	Create(ctx context.Context, technician *domain.Technician) error
	FindByID(ctx context.Context, id uint) (*domain.Technician, error)
	Update(ctx context.Context, technician *domain.Technician) error
	Delete(ctx context.Context, id uint) error
	FindAll(ctx context.Context) ([]domain.Technician, error)
}
