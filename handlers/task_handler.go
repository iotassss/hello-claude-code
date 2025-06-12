package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hello-gin/database"
	"hello-gin/models"
)

var validate = validator.New()

// GetTasks retrieves all tasks with optional filtering
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	query := database.DB.Preload("Category")

	// Filter by status
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// Filter by priority
	if priority := c.Query("priority"); priority != "" {
		query = query.Where("priority = ?", priority)
	}

	// Filter by category
	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	// Search by title or description
	if search := c.Query("search"); search != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Sort by
	sortBy := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")
	query = query.Order(sortBy + " " + order)

	if err := query.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks, "count": len(tasks)})
}

// GetTask retrieves a single task by ID
func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := database.DB.Preload("Category").First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// CreateTask creates a new task
func CreateTask(c *gin.Context) {
	var req models.CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default values
	if req.Priority == "" {
		req.Priority = "medium"
	}

	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		Status:      "pending",
		DueDate:     req.DueDate,
		CategoryID:  req.CategoryID,
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	// Load the category for the response
	database.DB.Preload("Category").First(&task, task.ID)

	c.JSON(http.StatusCreated, gin.H{"data": task})
}

// UpdateTask updates an existing task
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	var req models.UpdateTaskRequest

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
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
	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Status != nil {
		task.Status = *req.Status
	}
	if req.Priority != nil {
		task.Priority = *req.Priority
	}
	if req.DueDate != nil {
		task.DueDate = req.DueDate
	}
	if req.CategoryID != nil {
		task.CategoryID = req.CategoryID
	}

	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	// Load the category for the response
	database.DB.Preload("Category").First(&task, task.ID)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DeleteTask deletes a task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// GetTaskStats returns task statistics
func GetTaskStats(c *gin.Context) {
	var stats struct {
		Total       int64 `json:"total"`
		Pending     int64 `json:"pending"`
		InProgress  int64 `json:"in_progress"`
		Completed   int64 `json:"completed"`
		Overdue     int64 `json:"overdue"`
	}

	database.DB.Model(&models.Task{}).Count(&stats.Total)
	database.DB.Model(&models.Task{}).Where("status = ?", "pending").Count(&stats.Pending)
	database.DB.Model(&models.Task{}).Where("status = ?", "in_progress").Count(&stats.InProgress)
	database.DB.Model(&models.Task{}).Where("status = ?", "completed").Count(&stats.Completed)
	database.DB.Model(&models.Task{}).Where("due_date < ? AND status != ?", time.Now(), "completed").Count(&stats.Overdue)

	c.JSON(http.StatusOK, gin.H{"data": stats})
}

// ToggleTaskStatus toggles task status between pending/in_progress and completed
func ToggleTaskStatus(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if task.Status == "completed" {
		task.Status = "pending"
	} else {
		task.Status = "completed"
	}

	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task status"})
		return
	}

	// Load the category for the response
	database.DB.Preload("Category").First(&task, task.ID)

	c.JSON(http.StatusOK, gin.H{"data": task})
}