package api

import (
	"productManagement/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController interface {
	GetProductByFilter(ctx *gin.Context)
	ExportProduct(ctx *gin.Context)
}

type GetProductsByFilterRequest struct {
	Status    string        `json:"status"`
	Category  string        `json:"category"`
	Paging    models.Paging `json:"paging"`
	PriceSort string        `json:"price_sort"` // ASC, DESC
	CitySort  string        `json:"city_sort"`  // ASC, DESC
}

type ExportProductRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ProductDTO struct {
	ID         uuid.UUID  `json:"id"`
	Reference  string     `json:"reference"`
	Name       string     `json:"name"`
	AddedDate  string     `json:"added_date"`
	Status     string     `json:"status"`
	CategoryID *uuid.UUID `json:"category_id"`
	Price      float64    `json:"price"`
	StockCity  string     `json:"stock_city"`
	SupplierID *uuid.UUID `json:"supplier_id"`
	Quantity   int        `json:"quantity"`
}
