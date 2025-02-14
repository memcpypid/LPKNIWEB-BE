package controllers

import (
	"LPKNI/lpkni_project/config"
	"LPKNI/lpkni_project/models"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func isValidEmailAdmin(email_admin string) bool {

	re := regexp.MustCompile(`^[a-z0-9]+([.-_]?[a-z0-9]+)*@[a-z0-9]+([.-_]?[a-z0-9]+)*\.[a-z]{2,}$`)
	return re.MatchString(email_admin)
}

func RegisterAdmin(c *gin.Context) {
	var admin models.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !isValidEmailAdmin(admin.EmailAdmin) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format email tidak valid"})
		return
	}

	var existingAdmin models.Admin
	if err := config.DB.Where("email_admin = ?", admin.EmailAdmin).First(&existingAdmin).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email Admin sudah terdaftar, coba yang baru lagi"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password :) "})
		return
	}
	admin.Password = string(hashedPassword)

	if err := config.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin berhasil didaftarkan", "admin": admin})
}

func LoginAdmin(c *gin.Context) {
	var admin models.Admin
	var input struct {
		EmailAdmin string `json:"email_admin"`
		Password   string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("email_admin = ?", input.EmailAdmin).First(&admin).Error; err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Admin tidak ditemukan"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "admin": admin})
}

func GetAdminByID(c *gin.Context) {
	id := c.Param("id")
	var admin models.Admin

	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	if err := config.DB.First(&admin, intID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"admin": admin})
}

func GetAllAdmins(c *gin.Context) {
	var admins []models.Admin

	if err := config.DB.Find(&admins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data admin: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"admins": admins})
}

func DeleteAdmin(c *gin.Context) {
	id := c.Param("id")
	var admin models.Admin

	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	if err := config.DB.First(&admin, intID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus admin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin berhasil dihapus"})
}

func UploadSuratIjinAkses(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	// Ambil file dari form
	file, err := c.FormFile("surat_ijin_akses")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File tidak ditemukan"})
		return
	}

	// Tentukan lokasi penyimpanan file
	filePath := "uploads/admin/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	// Update data admin dengan path file
	var admin models.Admin
	if err := config.DB.First(&admin, intID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin tidak ditemukan"})
		return
	}

	admin.SuratIjinAkses = filePath
	if err := config.DB.Save(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data admin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Surat Ijin Akses berhasil diunggah", "file_path": filePath})
}
