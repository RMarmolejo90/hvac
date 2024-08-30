package ports

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
)

type NotePort interface {
    Create(ctx context.Context, note *domain.Note) error
    FindByID(ctx context.Context, id uint) (*domain.Note, error)
    Update(ctx context.Context, note *domain.Note) error
    Delete(ctx context.Context, id uint) error
    FindAllByCustomerID(ctx context.Context, customerID uint) ([]domain.Note, error)
    FindAllByLocationID(ctx context.Context, locationID uint) ([]domain.Note, error)
    FindAllByJobID(ctx context.Context, jobID uint) ([]domain.Note, error)
}
