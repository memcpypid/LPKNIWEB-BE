package controllers

import (
	"LPKNI/lpkniService/config"
	"LPKNI/lpkniService/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all pengaduans
func GetAllPengaduan(c *gin.Context) {
	var pengaduans []models.PengaduanKonsumen
	if err := config.DB.Preload("Media").Find(&pengaduans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch data"})
		return
	}
	c.JSON(http.StatusOK, pengaduans)
}

// Get one pengaduan by ID
func GetPengaduanById(c *gin.Context) {
	id := c.Param("id")
	var pengaduan models.PengaduanKonsumen
	if err := config.DB.Preload("Media").First(&pengaduan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengaduan not found"})
		return
	}
	c.JSON(http.StatusOK, pengaduan)
}

// Create a new pengaduan
func CreatePengaduan(c *gin.Context) {
	var pengaduan models.PengaduanKonsumen
	if err := c.ShouldBindJSON(&pengaduan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Create(&pengaduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create pengaduan"})
		return
	}
	c.JSON(http.StatusCreated, pengaduan)
}

// Update a pengaduan
func UpdatePengaduan(c *gin.Context) {
	id := c.Param("id")
	var pengaduan models.PengaduanKonsumen
	if err := config.DB.First(&pengaduan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengaduan not found"})
		return
	}
	if err := c.ShouldBindJSON(&pengaduan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Save(&pengaduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update pengaduan"})
		return
	}
	c.JSON(http.StatusOK, pengaduan)
}

// Delete a pengaduan
func DeletePengaduan(c *gin.Context) {
	id := c.Param("id")
	var pengaduan models.PengaduanKonsumen
	if err := config.DB.First(&pengaduan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengaduan not found"})
		return
	}
	if err := config.DB.Delete(&pengaduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete pengaduan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pengaduan deleted successfully"})
}

// CRUD for models.MediaPengaduan

// Get media by pengaduan ID
func GetMediaByPengaduan(c *gin.Context) {
	pengaduanID := c.Param("pengaduan_id")
	var media []models.MediaPengaduan
	if err := config.DB.Where("pengaduan_id = ?", pengaduanID).Find(&media).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch media"})
		return
	}
	c.JSON(http.StatusOK, media)
}

// Create media
func CreateMedia(c *gin.Context) {
	var media models.MediaPengaduan
	if err := c.ShouldBindJSON(&media); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Create(&media).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create media"})
		return
	}
	c.JSON(http.StatusCreated, media)
}

// Update media
func UpdateMedia(c *gin.Context) {
	id := c.Param("id")
	var media models.MediaPengaduan
	if err := config.DB.First(&media, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}
	if err := c.ShouldBindJSON(&media); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Save(&media).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update media"})
		return
	}
	c.JSON(http.StatusOK, media)
}

// Delete media
func DeleteMedia(c *gin.Context) {
	id := c.Param("id")
	var media models.MediaPengaduan
	if err := config.DB.First(&media, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}
	if err := config.DB.Delete(&media).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete media"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Media deleted successfully"})
}
