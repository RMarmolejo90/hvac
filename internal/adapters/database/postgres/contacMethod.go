package postgres

import (
	"context"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/postgresDB"
	"gorm.io/gorm"
)

type ContactMethodRepository struct {
	db *gorm.DB
}

func NewContactMethodRepository(db *gorm.DB) *ContactMethodRepository {
	return &ContactMethodRepository{db: postgresDB.DB}
}

func (r *ContactMethodRepository) Create(ctx context.Context, contactMethod *domain.ContactMethod) error {
	return r.db.WithContext(ctx).Create(contactMethod).Error
}

func (r *ContactMethodRepository) FindByID(ctx context.Context, id uint) (*domain.ContactMethod, error) {
	var contactMethod domain.ContactMethod
	if err := r.db.WithContext(ctx).First(&contactMethod, id).Error; err != nil {
		return nil, err
	}
	return &contactMethod, nil
}

func (r *ContactMethodRepository) Update(ctx context.Context, contactMethod *domain.ContactMethod) error {
	return r.db.WithContext(ctx).Save(contactMethod).Error
}

func (r *ContactMethodRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.ContactMethod{}, id).Error
}

func (r *ContactMethodRepository) FindAllByCustomerID(ctx context.Context, customerID uint) ([]domain.ContactMethod, error) {
	var contactMethods []domain.ContactMethod
	if err := r.db.WithContext(ctx).Where("customer_id = ?", customerID).Find(&contactMethods).Error; err != nil {
		return nil, err
	}
	return contactMethods, nil
}
