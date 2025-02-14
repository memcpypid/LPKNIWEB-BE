package controllers

import (
	"LPKNI/lpkni_project/config"
	"LPKNI/lpkni_project/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func PendaftaranAnggota(c *gin.Context) {
	var anggota models.Anggota
	if err := c.ShouldBindJSON(&anggota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidEmail(anggota.EmailAnggota) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format email tidak valid"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(anggota.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}
	anggota.Password = string(hashedPassword)

	if err := config.DB.Create(&anggota).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Pendaftaran berhasil",
		"anggota": anggota,
	})
}

func GetAnggotaByID(c *gin.Context) {
	id := c.Param("id")
	var anggota models.Anggota

	if err := config.DB.Preload("Keuangans").First(&anggota, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anggota tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"anggota":  anggota,
		"keuangan": anggota.Keuangans,
	})
}

func GetAllAnggota(c *gin.Context) {
	var anggota []models.Anggota

	if err := config.DB.Preload("Keuangans").Find(&anggota).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data anggota"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"anggota": anggota,
	})
}

func UpdateAnggota(c *gin.Context) {
	id := c.Param("id")
	var anggota models.Anggota

	if err := config.DB.First(&anggota, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anggota tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&anggota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&anggota).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data anggota"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data anggota berhasil diperbarui",
		"anggota": anggota,
	})
}

func DeleteAnggota(c *gin.Context) {
	id := c.Param("id")
	var anggota models.Anggota

	if err := config.DB.First(&anggota, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anggota tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&anggota).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data anggota"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data anggota berhasil dihapus",
	})
}

func LoginAnggota(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var anggota models.Anggota

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("email_anggota = ?", input.Email).First(&anggota).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(anggota.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func UploadSuratIjinAksesUs(c *gin.Context) {

	file, _ := c.FormFile("surat_ijin_akses")
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File surat izin akses tidak ditemukan"})
		return
	}

	dir := "./uploads/surat_ijin_akses"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat direktori"})
		return
	}

	extension := filepath.Ext(file.Filename)
	uniqueFileName := fmt.Sprintf("%s%s", uuid.New().String(), extension)
	filePath := filepath.Join(dir, uniqueFileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	var anggota models.Anggota
	anggotaID := c.Param("anggota_id")
	if err := config.DB.First(&anggota, anggotaID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anggota tidak ditemukan"})
		return
	}

	anggota.SuratIjinAkses = filePath
	if err := config.DB.Save(&anggota).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data surat izin akses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Surat izin akses berhasil di-upload",
		"file_path": filePath,
	})
}
