package seed

import (
    // "inventory-backend/config"
    "inventory-backend/models"
    "log"
)

func LoadSeedData() {
	// List of items to be added
	items := []models.Item{
		{Name: "Laptop", Stock: 5, Price: 1500},
		{Name: "Mouse", Stock: 10, Price: 25},
		{Name: "Keyboard", Stock: 7, Price: 50},
	}

	for _, item := range items {
		var existingItem models.Item
		// Checks if an item with the same name already exists
		if err := models.DB.Where("name = ?", item.Name).First(&existingItem).Error; err != nil {
			if err.Error() == "record not found" {
				// Inserts the item only if it does not exist
				if err := models.DB.Create(&item).Error; err != nil {
					log.Printf("Error inserting item %s: %v", item.Name, err)
				}
			} else {
				log.Printf("Error checking item %s: %v", item.Name, err)
			}
		}
	}
}