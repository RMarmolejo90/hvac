package services

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
    "github.com/rmarmolejo90/hvac/internal/app/ports"
)

type TagService struct {
    repo ports.TagPort
}

func NewTagService(repo ports.TagPort) *TagService {
    return &TagService{repo: repo}
}

func (s *TagService) CreateTag(ctx context.Context, tag *domain.Tag) error {
    return s.repo.Create(ctx, tag)
}

func (s *TagService) GetTagByID(ctx context.Context, id uint) (*domain.Tag, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *TagService) UpdateTag(ctx context.Context, tag *domain.Tag) error {
    return s.repo.Update(ctx, tag)
}

func (s *TagService) DeleteTag(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *TagService) ListTags(ctx context.Context) ([]domain.Tag, error) {
    return s.repo.FindAll(ctx)
}
