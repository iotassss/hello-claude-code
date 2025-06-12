package database

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hello-gin/models"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")
}

func Migrate() {
	err := DB.AutoMigrate(&models.Category{}, &models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully")
}

func SeedData() {
	// Check if categories already exist
	var count int64
	DB.Model(&models.Category{}).Count(&count)
	if count > 0 {
		log.Println("Categories already seeded")
		return
	}

	// Seed default categories
	categories := []models.Category{
		{Name: "Work", Color: "#e74c3c", Description: "Work-related tasks"},
		{Name: "Personal", Color: "#2ecc71", Description: "Personal tasks and goals"},
		{Name: "Shopping", Color: "#f39c12", Description: "Shopping lists and purchases"},
		{Name: "Health", Color: "#9b59b6", Description: "Health and fitness tasks"},
		{Name: "Finance", Color: "#1abc9c", Description: "Financial and budget tasks"},
		{Name: "Education", Color: "#34495e", Description: "Learning and educational tasks"},
	}

	for _, category := range categories {
		result := DB.Create(&category)
		if result.Error != nil {
			log.Printf("Failed to seed category %s: %v", category.Name, result.Error)
		}
	}

	log.Println("Default categories seeded successfully")
}