package seed

import (
    "inventory-backend/config"
    "inventory-backend/models"
)

func LoadSeedData() {
    items := []models.Item{
        {Name: "Laptop", Stock: 5, Price: 1500.0},
        {Name: "Mouse", Stock: 10, Price: 25.0},
        {Name: "Keyboard", Stock: 7, Price: 50.0},
    }

    for _, item := range items {
        config.DB.Create(&item)
    }
}

