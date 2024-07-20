package main

import (
	categoryRoutes "project/category-api/routes"
	"project/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	router := gin.Default()
	categoryRoutes.RegisterRoutes(router)
	router.Run(":8081")
}
