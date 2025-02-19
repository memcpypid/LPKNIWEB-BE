package controllers

import (
	"LPKNI/lpkniService/config"
	"LPKNI/lpkniService/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateBerita - Create a new Berita
func CreateBerita(c *gin.Context) {
	var berita models.Berita


	// Bind JSON data to the model
	if err := c.ShouldBindJSON(&berita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(berita.Kategori) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kategori Wajib Di isi"})
		return
	}
	// Set timestamps
	berita.CreatedAt = time.Now()
	berita.UpdatedAt = time.Now()

	// Save Berita to the database
	if err := config.DB.Create(&berita).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, berita)
}

// GetAllBerita - Get all Berita with media and categories
func GetAllBerita(c *gin.Context) {
	var beritaList []models.Berita

	if err := config.DB.Preload("Media").Preload("Kategori").Preload("Wilayah").Preload("Daerah").Find(&beritaList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, beritaList)
}

// GetBeritaByID - Get a single Berita by ID
func GetBeritaByID(c *gin.Context) {
	id := c.Param("id")
	var berita models.Berita

	if err := config.DB.Preload("Media").Preload("Kategori").First(&berita, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Berita not found"})
		return
	}

	c.JSON(http.StatusOK, berita)
}

// UpdateBerita - Update a Berita by ID
func UpdateBerita(c *gin.Context) {
	id := c.Param("id")
	var updatedBerita models.Berita

	// Bind JSON data to the model
	if err := c.ShouldBindJSON(&updatedBerita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find Berita by ID
	var berita models.Berita
	if err := config.DB.First(&berita, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Berita not found"})
		return
	}

	// Update fields
	berita.Judul = updatedBerita.Judul
	berita.Konten = updatedBerita.Konten
	berita.Status = updatedBerita.Status
	berita.Penulis = updatedBerita.Penulis
	berita.Tanggal = updatedBerita.Tanggal
	berita.UpdatedAt = time.Now()

	// Save updates
	if err := config.DB.Save(&berita).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Replace Media and Kategori
	if err := config.DB.Model(&berita).Association("Media").Replace(updatedBerita.Media); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&berita).Association("Kategori").Replace(updatedBerita.Kategori); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, berita)
}

// DeleteBerita - Delete a Berita by ID
func DeleteBerita(c *gin.Context) {
	id := c.Param("id")
	var berita models.Berita

	// Find Berita by ID
	if err := config.DB.First(&berita, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Berita not found"})
		return
	}

	// Delete related media (delete media where berita_id matches)
	if err := config.DB.Where("berita_id = ?", id).Delete(&models.MediaBerita{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Delete related categories (many-to-many relationship)
	if err := config.DB.Model(&berita).Association("Kategori").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Delete berita
	if err := config.DB.Delete(&berita).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berita deleted successfully"})
}






// CreateKategori - Create a new KategoriBerita
func CreateKategori(c *gin.Context) {
	var kategori models.KategoriBerita

	// Bind JSON data to the model
	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set timestamps
	kategori.CreatedAt = time.Now()
	kategori.UpdatedAt = time.Now()

	// Save Kategori to the database
	if err := config.DB.Create(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, kategori)
}

// GetAllKategori - Get all KategoriBerita
func GetAllKategori(c *gin.Context) {
	var kategoriList []models.KategoriBerita

	if err := config.DB.Find(&kategoriList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kategoriList)
}

// GetKategoriByID - Get a single KategoriBerita by ID
func GetKategoriByID(c *gin.Context) {
	id := c.Param("id")
	var kategori models.KategoriBerita

	if err := config.DB.First(&kategori, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori not found"})
		return
	}

	c.JSON(http.StatusOK, kategori)
}

// UpdateKategori - Update a KategoriBerita by ID
func UpdateKategori(c *gin.Context) {
	id := c.Param("id")
	var updatedKategori models.KategoriBerita

	// Bind JSON data to the model
	if err := c.ShouldBindJSON(&updatedKategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find Kategori by ID
	var kategori models.KategoriBerita
	if err := config.DB.First(&kategori, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori not found"})
		return
	}

	// Update fields
	kategori.Nama = updatedKategori.Nama
	kategori.UpdatedAt = time.Now()

	// Save updates
	if err := config.DB.Save(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kategori)
}

// DeleteKategori - Delete a KategoriBerita by ID
func DeleteKategori(c *gin.Context) {
	id := c.Param("id")
	var kategori models.KategoriBerita

	// Find Kategori by ID
	if err := config.DB.First(&kategori, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori not found"})
		return
	}

	// Delete kategori
	if err := config.DB.Delete(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori deleted successfully"})
}
