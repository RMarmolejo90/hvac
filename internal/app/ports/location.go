package ports

import (
	"context"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
)

type LocationPort interface {
	Create(ctx context.Context, location *domain.Location) error
	FindByID(ctx context.Context, id uint) (*domain.Location, error)
	Update(ctx context.Context, location *domain.Location) error
	Delete(ctx context.Context, id uint) error
	FindAllByCustomerID(ctx context.Context, customerID uint) ([]domain.Location, error)
}
