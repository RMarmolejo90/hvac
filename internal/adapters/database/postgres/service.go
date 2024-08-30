package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (r *ServiceRepository) Create(ctx context.Context, service *domain.Service) error {
	return r.db.WithContext(ctx).Create(service).Error
}

func (r *ServiceRepository) FindByID(ctx context.Context, id uint) (*domain.Service, error) {
	var service domain.Service
	if err := r.db.WithContext(ctx).First(&service, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &service, nil
}

func (r *ServiceRepository) Update(ctx context.Context, service *domain.Service) error {
	return r.db.WithContext(ctx).Save(service).Error
}

func (r *ServiceRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Service{}, id).Error
}

func (r *ServiceRepository) FindAll(ctx context.Context) ([]domain.Service, error) {
	var services []domain.Service
	if err := r.db.WithContext(ctx).Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}
