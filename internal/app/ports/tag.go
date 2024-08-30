package ports

import (
	"context"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
)

type TagPort interface {
	Create(ctx context.Context, tag *domain.Tag) error
	FindByID(ctx context.Context, id uint) (*domain.Tag, error)
	Update(ctx context.Context, tag *domain.Tag) error
	Delete(ctx context.Context, id uint) error
	FindAll(ctx context.Context) ([]domain.Tag, error)
}
