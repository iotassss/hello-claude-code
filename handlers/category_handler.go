package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hello-gin/database"
	"hello-gin/models"
)

// GetCategories retrieves all categories with task counts
func GetCategories(c *gin.Context) {
	var categories []models.Category

	if err := database.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	// Add task counts for each category
	for i := range categories {
		database.DB.Model(&models.Task{}).Where("category_id = ?", categories[i].ID).Count(&categories[i].TaskCount)
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// GetCategory retrieves a single category by ID
func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Add task count
	database.DB.Model(&models.Task{}).Where("category_id = ?", category.ID).Count(&category.TaskCount)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// CreateCategory creates a new category
func CreateCategory(c *gin.Context) {
	var req models.CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default color if not provided
	if req.Color == "" {
		req.Color = "#3498db"
	}

	category := models.Category{
		Name:        req.Name,
		Color:       req.Color,
		Description: req.Description,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": category})
}

// UpdateCategory updates an existing category
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	var req models.UpdateCategoryRequest

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields if provided
	if req.Name != nil {
		category.Name = *req.Name
	}
	if req.Color != nil {
		category.Color = *req.Color
	}
	if req.Description != nil {
		category.Description = *req.Description
	}

	if err := database.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteCategory deletes a category (only if no tasks are associated)
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Check if category has associated tasks
	var taskCount int64
	database.DB.Model(&models.Task{}).Where("category_id = ?", id).Count(&taskCount)

	if taskCount > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Cannot delete category with associated tasks"})
		return
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}