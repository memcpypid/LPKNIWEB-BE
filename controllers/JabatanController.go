package controllers

import (
	"LPKNI/lpkniService/config"
	"LPKNI/lpkniService/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all models.JabatanStruktural
func GetAllJabatan(c *gin.Context) {
	var jabatans []models.JabatanStruktural
	if err := config.DB.Find(&jabatans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch jabatan data"})
		return
	}
	c.JSON(http.StatusOK, jabatans)
}

// Get a single models.JabatanStruktural by ID
func GetJabatanById(c *gin.Context) {
	id := c.Param("id")
	var jabatan models.JabatanStruktural
	if err := config.DB.First(&jabatan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jabatan not found"})
		return
	}
	c.JSON(http.StatusOK, jabatan)
}

// Create a new models.JabatanStruktural
func CreateJabatan(c *gin.Context) {
	var jabatan models.JabatanStruktural
	if err := c.ShouldBindJSON(&jabatan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Create(&jabatan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create jabatan"})
		return
	}
	c.JSON(http.StatusCreated, jabatan)
}

// Update an existing models.JabatanStruktural by ID
func UpdateJabatan(c *gin.Context) {
	id := c.Param("id")
	var jabatan models.JabatanStruktural
	if err := config.DB.First(&jabatan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jabatan not found"})
		return
	}

	// Bind updated data to the model
	if err := c.ShouldBindJSON(&jabatan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Save(&jabatan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update jabatan"})
		return
	}
	c.JSON(http.StatusOK, jabatan)
}

// Delete a models.JabatanStruktural by ID
func DeleteJabatan(c *gin.Context) {
	id := c.Param("id")
	var jabatan models.JabatanStruktural
	if err := config.DB.First(&jabatan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jabatan not found"})
		return
	}

	if err := config.DB.Delete(&jabatan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete jabatan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Jabatan deleted successfully"})
}

// Count how many members are using a JabatanStruktural
func CountJabatanMembers(c *gin.Context) {
	id := c.Param("id")
	var jabatan models.JabatanStruktural
	if err := config.DB.First(&jabatan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jabatan not found"})
		return
	}

	// Count how many records in the `PengaduanKonsumen` table are using this JabatanStruktural
	var count int64
	if err := config.DB.Model(&models.DataAnggota{}).Where("jabatan_struktural_id = ?", jabatan.ID).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to count members for this jabatan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jabatan_id": jabatan.ID, "member_count": count})
}
