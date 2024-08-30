package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"gorm.io/gorm"
)

type TechnicianRepository struct {
	db *gorm.DB
}

func NewTechnicianRepository(db *gorm.DB) *TechnicianRepository {
	return &TechnicianRepository{db: db}
}

func (r *TechnicianRepository) Create(ctx context.Context, technician *domain.Technician) error {
	return r.db.WithContext(ctx).Create(technician).Error
}

func (r *TechnicianRepository) FindByID(ctx context.Context, id uint) (*domain.Technician, error) {
	var technician domain.Technician
	if err := r.db.WithContext(ctx).Preload("Jobs").Preload("Schedules").First(&technician, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &technician, nil
}

func (r *TechnicianRepository) Update(ctx context.Context, technician *domain.Technician) error {
	return r.db.WithContext(ctx).Save(technician).Error
}

func (r *TechnicianRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Technician{}, id).Error
}

func (r *TechnicianRepository) FindAll(ctx context.Context) ([]domain.Technician, error) {
	var technicians []domain.Technician
	if err := r.db.WithContext(ctx).Preload("Jobs").Preload("Schedules").Find(&technicians).Error; err != nil {
		return nil, err
	}
	return technicians, nil
}
