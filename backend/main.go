package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	// Connect DB
	database.Connect(dsn)

	// Auto migrate
	if err := database.DB.AutoMigrate(&models.Berita{}, &models.User{}, &models.Category{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	r := gin.Default()

	// ===============================
	// FIX CORS PALING AMAN
	// ===============================
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Static files
	r.Static("/uploads", "./public/uploads")

	// Ensure upload folder exists
	uploadDir := "./public/uploads/berita"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Fatalf("Gagal membuat folder upload: %v", err)
	}

	// ===============================
	// API ROUTES
	// ===============================
	api := r.Group("/api")
	{
		// AUTH
		api.POST("/auth/register", handlers.Register)
		api.POST("/auth/login", handlers.Login)

		// USERS
		api.GET("/users", handlers.GetUsers)          // GET semua user, tanpa token
		api.DELETE("/users/:id", handlers.DeleteUser) // DELETE user cepat, tanpa token
		api.GET("/users/:id", handlers.GetUserByID)
		api.PUT("/users/:id", handlers.UpdateUser)

		// BERITA CRUD
		api.GET("/berita", handlers.GetAllBerita)
		api.GET("/berita/:id", handlers.GetBeritaByID)
		api.POST("/berita", handlers.CreateBerita)
		api.PUT("/berita/:id", handlers.UpdateBerita)
		api.DELETE("/berita/:id", handlers.DeleteBerita)

		// 🟢 TAMBAHKAN INI - Route untuk hitung berita per kategori
		api.GET("/berita/count-by-category/:id", handlers.GetNewsCountByCategory)

		// CATEGORY CRUD
		api.GET("/categories", handlers.GetCategories)
		api.GET("/categories/:id", handlers.GetCategory)
		api.POST("/categories", handlers.CreateCategory)
		api.PUT("/categories/:id", handlers.UpdateCategory)
		api.DELETE("/categories/:id", handlers.DeleteCategory)
	}

	// Run Server
	log.Printf("Server berjalan di http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server run failed: %v", err)
	}
}
