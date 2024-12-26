package controllers

import (
    "inventory-backend/config"
    "inventory-backend/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

// GET /inventory - Lista todos os itens
func GetItems(c *gin.Context) {
    var items []models.Item
    config.DB.Find(&items)
    c.JSON(http.StatusOK, items)
}

// POST /inventory - Cria um novo item
func CreateItem(c *gin.Context) {
    var newItem models.Item

    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Create(&newItem)
    c.JSON(http.StatusCreated, newItem)
}
