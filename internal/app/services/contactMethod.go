package services

import (
	"context"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/ports"
)

type ContactMethodService struct {
	contactMethodRepo ports.ContactMethodPort
}

func NewContactMethodService(contactMethodRepo ports.ContactMethodPort) *ContactMethodService {
	return &ContactMethodService{contactMethodRepo: contactMethodRepo}
}

func (s *ContactMethodService) CreateContactMethod(ctx context.Context, contactMethod *domain.ContactMethod) error {
	return s.contactMethodRepo.Create(ctx, contactMethod)
}

func (s *ContactMethodService) FindContactMethodByID(ctx context.Context, id uint) (*domain.ContactMethod, error) {
	return s.contactMethodRepo.FindByID(ctx, id)
}

func (s *ContactMethodService) UpdateContactMethod(ctx context.Context, contactMethod *domain.ContactMethod) error {
	return s.contactMethodRepo.Update(ctx, contactMethod)
}

func (s *ContactMethodService) DeleteContactMethod(ctx context.Context, id uint) error {
	return s.contactMethodRepo.Delete(ctx, id)
}

func (s *ContactMethodService) ListContactMethodsByCustomerID(ctx context.Context, customerID uint) ([]domain.ContactMethod, error) {
	return s.contactMethodRepo.FindAllByCustomerID(ctx, customerID)
}
