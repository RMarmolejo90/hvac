package ports

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
)

type JobStatusHistoryPort interface {
    Create(ctx context.Context, history *domain.JobStatusHistory) error
    FindByJobID(ctx context.Context, jobID uint) ([]domain.JobStatusHistory, error)
}
