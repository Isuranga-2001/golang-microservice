package controllers

import (
	"net/http"
	"product-service/config"
	"product-service/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&product)
    c.JSON(http.StatusCreated, product)
}

func GetProducts(c *gin.Context) {
    var products []models.Product
    config.DB.Find(&products)
    c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
    var product models.Product
    if err := config.DB.First(&product, "id = ?", c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
    var product models.Product
    if err := config.DB.First(&product, "id = ?", c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&product)
    c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
    var product models.Product
    if err := config.DB.First(&product, "id = ?", c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    config.DB.Delete(&product)
    c.JSON(http.StatusNoContent, gin.H{"message": "Product deleted"})
}