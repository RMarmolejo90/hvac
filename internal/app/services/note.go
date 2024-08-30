package services

import (
    "context"
    "github.com/rmarmolejo90/hvac/internal/app/domain"
    "github.com/rmarmolejo90/hvac/internal/app/ports"
)

type NoteService struct {
    repo ports.NotePort
}

func NewNoteService(repo ports.NotePort) *NoteService {
    return &NoteService{repo: repo}
}

func (s *NoteService) CreateNote(ctx context.Context, note *domain.Note) error {
    return s.repo.Create(ctx, note)
}

func (s *NoteService) GetNoteByID(ctx context.Context, id uint) (*domain.Note, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *NoteService) UpdateNote(ctx context.Context, note *domain.Note) error {
    return s.repo.Update(ctx, note)
}

func (s *NoteService) DeleteNote(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *NoteService) ListNotesByCustomerID(ctx context.Context, customerID uint) ([]domain.Note, error) {
    return s.repo.FindAllByCustomerID(ctx, customerID)
}

func (s *NoteService) ListNotesByLocationID(ctx context.Context, locationID uint) ([]domain.Note, error) {
    return s.repo.FindAllByLocationID(ctx, locationID)
}

func (s *NoteService) ListNotesByJobID(ctx context.Context, jobID uint) ([]domain.Note, error) {
    return s.repo.FindAllByJobID(ctx, jobID)
}
