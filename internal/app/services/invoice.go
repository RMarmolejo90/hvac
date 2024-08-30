package services

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
    "github.com/rmarmolejo90/hvac/internal/app/ports"
)

type InvoiceService struct {
    repo ports.InvoicePort
}

func NewInvoiceService(repo ports.InvoicePort) *InvoiceService {
    return &InvoiceService{repo: repo}
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, invoice *domain.Invoice) error {
    return s.repo.Create(ctx, invoice)
}

func (s *InvoiceService) GetInvoiceByID(ctx context.Context, id uint) (*domain.Invoice, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *InvoiceService) UpdateInvoice(ctx context.Context, invoice *domain.Invoice) error {
    return s.repo.Update(ctx, invoice)
}

func (s *InvoiceService) DeleteInvoice(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *InvoiceService) ListInvoices(ctx context.Context) ([]domain.Invoice, error) {
    return s.repo.FindAll(ctx)
}
