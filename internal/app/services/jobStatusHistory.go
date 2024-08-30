package services

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
    "github.com/rmarmolejo90/hvac/internal/app/ports"
)

type JobStatusHistoryService struct {
    repo ports.JobStatusHistoryPort
}

func NewJobStatusHistoryService(repo ports.JobStatusHistoryPort) *JobStatusHistoryService {
    return &JobStatusHistoryService{repo: repo}
}

func (s *JobStatusHistoryService) CreateJobStatusHistory(ctx context.Context, history *domain.JobStatusHistory) error {
    return s.repo.Create(ctx, history)
}

func (s *JobStatusHistoryService) ListJobStatusHistoryByJobID(ctx context.Context, jobID uint) ([]domain.JobStatusHistory, error) {
    return s.repo.FindByJobID(ctx, jobID)
}
