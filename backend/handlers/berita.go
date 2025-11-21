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
	"golang.org/x/crypto/bcrypt"
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

// POST /api/register
func Register(c *gin.Context) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	user := models.User{
		Email:    payload.Email,
		Password: string(hash),
		Name:     payload.Name,
		Role:     "admin",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dibuat"})
}

// POST /api/login
func Login(c *gin.Context) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"role":  user.Role,
	})
}