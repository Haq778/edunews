package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}
	// optional ping
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("database ping error: %v", err)
	}

	// ensure upload path exists (relative to backend)
	uploadPath := os.Getenv("UPLOAD_PATH")
	if uploadPath == "" {
		uploadPath = "../public/uploads"
	}
	if err := os.MkdirAll(fmt.Sprintf("%s/berita", uploadPath), 0755); err != nil {
		log.Printf("warning: cannot create upload dir: %v", err)
	}

	log.Println("Database connected successfully")
}
