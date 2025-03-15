package repository

import (
	"context"
	"productManagement/internal/models"
	"productManagement/pkg/dbsvc"
)

type CategoryRepository interface {
}

type categoryRepository struct {
	db dbsvc.Database
}

func NewCategoryRepository(db dbsvc.Database) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(ctx context.Context, category *models.Category) error {
	return r.db.GormDB.Create(category).Error
}
