package routes

import (
    "inventory-backend/controllers"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Inventory route configuration
	r.GET("/inventory", controllers.GetItems)
	r.GET("/inventory/:id", controllers.GetItem)
	r.POST("/inventory", controllers.CreateItem)
	r.PUT("/inventory/:id", controllers.UpdateItem)
	r.DELETE("/inventory/:id", controllers.DeleteItem)
}
