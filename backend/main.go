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

	// DSN for PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	// Connect to DB
	database.Connect(dsn)

	// Auto migrate models
	if err := database.DB.AutoMigrate(&models.Berita{}, &models.User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	// Setup Gin
	r := gin.Default()
	r.Use(cors.Default())

	// Serve static uploads folder (Solusi A)
	// Bisa diakses via browser:
	// http://localhost:8080/uploads/berita/filename.jpg
	r.Static("/uploads", "./public/uploads")

	api := r.Group("/api")
	{
		// Berita CRUD
		api.GET("/berita", handlers.GetAllBerita)
		api.GET("/berita/:id", handlers.GetBeritaByID)
		api.POST("/berita", handlers.CreateBerita)
		api.PUT("/berita/:id", handlers.UpdateBerita)
		api.DELETE("/berita/:id", handlers.DeleteBerita)

		// User register + login
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
	}

	// Pastikan folder uploads/berita ada
	uploadDir := "./uploads/berita"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Fatalf("Gagal membuat folder upload: %v", err)
	}

	// Run server
	log.Printf("Server running at http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server run failed: %v", err)
	}
}
