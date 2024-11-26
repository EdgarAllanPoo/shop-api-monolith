package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdgarAllanPoo/shop-api-monolith/database"
	"github.com/EdgarAllanPoo/shop-api-monolith/models"
)

func GetProducts(c *gin.Context) {
	category := c.Query("category")
	var products []models.Product

	query := database.DB
	if category != "" {
		query = query.Where("category = ?", category)
	}

	if err := query.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
