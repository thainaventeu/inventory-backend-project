package routes

import (
    "inventory-backend/controllers"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.GET("/inventory", controllers.GetItems)
    router.POST("/inventory", controllers.CreateItem)
}
