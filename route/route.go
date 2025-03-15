package route

import (
	"net/http"
	"productManagement/internal/controller"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(e *gin.Engine, c *controller.Controller) {
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	e.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	product := e.Group("/products")
	product.POST("/filter", c.ProductController.GetProductByFilter)
	product.GET("/export", c.ProductController.ExportProduct)
}
