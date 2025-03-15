package api

import (
	"productManagement/internal/models"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetProductByFilter(ctx *gin.Context)
	ExportProduct(ctx *gin.Context)
}

type GetProductsByFilterRequest struct {
	Status string        `json:"status"`
	Paging models.Paging `json:"paging"`
}

type ExportProductRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
