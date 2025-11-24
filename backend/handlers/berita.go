package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"edunews-backend/database"
	"edunews-backend/models"

	"github.com/gin-gonic/gin"
)

// GET /api/berita
func GetAllBerita(c *gin.Context) {
	var list []models.Berita
	if err := database.DB.Order("created_at desc").Find(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

// GET /api/berita/:id
func GetBeritaByID(c *gin.Context) {
	id := c.Param("id")
	var b models.Berita
	if err := database.DB.First(&b, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Berita tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, b)
}

// POST /api/berita  (multipart/form-data)
func CreateBerita(c *gin.Context) {
	title := c.PostForm("title")
	category := c.PostForm("category")
	content := c.PostForm("content")
	excerpt := c.PostForm("excerpt")

	var coverFilename string
	file, err := c.FormFile("cover")
	if err == nil && file != nil {
		// 🔥 FIX: Ubah ke ./public/uploads agar sinkron dengan static di main.go
		uploadRoot := "./public/uploads" // Dari "./uploads"
		outDir := filepath.Join(uploadRoot, "berita")
		_ = os.MkdirAll(outDir, 0755)

		coverFilename = fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		dst := filepath.Join(outDir, coverFilename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
			return
		}
	}

	b := models.Berita{
		Title:     title,
		Category:  category,
		Content:   content,
		Excerpt:   excerpt,
		Cover:     coverFilename,
		CreatedAt: time.Now(),
	}

	if err := database.DB.Create(&b).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan berita"})
		return
	}

	c.JSON(http.StatusOK, b)
}

// PUT /api/berita/:id  (supports multipart or json)
func UpdateBerita(c *gin.Context) {
	id := c.Param("id")
	var b models.Berita
	if err := database.DB.First(&b, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Berita tidak ditemukan"})
		return
	}

	if v := c.PostForm("title"); v != "" {
		b.Title = v
	}
	if v := c.PostForm("category"); v != "" {
		b.Category = v
	}
	if v := c.PostForm("content"); v != "" {
		b.Content = v
	}
	if v := c.PostForm("excerpt"); v != "" {
		b.Excerpt = v
	}

	file, err := c.FormFile("cover")
	if err == nil && file != nil {
		// 🔥 FIX: Ubah ke ./public/uploads
		uploadRoot := "./public/uploads" // Dari "./uploads"
		outDir := filepath.Join(uploadRoot, "berita")
		_ = os.MkdirAll(outDir, 0755)

		coverFilename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		dst := filepath.Join(outDir, coverFilename)
		if err := c.SaveUploadedFile(file, dst); err == nil {
			if b.Cover != "" {
				_ = os.Remove(filepath.Join(outDir, b.Cover)) // hapus lama
			}
			b.Cover = coverFilename
		}
	}

	if err := database.DB.Save(&b).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan perubahan"})
		return
	}

	c.JSON(http.StatusOK, b)
}

// DELETE /api/berita/:id
func DeleteBerita(c *gin.Context) {
	id := c.Param("id")
	var b models.Berita
	if err := database.DB.First(&b, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Berita tidak ditemukan"})
		return
	}

	// delete DB record
	if err := database.DB.Delete(&b).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus berita"})
		return
	}

	// delete file
	if b.Cover != "" {
		// 🔥 FIX: Ubah ke ./public/uploads
		uploadRoot := "./public/uploads" // Dari "./uploads"
		_ = os.Remove(filepath.Join(uploadRoot, "berita", b.Cover))
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berita dihapus"})
}
