package ports

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
)

type ServicePort interface {
    Create(ctx context.Context, service *domain.Service) error
    FindByID(ctx context.Context, id uint) (*domain.Service, error)
    Update(ctx context.Context, service *domain.Service) error
    Delete(ctx context.Context, id uint) error
    FindAll(ctx context.Context) ([]domain.Service, error)
}
