package services

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
    "github.com/rmarmolejo90/hvac/internal/app/ports"
)

type QuoteService struct {
    repo ports.QuotePort
}

func NewQuoteService(repo ports.QuotePort) *QuoteService {
    return &QuoteService{repo: repo}
}

func (s *QuoteService) CreateQuote(ctx context.Context, quote *domain.Quote) error {
    return s.repo.Create(ctx, quote)
}

func (s *QuoteService) GetQuoteByID(ctx context.Context, id uint) (*domain.Quote, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *QuoteService) UpdateQuote(ctx context.Context, quote *domain.Quote) error {
    return s.repo.Update(ctx, quote)
}

func (s *QuoteService) DeleteQuote(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *QuoteService) ListQuotes(ctx context.Context) ([]domain.Quote, error) {
    return s.repo.FindAll(ctx)
}
