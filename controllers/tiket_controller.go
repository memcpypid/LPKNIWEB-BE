package controllers

import (
	"LPKNI/lpkni_project/config"
	"LPKNI/lpkni_project/models"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllTiket(c *gin.Context) {
	var tiket []models.Tiket
	if err := config.DB.Find(&tiket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tiket)
}

func GetTiketByID(c *gin.Context) {
	id := c.Param("id")
	var tiket models.Tiket
	if err := config.DB.First(&tiket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tiket tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, tiket)
}

func CreateTiket(c *gin.Context) {
	var tiket models.Tiket
	if err := c.ShouldBindJSON(&tiket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(tiket.EmailPengguna) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format email tidak valid"})
		return
	}

	tiket.TanggalBuat = time.Now()
	tiket.TanggalDiperbarui = time.Now()

	if err := config.DB.Create(&tiket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tiket)
}

func UpdateTiket(c *gin.Context) {
	id := c.Param("id")
	var tiket models.Tiket

	if err := config.DB.First(&tiket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tiket tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&tiket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tiket.TanggalDiperbarui = time.Now()

	if err := config.DB.Save(&tiket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tiket)
}

func DeleteTiket(c *gin.Context) {
	id := c.Param("id")
	var tiket models.Tiket

	if err := config.DB.First(&tiket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tiket tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&tiket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tiket berhasil dihapus"})
}

func UpdateStatusTiket(c *gin.Context) {
	id := c.Param("id")
	var tiket models.Tiket

	if err := config.DB.First(&tiket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tiket tidak ditemukan"})
		return
	}

	status := c.PostForm("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status tidak boleh kosong"})
		return
	}

	if status != "open" && status != "in_progress" && status != "resolved" && status != "closed" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status tidak valid"})
		return
	}

	tiket.Status = status
	tiket.TanggalDiperbarui = time.Now()

	if err := config.DB.Save(&tiket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tiket)
}
