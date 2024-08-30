package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/postgresDB"
	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{db: postgresDB.DB}
}

func (r *JobRepository) Create(ctx context.Context, job *domain.Job) error {
	return r.db.WithContext(ctx).Create(job).Error
}

func (r *JobRepository) FindByID(ctx context.Context, id uint) (*domain.Job, error) {
	var job domain.Job
	if err := r.db.WithContext(ctx).Preload("StatusHistory").Preload("Services").Preload("Tags").Preload("Invoices").Preload("Quotes").Preload("Schedules").First(&job, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &job, nil
}

func (r *JobRepository) Update(ctx context.Context, job *domain.Job) error {
	return r.db.WithContext(ctx).Save(job).Error
}

func (r *JobRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Job{}, id).Error
}

func (r *JobRepository) FindAll(ctx context.Context) ([]domain.Job, error) {
	var jobs []domain.Job
	if err := r.db.WithContext(ctx).Preload("StatusHistory").Preload("Services").Preload("Tags").Preload("Invoices").Preload("Quotes").Preload("Schedules").Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}
