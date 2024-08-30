package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(ctx context.Context, customer *domain.Customer) error {
	return r.db.WithContext(ctx).Create(customer).Error
}

func (r *CustomerRepository) FindByID(ctx context.Context, id uint) (*domain.Customer, error) {
	var customer domain.Customer
	if err := r.db.WithContext(ctx).Preload("Locations").Preload("Jobs").First(&customer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	return r.db.WithContext(ctx).Save(customer).Error
}

func (r *CustomerRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Customer{}, id).Error
}

func (r *CustomerRepository) FindAll(ctx context.Context) ([]domain.Customer, error) {
	var customers []domain.Customer
	if err := r.db.WithContext(ctx).Preload("Locations").Preload("Jobs").Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
