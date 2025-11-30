package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"edunews-backend/database"
	"edunews-backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllBerita(c *gin.Context) {
	var berita []models.Berita

	if err := database.DB.
		Preload("Category").
		Find(&berita).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil berita"})
		return
	}

	c.JSON(http.StatusOK, berita)
}

func GetBeritaByID(c *gin.Context) {
	id := c.Param("id")

	var berita models.Berita
	if err := database.DB.
		Preload("Category").
		First(&berita, id).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Berita tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, berita)
}

func CreateBerita(c *gin.Context) {
	title := c.PostForm("title")
	cid := c.PostForm("category_id")
	content := c.PostForm("content")
	excerpt := c.PostForm("excerpt")

	var coverFilename string
	file, err := c.FormFile("cover")
	if err == nil && file != nil {
		uploadRoot := "./public/uploads"
		outDir := filepath.Join(uploadRoot, "berita")
		_ = os.MkdirAll(outDir, 0755)

		coverFilename = fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		dst := filepath.Join(outDir, coverFilename)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(500, gin.H{"error": "Gagal menyimpan file"})
			return
		}
	}

	b := models.Berita{
		Title:      title,
		CategoryID: parseUint(cid),
		Content:    content,
		Excerpt:    excerpt,
		Cover:      coverFilename,
		CreatedAt:  time.Now(),
	}

	if err := database.DB.Create(&b).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, b)
}

func parseUint(s string) uint {
	v, _ := strconv.ParseUint(s, 10, 32)
	return uint(v)
}

// PUT /api/berita/:id  (supports multipart or json)
func UpdateBerita(c *gin.Context) {
	id := c.Param("id")

	var b models.Berita
	if err := database.DB.First(&b, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Berita tidak ditemukan"})
		return
	}

	// =============================
	// VALIDASI CATEGORY ID
	// =============================
	if v := c.PostForm("category_id"); v != "" {
		cid := parseUint(v)

		// Jika 0 → abaikan, jangan simpan karena 0 bukan kategori valid
		if cid != 0 {
			// Pastikan kategori benar-benar ada
			var cat models.Category
			if err := database.DB.First(&cat, cid).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Kategori tidak valid",
				})
				return
			}
			b.CategoryID = cid
		}
	}

	// FIELD LAIN
	if v := c.PostForm("title"); v != "" {
		b.Title = v
	}
	if v := c.PostForm("content"); v != "" {
		b.Content = v
	}
	if v := c.PostForm("excerpt"); v != "" {
		b.Excerpt = v
	}

	// COVER (UPLOAD)
	file, err := c.FormFile("cover")
	if err == nil && file != nil {
		uploadRoot := "./public/uploads"
		outDir := filepath.Join(uploadRoot, "berita")
		_ = os.MkdirAll(outDir, 0755)

		coverFilename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		dst := filepath.Join(outDir, coverFilename)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(500, gin.H{"error": "Gagal upload cover"})
			return
		}

		// hapus file lama
		if b.Cover != "" {
			oldFile := filepath.Join(outDir, b.Cover)
			_ = os.Remove(oldFile)
		}

		b.Cover = coverFilename
	}

	// SAVE
	if err := database.DB.Save(&b).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menyimpan perubahan",
			"detail": err.Error(),   // debug
		})
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
