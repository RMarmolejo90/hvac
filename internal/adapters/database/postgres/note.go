package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"gorm.io/gorm"
)

type NoteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db: db}
}

func (r *NoteRepository) Create(ctx context.Context, note *domain.Note) error {
	return r.db.WithContext(ctx).Create(note).Error
}

func (r *NoteRepository) FindByID(ctx context.Context, id uint) (*domain.Note, error) {
	var note domain.Note
	if err := r.db.WithContext(ctx).First(&note, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &note, nil
}

func (r *NoteRepository) Update(ctx context.Context, note *domain.Note) error {
	return r.db.WithContext(ctx).Save(note).Error
}

func (r *NoteRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Note{}, id).Error
}

func (r *NoteRepository) FindAllByCustomerID(ctx context.Context, customerID uint) ([]domain.Note, error) {
	var notes []domain.Note
	if err := r.db.WithContext(ctx).Where("customer_id = ?", customerID).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *NoteRepository) FindAllByLocationID(ctx context.Context, locationID uint) ([]domain.Note, error) {
	var notes []domain.Note
	if err := r.db.WithContext(ctx).Where("location_id = ?", locationID).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *NoteRepository) FindAllByJobID(ctx context.Context, jobID uint) ([]domain.Note, error) {
	var notes []domain.Note
	if err := r.db.WithContext(ctx).Where("job_id = ?", jobID).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}
