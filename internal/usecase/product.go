package usecase

import (
	"context"
	api "productManagement/api/interface"
	"productManagement/internal/models"
	"productManagement/internal/repository"
)

type ProductUseCase interface {
	GetProductByFilter(ctx context.Context, req *api.GetProductsByFilterRequest) ([]*models.Product, int64, error)
	ExportProduct(ctx context.Context, req *api.ExportProductRequest) ([]*models.Product, error)
}

type productUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		repo: repo,
	}
}

func (p *productUseCase) GetProductByFilter(ctx context.Context, req *api.GetProductsByFilterRequest) ([]*models.Product, int64, error) {
	products, total, err := p.repo.GetProductByFilter(ctx, req.Status, req.Paging.Limit, req.Paging.Offset)
	if err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func (p *productUseCase) ExportProduct(ctx context.Context, req *api.ExportProductRequest) ([]*models.Product, error) {
	products, err := p.repo.GetProducts(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	return products, nil
}
