package main

import (
	"fmt"
	"log"
	"os"

	"edunews-backend/database"
	"edunews-backend/handlers"
	"edunews-backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func mustEnv(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}

func main() {
	// Load environment
	port := mustEnv("PORT", "8080")
	dbHost := mustEnv("DB_HOST", "localhost")
	dbPort := mustEnv("DB_PORT", "5432")
	dbName := mustEnv("DB_NAME", "edunews")
	dbUser := mustEnv("DB_USER", "postgres")
	dbPass := mustEnv("DB_PASS", "12345")

	// DSN PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	// Connect to DB
	database.Connect(dsn)

	// Auto migrate
	if err := database.DB.AutoMigrate(&models.Berita{}, &models.User{}, &models.Category{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	// Setup Gin
	r := gin.Default()
	r.Use(cors.Default())

	// Serve uploads
	r.Static("/uploads", "./public/uploads")

	// Pastikan folder uploads/berita ada
	uploadDir := "./public/uploads/berita"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Fatalf("Gagal membuat folder upload: %v", err)
	}

	// API group
	api := r.Group("/api")
	{
		// Berita CRUD
		api.GET("/berita", handlers.GetAllBerita)
		api.GET("/berita/:id", handlers.GetBeritaByID)
		api.POST("/berita", handlers.CreateBerita)
		api.PUT("/berita/:id", handlers.UpdateBerita)
		api.DELETE("/berita/:id", handlers.DeleteBerita)

		// Category CRUD
		api.GET("/categories", handlers.GetCategories)
		api.POST("/categories", handlers.CreateCategory)
		api.PUT("/categories/:id", handlers.UpdateCategory)
		api.DELETE("/categories/:id", handlers.DeleteCategory)

		// Users read-only
		api.GET("/users", handlers.GetUsers)

		// Auth
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
	}

	// Run server
	log.Printf("Server running at http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server run failed: %v", err)
	}
}
