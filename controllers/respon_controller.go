package controllers

import (
	"LPKNI/lpkni_project/config"
	"LPKNI/lpkni_project/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllRespon(c *gin.Context) {
	var respon []models.Respon
	if err := config.DB.Preload("Tiket").Find(&respon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, respon)
}

func GetResponByID(c *gin.Context) {
	id := c.Param("id")
	var respon models.Respon
	if err := config.DB.Preload("Tiket").First(&respon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Respon tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, respon)
}

func CreateRespon(c *gin.Context) {
	var respon models.Respon
	if err := c.ShouldBindJSON(&respon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	respon.TanggalRespon = time.Now()

	if err := config.DB.Create(&respon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, respon)
}

func UpdateRespon(c *gin.Context) {
	id := c.Param("id")
	var respon models.Respon

	if err := config.DB.First(&respon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Respon tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&respon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	respon.TanggalRespon = time.Now()

	if err := config.DB.Save(&respon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, respon)
}

func DeleteRespon(c *gin.Context) {
	id := c.Param("id")
	var respon models.Respon

	if err := config.DB.First(&respon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Respon tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&respon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Respon berhasil dihapus"})
}

func GetResponByTiketID(c *gin.Context) {
	tiketIDStr := c.Param("tiket_id")
	tiketID, err := strconv.Atoi(tiketIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Tiket tidak valid"})
		return
	}

	var respon []models.Respon
	if err := config.DB.Where("tiket_id = ?", tiketID).Preload("Tiket").Find(&respon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(respon) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Tidak ada respon untuk tiket ini"})
		return
	}

	c.JSON(http.StatusOK, respon)
}
