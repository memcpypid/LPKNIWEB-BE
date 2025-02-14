package controllers

import (
	"LPKNI/lpkni_project/config"
	"LPKNI/lpkni_project/models"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllKeuangan(c *gin.Context) {
	var keuangan []models.Keuangan
	if err := config.DB.Preload("Anggota").Find(&keuangan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := make([]gin.H, 0, len(keuangan))

	for _, k := range keuangan {
		response = append(response, gin.H{
			"id":         k.ID,
			"anggota_id": k.AnggotaID,
			"anggota": gin.H{
				"anggota_id":   k.Anggota.AnggotaID,
				"nama_anggota": k.Anggota.NamaAnggota,
				"no_hp":        k.Anggota.NoHp,
			},
			"nominal":         k.Nominal,
			"tipe_transaksi":  k.TipeTransaksi,
			"deskripsi":       k.Deskripsi,
			"bukti_transaksi": k.BuktiTransaksi,
			"created_at":      k.CreatedAt,
			"updated_at":      k.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}
func GetKeuanganByID(c *gin.Context) {
	id := c.Param("id")
	var keuangan models.Keuangan
	if err := config.DB.Preload("Anggota").First(&keuangan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi tidak ditemukan"})
		return
	}
	response := gin.H{
		"id":         keuangan.ID,
		"anggota_id": keuangan.AnggotaID,
		"anggota": gin.H{
			"anggota_id":   keuangan.Anggota.AnggotaID,
			"nama_anggota": keuangan.Anggota.NamaAnggota,
			"no_hp":        keuangan.Anggota.NoHp,
		},
		"nominal":         keuangan.Nominal,
		"tipe_transaksi":  keuangan.TipeTransaksi,
		"deskripsi":       keuangan.Deskripsi,
		"bukti_transaksi": keuangan.BuktiTransaksi,
		"created_at":      keuangan.CreatedAt,
		"updated_at":      keuangan.UpdatedAt,
	}
	c.JSON(http.StatusOK, response)
}

func CreateKeuangan(c *gin.Context) {
	var keuangan models.Keuangan

	if err := c.ShouldBindJSON(&keuangan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := keuangan.ValidateTipeTransaksi(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var anggota models.Anggota
	if err := config.DB.First(&anggota, keuangan.AnggotaID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Anggota tidak ditemukan"})
		return
	}

	if err := config.DB.Create(&keuangan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Preload("Anggota").First(&keuangan, keuangan.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memuat data anggota"})
		return
	}

	c.JSON(http.StatusCreated, keuangan)
}

func UpdateKeuangan(c *gin.Context) {
	id := c.Param("id")
	var keuangan models.Keuangan

	if err := config.DB.First(&keuangan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&keuangan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&keuangan)
	c.JSON(http.StatusOK, keuangan)
}

func DeleteKeuangan(c *gin.Context) {
	id := c.Param("id")
	var keuangan models.Keuangan

	if err := config.DB.First(&keuangan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi tidak ditemukan"})
		return
	}

	config.DB.Delete(&keuangan)
	c.JSON(http.StatusOK, gin.H{"message": "Transaksi berhasil dihapus"})
}

func UploadBuktiTransaksi(c *gin.Context) {
	var keuangan models.Keuangan
	id := c.Param("id")

	if err := config.DB.First(&keuangan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Keuangan tidak ditemukan"})
		return
	}

	file, _ := c.FormFile("bukti_transaksi")
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File bukti transaksi tidak ditemukan"})
		return
	}
	validExtensions := []string{".jpg", ".jpeg", ".png", ".pdf"}
	fileExt := filepath.Ext(file.Filename)
	isValid := false

	for _, ext := range validExtensions {
		if ext == fileExt {
			isValid = true
			break
		}
	}

	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File tidak valid, hanya gambar dan PDF yang diperbolehkan"})
		return
	}

	filePath := "./uploads/" + file.Filename

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	keuangan.BuktiTransaksi = filePath
	if err := config.DB.Save(&keuangan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui bukti transaksi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Bukti transaksi berhasil di-upload",
		"keuangan": keuangan,
	})
}
func GetKeuanganStats(c *gin.Context) {

	bulan := c.DefaultQuery("bulan", fmt.Sprintf("%d", time.Now().Month()))
	tahun := c.DefaultQuery("tahun", fmt.Sprintf("%d", time.Now().Year()))

	type result struct {
		TotalPemasukan   float64
		TotalPengeluaran float64
		CountPemasukan   int64
		CountPengeluaran int64
		TotalTransaksi   int64
	}

	var res result

	config.DB.Model(&models.Keuangan{}).
		Where("bulan = ? AND tahun = ?", bulan, tahun).
		Select(`
			SUM(CASE WHEN tipe_transaksi = 'pemasukan' THEN nominal ELSE 0 END) AS total_pemasukan,
			SUM(CASE WHEN tipe_transaksi = 'pengeluaran' THEN nominal ELSE 0 END) AS total_pengeluaran,
			COUNT(CASE WHEN tipe_transaksi = 'pemasukan' THEN 1 ELSE NULL END) AS count_pemasukan,
			COUNT(CASE WHEN tipe_transaksi = 'pengeluaran' THEN 1 ELSE NULL END) AS count_pengeluaran,
			COUNT(*) AS total_transaksi`).
		Scan(&res)

	var avgTransaksi float64
	if res.TotalTransaksi > 0 {
		avgTransaksi = (res.TotalPemasukan + res.TotalPengeluaran) / float64(res.TotalTransaksi)
	}

	c.JSON(http.StatusOK, gin.H{
		"tahun":               tahun,
		"bulan":               bulan,
		"total_pemasukan":     res.TotalPemasukan,
		"total_pengeluaran":   res.TotalPengeluaran,
		"jumlah_pemasukan":    res.CountPemasukan,
		"jumlah_pengeluaran":  res.CountPengeluaran,
		"total_transaksi":     res.TotalTransaksi,
		"rata_rata_transaksi": avgTransaksi,
	})
}
