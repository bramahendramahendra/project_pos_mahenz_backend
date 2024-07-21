package main

import (
	categoryRoutes "project/category-api/routes"
	"project/config"

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

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
func main() {
	config.InitDB()

	router := gin.Default()
	categoryRoutes.RegisterRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
