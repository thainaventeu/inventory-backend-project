package config

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"inventory-backend/models"
)




var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres_user password=yourpassword dbname=inventory port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

	// Auto-migrate models
    err = db.AutoMigrate(&models.Item{})
    if err != nil {
        panic("Failed to migrate database: " + err.Error())
    }

	DB = db
    fmt.Println("Database connected successfully!")
    
}
