package main

import (
	"github.com/gin-gonic/gin"
	"hello-gin/database"
	"hello-gin/handlers"
)

func main() {
	// Initialize database
	database.Connect()
	database.Migrate()
	database.SeedData()

	// Initialize Gin router
	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Serve static files
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// Routes
	api := r.Group("/api/v1")
	{
		// Task routes
		api.GET("/tasks", handlers.GetTasks)
		api.GET("/tasks/stats", handlers.GetTaskStats)
		api.GET("/tasks/:id", handlers.GetTask)
		api.POST("/tasks", handlers.CreateTask)
		api.PUT("/tasks/:id", handlers.UpdateTask)
		api.DELETE("/tasks/:id", handlers.DeleteTask)
		api.PATCH("/tasks/:id/toggle", handlers.ToggleTaskStatus)

		// Category routes
		api.GET("/categories", handlers.GetCategories)
		api.GET("/categories/:id", handlers.GetCategory)
		api.POST("/categories", handlers.CreateCategory)
		api.PUT("/categories/:id", handlers.UpdateCategory)
		api.DELETE("/categories/:id", handlers.DeleteCategory)
	}

	// Frontend route
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "TODO App",
		})
	})

	r.Run(":8080")
}