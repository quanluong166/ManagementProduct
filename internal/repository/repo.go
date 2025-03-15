package repository

import "productManagement/pkg/dbsvc"

type Repository struct {
	CategoryRepository CategoryRepository
	ProductRepository  ProductRepository
}

func NewRepository(database dbsvc.Database) *Repository {
	return &Repository{
		CategoryRepository: NewCategoryRepository(database),
		ProductRepository:  NewProductRepository(database),
	}
}
