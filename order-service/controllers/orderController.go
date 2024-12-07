package controllers

import (
	"net/http"
	"order-service/config"
	"order-service/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&order)
    c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
    var orders []models.Order
    config.DB.Find(&orders)
    c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
    var order models.Order
    if err := config.DB.First(&order, "id = ?", c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    c.JSON(http.StatusOK, order)
}

func UpdateOrder(c *gin.Context) {
    var order models.Order
    if err := config.DB.First(&order, "id = ?", c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&order)
    c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
    var order models.Order
    if err := config.DB.First(&order, "id = ?", c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    config.DB.Delete(&order)
    c.JSON(http.StatusNoContent, gin.H{"message": "Order deleted"})
}