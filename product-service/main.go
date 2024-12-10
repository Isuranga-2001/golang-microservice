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
 
    router.POST("/orders", controllers.CreateProduct)
    router.GET("/orders", controllers.GetProducts)
    router.GET("/orders/:id", controllers.GetProduct)
    router.PUT("/orders/:id", controllers.UpdateProduct)
    router.DELETE("/orders/:id", controllers.DeleteProduct)

    log.Println("Server running on port 8080")
    router.Run(":8080")
}