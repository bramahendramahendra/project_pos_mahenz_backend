package routes

import (
	"project/category-api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	r := router.Group("/categories")
	{
		r.GET("/", handler.GetAllCategories)
		r.GET("/:id", handler.GetCategoryByID)
		r.POST("/", handler.CreateCategory)
		r.PUT("/:id", handler.UpdateCategory)
		r.DELETE("/:id", handler.DeleteCategory)
		r.DELETE("/permanently/:id", handler.DeletePermanentlyCategory)
	}
}
