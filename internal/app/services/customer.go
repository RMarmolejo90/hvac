package services

import (
	"context"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/ports"
)

type CustomerService struct {
	repo ports.CustomerPort
}

func NewCustomerService(repo ports.CustomerPort) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, customer *domain.Customer) error {
	return s.repo.Create(ctx, customer)
}

func (s *CustomerService) GetCustomerByID(ctx context.Context, id uint) (*domain.Customer, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, customer *domain.Customer) error {
	return s.repo.Update(ctx, customer)
}

func (s *CustomerService) DeleteCustomer(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *CustomerService) ListCustomers(ctx context.Context) ([]domain.Customer, error) {
	return s.repo.FindAll(ctx)
}
