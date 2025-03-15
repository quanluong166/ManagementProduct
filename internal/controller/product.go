package controller

import (
	"errors"
	"net/http"
	api "productManagement/api/interface"
	"productManagement/internal/models"
	"productManagement/internal/usecase"
	"productManagement/internal/utils"

	"github.com/gin-gonic/gin"
)

type productController struct {
	p usecase.ProductUseCase
}

func NewProductController(p usecase.ProductUseCase) api.ProductController {
	return &productController{
		p: p,
	}
}

// @Summary Get product by price and category filter
// @Description Get list of products by filter price and category
// @Tags Products
// @Accept json
// @Param body body model.GetProductsByFilterRequest true "Get product by filter"
// @Produce json
// @Success 200 {object} string "Get product successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Product not found"
// @Failure 500 {object} string "Internal server error"
// @Router /products/filter [post]
func (c *productController) GetProductByFilter(ctx *gin.Context) {
	var req *api.GetProductsByFilterRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ResponseWithErrorAndMessage(http.StatusBadRequest, err, ctx)
		return
	}

	if _, ok := models.StatusMapping[req.Status]; !ok {
		ResponseWithErrorAndMessage(http.StatusBadRequest, errors.New("INVALID_STATUS"), ctx)
	}

	if _, ok := models.CategoryMapping[req.Category]; !ok {
		ResponseWithErrorAndMessage(http.StatusBadRequest, errors.New("INVALID_CATEGORY"), ctx)
	}

	products, total, err := c.p.GetProductByFilter(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var res []*api.ProductDTO

	for _, product := range products {
		dto := c.convertToDTO(product)
		res = append(res, dto)
	}

	ctx.JSON(200, gin.H{"products": res, "total": total})
}

func (c *productController) ExportProduct(ctx *gin.Context) {
	var req *api.GetProductsByFilterRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ResponseWithErrorAndMessage(http.StatusBadRequest, err, ctx)
		return
	}

	products, _, err := c.p.GetProductByFilter(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	filePath, err := utils.GenerateProductPDF(products)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
		return
	}

	// Serve the PDF file
	ctx.File(filePath)
}

func (c *productController) convertToDTO(product *models.Product) *api.ProductDTO {
	dateFormat := utils.ConvertToDateOnly(product.AddedDate)
	productDTO := &api.ProductDTO{
		ID:        product.ID,
		Reference: product.Reference,
		Name:      product.Name,
		Price:     product.Price,
		Status:    product.Status,
		StockCity: product.StockCity,
		AddedDate: dateFormat,
		Quantity:  product.Quantity,
	}

	return productDTO
}
