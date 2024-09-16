package main

import (
	categoryRoutes "project/category-api/routes"
	"project/config"
	productRoutes "project/product-api/routes"

	_ "project/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Point of Sale API
// @version 1.0
// @description This is a sample server for managing Point of Sale.
// @host localhost:8080
// @BasePath /
func main() {
	config.InitDB()

	router := gin.Default()
	categoryRoutes.RegisterRoutes(router)
	productRoutes.RegisterRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
