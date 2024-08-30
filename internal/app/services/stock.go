package services

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
    "github.com/rmarmolejo90/hvac/internal/app/ports"
)

type StockService struct {
    repo ports.StockPort
}

func NewStockService(repo ports.StockPort) *StockService {
    return &StockService{repo: repo}
}

func (s *StockService) CreateStock(ctx context.Context, stock *domain.Stock) error {
    return s.repo.Create(ctx, stock)
}

func (s *StockService) GetStockByID(ctx context.Context, id uint) (*domain.Stock, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *StockService) UpdateStock(ctx context.Context, stock *domain.Stock) error {
    return s.repo.Update(ctx, stock)
}

func (s *StockService) DeleteStock(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *StockService) ListStockByTechnicianID(ctx context.Context, technicianID uint) ([]domain.Stock, error) {
    return s.repo.FindAllByTechnicianID(ctx, technicianID)
}
