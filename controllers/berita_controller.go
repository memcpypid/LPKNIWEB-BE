package controllers

import (
	"LPKNI/lpkni_project/config"
	"LPKNI/lpkni_project/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func CreateNews(c *gin.Context) {
	var news models.News

	
	news.Title = c.PostForm("title")
	news.Description = c.PostForm("description")
	news.Category = c.PostForm("category")
	dateStr := c.PostForm("date")

	
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal salah"})
		return
	}
	news.Date = parsedDate

	
	file, err := c.FormFile("gambar")
	if err == nil { 
		
		ext := filepath.Ext(file.Filename)

		
		newFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

		
		newFileName = strings.ReplaceAll(newFileName, " ", "_")

		
		filePath := filepath.Join("uploads", newFileName)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan gambar"})
			return
		}

		
		news.ImageURL = "/uploads/" + newFileName
	}


	
	if err := config.DB.Create(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan berita"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berita berhasil ditambahkan", "data": news})
}

func GetAllNews(c *gin.Context) {
	var news []models.News
	var total int64

	
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	
	config.DB.Model(&models.News{}).Count(&total)

	
	if err := config.DB.Limit(limit).Offset(offset).Find(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan dalam memuat berita: " + err.Error()})
		return
	}

	
	if len(news) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Tidak ada berita yang ditemukan"})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{
		"data":       news,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}

func GetNewsByID(c *gin.Context) {
	id := c.Param("id")
	var news models.News

	if err := config.DB.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "news tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, news)
}


func UpdateNews(c *gin.Context) {
    id := c.Param("id")
    var news models.News

    
    if err := config.DB.First(&news, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Berita tidak ditemukan"})
        return
    }

    
    title := c.PostForm("title")
    description := c.PostForm("description")
    category := c.PostForm("category")
    dateStr := c.PostForm("date")

    
    parsedDate, err := time.Parse("2006-01-02", dateStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal salah"})
        return
    }

    
    news.Title = title
    news.Description = description
    news.Category = category
    news.Date = parsedDate

    
    file, err := c.FormFile("gambar")
    if err == nil { 
        
        ext := filepath.Ext(file.Filename)
  
        newFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
        newFileName = strings.ReplaceAll(newFileName, " ", "_")

        
        filePath := filepath.Join("uploads", newFileName)
        if err := c.SaveUploadedFile(file, filePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan gambar"})
            return
        }
        if news.ImageURL != "" {
            oldImagePath := filepath.Join("uploads", filepath.Base(news.ImageURL))
            os.Remove(oldImagePath)
        }

        
        news.ImageURL = "/uploads/" + newFileName
    }

    
    if err := config.DB.Save(&news).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui berita"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Berita berhasil diperbarui", "data": news})
}

func DeleteNews(c *gin.Context) {
    id := c.Param("id")
    var news models.News

    // Cek apakah berita ada di database
    if err := config.DB.First(&news, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Berita tidak ditemukan"})
        return
    }

    // Jika berita memiliki gambar, hapus dari storage
    if news.ImageURL != "" {
        imagePath := filepath.Join("uploads", filepath.Base(news.ImageURL))
        os.Remove(imagePath) // Hapus file gambar dari sistem penyimpanan
    }

    // Hapus berita dari database
    if err := config.DB.Delete(&news).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus berita"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Berita berhasil dihapus"})
}
