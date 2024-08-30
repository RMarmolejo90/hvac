package postgres

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
    "gorm.io/gorm"
)

type JobStatusHistoryRepository struct {
    db *gorm.DB
}

func NewJobStatusHistoryRepository(db *gorm.DB) *JobStatusHistoryRepository {
    return &JobStatusHistoryRepository{db: db}
}

func (r *JobStatusHistoryRepository) Create(ctx context.Context, history *domain.JobStatusHistory) error {
    return r.db.WithContext(ctx).Create(history).Error
}

func (r *JobStatusHistoryRepository) FindByJobID(ctx context.Context, jobID uint) ([]domain.JobStatusHistory, error) {
    var history []domain.JobStatusHistory
    if err := r.db.WithContext(ctx).Where("job_id = ?", jobID).Find(&history).Error; err != nil {
        return nil, err
    }
    return history, nil
}
