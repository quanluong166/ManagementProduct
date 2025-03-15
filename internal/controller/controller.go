package controller

import (
	api "productManagement/api/interface"
	"productManagement/internal/usecase"
	"productManagement/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ProductController api.ProductController
	// CategoryController CategoryController
}

func NewController(uc *usecase.UseCase) *Controller {
	return &Controller{
		ProductController: NewProductController(uc.ProductUseCase),
	}
}

func ResponseWithErrorAndMessage(status int, err error, c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if err != nil {
		logger.Error("Err in controller", err)
		c.AbortWithStatusJSON(status, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(status, nil)
}

func ResponseWithStatusAndData(status int, data interface{}, c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(status, data)
}
