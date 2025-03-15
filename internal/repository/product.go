package repository

import (
	"context"
	"productManagement/internal/models"
	"productManagement/pkg/dbsvc"
)

type ProductRepository interface {
	GetProductByFilter(ctx context.Context, city_sort, price_sort, category, status string, limit, offset int) ([]*models.Product, int64, error)
	GetProducts(ctx context.Context, limit, offset int) ([]*models.Product, error)
}

type productRepository struct {
	db dbsvc.Database
}

func NewProductRepository(db dbsvc.Database) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) GetProductByFilter(ctx context.Context, city_sort, price_sort, category, status string, limit, offset int) ([]*models.Product, int64, error) {
	var products []*models.Product
	query := r.db.GormDB.Model(&models.Product{})

	if len(status) > 0 {
		query = query.Where("status = ?", status)
	}

	if len(category) > 0 {
		query = query.Where("category = ?", category)
	}

	total := query.Find(&products).RowsAffected

	if limit == 0 {
		limit = 10
	}

	if len(price_sort) > 0 {
		query = query.Order("price " + price_sort)
	}

	if len(city_sort) > 0 {
		query = query.Order("stock_city " + city_sort)
	}

	if err := query.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *productRepository) GetProducts(ctx context.Context, limit, offset int) ([]*models.Product, error) {
	var products []*models.Product
	query := r.db.GormDB.Model(&models.Product{})

	if limit == 0 {
		limit = 10
	}

	if err := query.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
