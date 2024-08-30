package ports

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
)

type JobPort interface {
    Create(ctx context.Context, job *domain.Job) error
    FindByID(ctx context.Context, id uint) (*domain.Job, error)
    Update(ctx context.Context, job *domain.Job) error
    Delete(ctx context.Context, id uint) error
    FindAll(ctx context.Context) ([]domain.Job, error)
}
