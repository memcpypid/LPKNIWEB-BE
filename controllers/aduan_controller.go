package controllers

import (
	"LPKNI/lpkni_project/config"
	"LPKNI/lpkni_project/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAduan(c *gin.Context) {
	var aduan models.Aduan

	if err := c.ShouldBindJSON(&aduan); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&aduan).Error; err != nil {
		log.Printf("Error saving Aduan: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Aduan berhasil dibuat", "aduan": aduan})
}

func GetAduanByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var aduan models.Aduan
	if err := config.DB.First(&aduan, id).Error; err != nil {
		log.Printf("Error retrieving Aduan with ID %v: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Aduan not found"})
		return
	}

	c.JSON(http.StatusOK, aduan)
}

func UpdateAduan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var aduan models.Aduan
	if err := config.DB.First(&aduan, id).Error; err != nil {
		log.Printf("Error retrieving Aduan with ID %v: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Aduan not found"})
		return
	}

	if err := c.ShouldBindJSON(&aduan); err != nil {
		log.Printf("Error binding JSON for update: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&aduan).Error; err != nil {
		log.Printf("Error saving updated Aduan: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aduan)
}

func DeleteAduan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var aduan models.Aduan
	if err := config.DB.First(&aduan, id).Error; err != nil {
		log.Printf("Error retrieving Aduan with ID %v: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Aduan not found"})
		return
	}

	if err := config.DB.Delete(&aduan).Error; err != nil {
		log.Printf("Error deleting Aduan with ID %v: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aduan deleted successfully"})
}

func GetAllAduan(c *gin.Context) {
	var aduans []models.Aduan

	if err := config.DB.Find(&aduans).Error; err != nil {
		log.Printf("Error retrieving all Aduan: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aduans)
}

func UploadDokumenAduan(c *gin.Context) {
	aduanID := c.Param("aduan_id")
	if aduanID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Aduan ID tidak valid"})
		return
	}

	var aduan models.Aduan

	if err := config.DB.First(&aduan, aduanID).Error; err != nil {
		log.Printf("Aduan dengan ID %v tidak ditemukan", aduanID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Aduan tidak ditemukan"})
		return
	}

	files := c.Request.MultipartForm.File["dokumen"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dokumen tidak ditemukan"})
		return
	}

	var aduanDokumen []models.AduanDokumen
	for _, file := range files {
		filePath := "./uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			log.Printf("Error saving file %v: %v", file.Filename, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		aduanDokumen = append(aduanDokumen, models.AduanDokumen{
			AduanID:  aduan.ID,
			FileURL:  filePath,
			FileName: file.Filename,
		})
	}

	for _, doc := range aduanDokumen {
		if err := config.DB.Create(&doc).Error; err != nil {
			log.Printf("Error saving document %v: %v", doc.FileName, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan dokumen", "detail": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dokumen berhasil di-upload", "aduan_id": aduan.ID})
}
