package handlers

import (
	"net/http"

	"edunews-backend/database"
	"edunews-backend/models"

	"github.com/gin-gonic/gin"
)

// GET /users
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
