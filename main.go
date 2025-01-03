package main

import (
	"inventory-backend/config"
	"inventory-backend/routes"
	"inventory-backend/seed"
    "inventory-backend/models"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	config.ConnectDatabase()
    
    // Assign the database connection to the models package
    models.DB = config.DB

	// Load seed data
	seed.LoadSeedData()

	// Initialize the Gin server
	r := gin.Default()

	// Configure rate limiting
	limiter := tollbooth.NewLimiter(1, nil) // 1 request per second
	limiter.SetBurst(5)                    // Burst capacity of 5 requests

	// Apply rate limiting middleware
	r.Use(tollbooth_gin.LimitHandler(limiter))

	// Configure routes
	routes.SetupRoutes(r)

	// Run the server
	r.Run(":8080")
}