package controllers

import (
	"LPKNI/lpkniService/config"
	"LPKNI/lpkniService/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateDataUser untuk membuat data pengguna baru
func CreateDataUser(c *gin.Context) {
	var dataUser models.DataUser
	if err := c.ShouldBindJSON(&dataUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan data pengguna ke database
	if err := config.DB.Create(&dataUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DataUser creation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "DataUser created successfully", "dataUser": dataUser})
}

// GetDataUser untuk mendapatkan detail data pengguna berdasarkan ID
func GetDataUserById(c *gin.Context) {
	dataUserID := c.Param("id")
	var dataUser models.DataUser

	if err := config.DB.First(&dataUser, dataUserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "DataUser not found"})
		return
	}

	c.JSON(http.StatusOK, dataUser)
}

// UpdateDataUser untuk memperbarui data pengguna
func UpdateDataUser(c *gin.Context) {
	dataUserID := c.Param("id")
	var dataUser models.DataUser

	// Bind data dari request body
	if err := c.ShouldBindJSON(&dataUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mencari DataUser berdasarkan ID dan mengupdate
	if err := config.DB.Model(&dataUser).Where("id = ?", dataUserID).Updates(dataUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update DataUser"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "DataUser updated successfully"})
}

// DeleteDataUser untuk menghapus data pengguna berdasarkan ID
func DeleteDataUser(c *gin.Context) {
	dataUserID := c.Param("id")
	var dataUser models.DataUser

	if err := config.DB.Delete(&dataUser, dataUserID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete DataUser"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "DataUser deleted successfully"})
}
