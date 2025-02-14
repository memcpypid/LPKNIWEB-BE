package controllers

import (
	"LPKNI/lpkni_project/config"
	"LPKNI/lpkni_project/models"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllPendaftaran(c *gin.Context) {
	var pendaftaran []models.Pendaftaran
	if err := config.DB.Find(&pendaftaran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pendaftaran)
}

func GetPendaftaranByID(c *gin.Context) {
	id := c.Param("id")
	var pendaftaran models.Pendaftaran
	if err := config.DB.First(&pendaftaran, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pendaftaran tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, pendaftaran)
}

func CreatePendaftaran(c *gin.Context) {
	var pendaftaran models.Pendaftaran

	
	if err := c.ShouldBind(&pendaftaran); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	fileFoto3x4, _ := c.FormFile("foto3x4")
	if fileFoto3x4 != nil {
		ext := filepath.Ext(fileFoto3x4.Filename)
		newFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
		newFileName = strings.ReplaceAll(newFileName, " ", "_")
		filePath := filepath.Join("uploads", newFileName)
		if err := c.SaveUploadedFile(fileFoto3x4, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto 3x4"})
			return
		}
		pendaftaran.Foto3x4 = "/uploads/" + newFileName
	}

	fileFotoKtp, _ := c.FormFile("fotoKtp")
	if fileFotoKtp != nil {
		ext := filepath.Ext(fileFotoKtp.Filename)
		newFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
		newFileName = strings.ReplaceAll(newFileName, " ", "_")
		filePath := filepath.Join("uploads", newFileName)
		if err := c.SaveUploadedFile(fileFotoKtp, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto KTP"})
			return
		}
		pendaftaran.FotoKtp = "/uploads/" + newFileName
	}

	
	if err := config.DB.Create(&pendaftaran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan pendaftaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pendaftaran berhasil ditambahkan", "data": pendaftaran})
}


func UpdatePendaftaran(c *gin.Context) {
	id := c.Param("id")
	var pendaftaran models.Pendaftaran

	
	if err := config.DB.First(&pendaftaran, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pendaftaran tidak ditemukan"})
		return
	}

	
	if err := c.ShouldBind(&pendaftaran); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	foto3x4, _ := c.FormFile("foto3x4")
	fotoKtp, _ := c.FormFile("fotoKtp")
	if foto3x4 != nil {
		pendaftaran.Foto3x4 = foto3x4.Filename
		if err := c.SaveUploadedFile(foto3x4, "/uploads/"+foto3x4.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan Foto 3x4"})
			return
		}
	}
	if fotoKtp != nil {
		pendaftaran.FotoKtp = fotoKtp.Filename
		if err := c.SaveUploadedFile(fotoKtp, "/uploads/"+fotoKtp.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan Foto KTP"})
			return
		}
	}

	
	if err := config.DB.Save(&pendaftaran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pendaftaran)
}

func DeletePendaftaran(c *gin.Context) {
	id := c.Param("id")
	var pendaftaran models.Pendaftaran

	
	if err := config.DB.First(&pendaftaran, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pendaftaran tidak ditemukan"})
		return
	}

	
	if err := config.DB.Delete(&pendaftaran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus pendaftaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pendaftaran berhasil dihapus"})
}

func UploadFoto(c *gin.Context) {
	id := c.Param("id")
	var pendaftaran models.Pendaftaran

	
	if err := config.DB.First(&pendaftaran, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pendaftaran tidak ditemukan"})
		return
	}

	
	foto3x4, _ := c.FormFile("foto3x4")
	if foto3x4 != nil {
		if err := c.SaveUploadedFile(foto3x4, "/uploads/"+foto3x4.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupload Foto 3x4"})
			return
		}
		pendaftaran.Foto3x4 = foto3x4.Filename
	}

	
	fotoKtp, _ := c.FormFile("fotoKtp")
	if fotoKtp != nil {
		if err := c.SaveUploadedFile(fotoKtp, "/uploads/"+fotoKtp.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupload Foto KTP"})
			return
		}
		pendaftaran.FotoKtp = fotoKtp.Filename
	}

	
	if err := config.DB.Save(&pendaftaran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui pendaftaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Foto berhasil di-upload", "pendaftaran": pendaftaran})
}
