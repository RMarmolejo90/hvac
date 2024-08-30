package ports

import (
	"context"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
)

type CustomerPort interface {
	Create(ctx context.Context, customer *domain.Customer) error
	FindByID(ctx context.Context, id uint) (*domain.Customer, error)
	Update(ctx context.Context, customer *domain.Customer) error
	Delete(ctx context.Context, id uint) error
	FindAll(ctx context.Context) ([]domain.Customer, error)
}
