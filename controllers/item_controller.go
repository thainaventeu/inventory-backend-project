package controllers

import (
    //"inventory-backend/config"
    "inventory-backend/models"
    "net/http"
	"strconv"
    "github.com/gin-gonic/gin"
)

// GET /inventory - List all items
// GET /inventory - List all items with pagination, sorting, and filtering
// GET /inventory - List all items with pagination, sorting, and filtering
func GetItems(c *gin.Context) {
    var items []models.Item
    db := models.DB

    // Getting the query parameters
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
    sort := c.DefaultQuery("sort", "id")
    order := c.DefaultQuery("order", "asc")
    minStock, _ := strconv.Atoi(c.DefaultQuery("min_stock", "0"))

    // Filtering
    query := db.Where("stock >= ?", minStock)


    // Classification
    validSortFields := map[string]bool{"name": true, "stock": true, "price": true}
    if validSortFields[sort] {
        query = query.Order(sort + " " + order)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort field"})
        return
    }

    // Pagination
    offset := (page - 1) * size
    query = query.Offset(offset).Limit(size)

    // Run query
    if err := query.Debug().Find(&items).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return response
    c.JSON(http.StatusOK, gin.H{
        "data":  items,
        "page":  page,
        "size":  size,
        "total": len(items),
    })
}


// GET /inventory - Get an item by ID
func GetItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	if err := models.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// POST /inventory - Create a new item
func CreateItem(c *gin.Context) {
	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// PUT /inventory - Update an item by ID 
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	if err := models.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DELETE /inventory - Delete an item by ID 
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	if err := models.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := models.DB.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}