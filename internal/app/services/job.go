package services

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
    "github.com/rmarmolejo90/hvac/internal/app/ports"
)

type JobService struct {
    repo ports.JobPort
}

func NewJobService(repo ports.JobPort) *JobService {
    return &JobService{repo: repo}
}

func (s *JobService) CreateJob(ctx context.Context, job *domain.Job) error {
    return s.repo.Create(ctx, job)
}

func (s *JobService) GetJobByID(ctx context.Context, id uint) (*domain.Job, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *JobService) UpdateJob(ctx context.Context, job *domain.Job) error {
    return s.repo.Update(ctx, job)
}

func (s *JobService) DeleteJob(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *JobService) ListJobs(ctx context.Context) ([]domain.Job, error) {
    return s.repo.FindAll(ctx)
}
