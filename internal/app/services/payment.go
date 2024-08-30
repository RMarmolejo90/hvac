package services

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
    "github.com/rmarmolejo90/hvac/internal/app/ports"
)

type PaymentService struct {
    repo ports.PaymentPort
}

func NewPaymentService(repo ports.PaymentPort) *PaymentService {
    return &PaymentService{repo: repo}
}

func (s *PaymentService) CreatePayment(ctx context.Context, payment *domain.Payment) error {
    return s.repo.Create(ctx, payment)
}

func (s *PaymentService) GetPaymentByID(ctx context.Context, id uint) (*domain.Payment, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *PaymentService) UpdatePayment(ctx context.Context, payment *domain.Payment) error {
    return s.repo.Update(ctx, payment)
}

func (s *PaymentService) DeletePayment(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *PaymentService) ListPaymentsByInvoiceID(ctx context.Context, invoiceID uint) ([]domain.Payment, error) {
    return s.repo.FindAllByInvoiceID(ctx, invoiceID)
}
