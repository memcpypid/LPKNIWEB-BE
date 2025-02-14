package controllers

import (
	"net/http"
	"time"

	"LPKNI/lpkni_project/config"
	"LPKNI/lpkni_project/models"

	"github.com/gin-gonic/gin"
)

func CreatePiutang(c *gin.Context) {
	var piutang models.Piutang
	if err := c.ShouldBindJSON(&piutang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	piutang.TanggalPengaduan = time.Now()

	if err := config.DB.Create(&piutang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, piutang)
}

func GetPiutangByID(c *gin.Context) {
	id := c.Param("id")
	var piutang models.Piutang

	if err := config.DB.First(&piutang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengaduan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, piutang)
}

func GetAllPiutangs(c *gin.Context) {
	var piutangs []models.Piutang

	if err := config.DB.Find(&piutangs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, piutangs)
}

func UpdatePiutangStatus(c *gin.Context) {
	id := c.Param("id")
	var piutang models.Piutang

	if err := config.DB.First(&piutang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengaduan tidak ditemukan"})
		return
	}

	var statusUpdate struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	piutang.Status = statusUpdate.Status
	if err := config.DB.Save(&piutang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, piutang)
}

func DeletePiutang(c *gin.Context) {
	id := c.Param("id")
	var piutang models.Piutang

	if err := config.DB.First(&piutang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengaduan tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&piutang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengaduan berhasil dihapus"})
}
