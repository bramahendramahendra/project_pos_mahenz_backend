package routes

import (
	"project/category-api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1/categories")
	{
		v1.GET("/", handler.GetAllCategories)
		v1.GET("/:id", handler.GetCategoryByID)
		v1.POST("/", handler.CreateCategory)
		v1.PUT("/:id", handler.UpdateCategory)
		v1.DELETE("/:id", handler.DeleteCategory)
		v1.DELETE("/permanently/:id", handler.DeletePermanentlyCategory)
	}

	return router
}
