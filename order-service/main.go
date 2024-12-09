package main

import (
	"log"
	"order-service/config"
	"order-service/controllers"
	"order-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize database connection
    config.ConnectDB()

    // Auto migrate the Student model
    config.DB.AutoMigrate(&models.Order{})
    config.DB.AutoMigrate(&models.OrderItem{})

    // Set up Gin router
    router := gin.Default()
 
    router.POST("/orders", controllers.CreateOrder)
    router.GET("/orders", controllers.GetOrders)
    router.GET("/orders/:id", controllers.GetOrder)
    router.PUT("/orders/:id", controllers.UpdateOrder)
    router.DELETE("/orders/:id", controllers.DeleteOrder)

    log.Println("Server running on port 8080")
    router.Run(":8080")
}