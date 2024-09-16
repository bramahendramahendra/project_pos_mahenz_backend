package routes

import (
	"project/product-api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	r := router.Group("/products")
	{
		r.GET("/", handler.GetAllProducts)
		r.GET("/:id", handler.GetProductByID)
		r.GET("/category/:id_category", handler.GetProductsByCategoryID) // Tambahkan ini
		r.POST("/", handler.CreateProduct)
		r.PUT("/:id", handler.UpdateProduct)
		r.DELETE("/:id", handler.DeleteProduct)
		r.DELETE("/permanently/:id", handler.DeletePermanentlyProduct)
		r.GET("/with-deleted", handler.GetAllProductsWithDeleted)
	}
}
