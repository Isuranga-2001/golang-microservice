package main

import (
	"log"
	"product-service/config"
	"product-service/controllers"
	"product-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize database connection
    config.ConnectDB()

    // Auto migrate the Student model
    config.DB.AutoMigrate(&models.Product{})
    config.DB.AutoMigrate(&models.ProductType{})

    // Set up Gin router
    router := gin.Default()
 
    router.POST("/products", controllers.CreateProduct)
    router.GET("/products", controllers.GetProducts)
    router.GET("/products/:id", controllers.GetProduct)
    router.PUT("/products/:id", controllers.UpdateProduct)
    router.DELETE("/products/:id", controllers.DeleteProduct)

    log.Println("Server running on port 8081")
    router.Run(":8081")
}