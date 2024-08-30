package postgres

import (
	"context"
	"errors"

	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) Create(ctx context.Context, tag *domain.Tag) error {
	return r.db.WithContext(ctx).Create(tag).Error
}

func (r *TagRepository) FindByID(ctx context.Context, id uint) (*domain.Tag, error) {
	var tag domain.Tag
	if err := r.db.WithContext(ctx).Preload("Jobs").Preload("Locations").First(&tag, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

func (r *TagRepository) Update(ctx context.Context, tag *domain.Tag) error {
	return r.db.WithContext(ctx).Save(tag).Error
}

func (r *TagRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Tag{}, id).Error
}

func (r *TagRepository) FindAll(ctx context.Context) ([]domain.Tag, error) {
	var tags []domain.Tag
	if err := r.db.WithContext(ctx).Preload("Jobs").Preload("Locations").Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}
