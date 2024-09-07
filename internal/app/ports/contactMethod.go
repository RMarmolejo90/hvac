package ports

import (
	"context"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
)

type ContactMethodPort interface {
	Create(ctx context.Context, contactMethod *domain.ContactMethod) error
	FindByID(ctx context.Context, id uint) (*domain.ContactMethod, error)
	Update(ctx context.Context, contactMethod *domain.ContactMethod) error
	Delete(ctx context.Context, id uint) error
	FindAllByCustomerID(ctx context.Context, customerID uint) ([]domain.ContactMethod, error)
}
