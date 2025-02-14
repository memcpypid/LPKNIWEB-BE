package controllers

import (
	"LPKNI/lpkni_project/config"
	"LPKNI/lpkni_project/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllPengajuanSponsor(c *gin.Context) {
	var pengajuanSponsors []models.PengajuanSponsor
	if err := config.DB.Find(&pengajuanSponsors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pengajuanSponsors)
}

func GetPengajuanSponsorByID(c *gin.Context) {
	id := c.Param("id")
	var pengajuanSponsor models.PengajuanSponsor

	if err := config.DB.First(&pengajuanSponsor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengajuan sponsor tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, pengajuanSponsor)
}

func isValidEmailSponsor(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func CreatePengajuanSponsor(c *gin.Context) {
	var pengajuanSponsor models.PengajuanSponsor

	if err := c.ShouldBindJSON(&pengajuanSponsor); err != nil {
		fmt.Println("Error Binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Email diterima:", pengajuanSponsor.ApplicantEmail)

	if !isValidEmailSponsor(pengajuanSponsor.ApplicantEmail) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format email tidak valid"})
		return
	}

	if pengajuanSponsor.AmountRequested <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Jumlah sponsor harus lebih dari 0"})
		return
	}

	if pengajuanSponsor.StartDate != nil && pengajuanSponsor.EndDate != nil {
		if pengajuanSponsor.EndDate.Before(*pengajuanSponsor.StartDate) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tanggal akhir tidak boleh lebih awal dari tanggal mulai"})
			return
		}
	}

	if err := config.DB.Create(&pengajuanSponsor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Pengajuan sponsor berhasil dibuat",
		"data":    pengajuanSponsor,
	})
}
func UpdatePengajuanSponsor(c *gin.Context) {
	id := c.Param("id")
	var pengajuanSponsor models.PengajuanSponsor

	if err := config.DB.First(&pengajuanSponsor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengajuan sponsor tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&pengajuanSponsor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pengajuanSponsor.UpdatedAt = time.Now()

	if err := config.DB.Save(&pengajuanSponsor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pengajuanSponsor)
}

func DeletePengajuanSponsor(c *gin.Context) {
	id := c.Param("id")
	var pengajuanSponsor models.PengajuanSponsor

	if err := config.DB.First(&pengajuanSponsor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengajuan sponsor tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&pengajuanSponsor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengajuan sponsor berhasil dihapus"})
}

func UploadProposal(c *gin.Context) {

	id := c.Param("id")
	var pengajuanSponsor models.PengajuanSponsor

	if err := config.DB.First(&pengajuanSponsor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengajuan sponsor tidak ditemukan"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File tidak ditemukan"})
		return
	}

	allowedExtensions := map[string]bool{".pdf": true, ".doc": true, ".docx": true}
	ext := filepath.Ext(file.Filename)
	if !allowedExtensions[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak valid. Hanya PDF, DOC, DOCX"})
		return
	}

	if file.Size > 10<<20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ukuran file terlalu besar (maksimal 10MB)"})
		return
	}

	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	filename := fmt.Sprintf("proposal_%s%s", id, ext)
	filePath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	pengajuanSponsor.Attachment = filename
	if err := config.DB.Save(&pengajuanSponsor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "File proposal berhasil diunggah",
		"file_url": "/uploads/" + filename,
	})
}
