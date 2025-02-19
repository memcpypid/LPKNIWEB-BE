package controllers

import (
	"LPKNI/lpkniService/config"
	"LPKNI/lpkniService/models"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateUserData - Create a new user and upload the image
func CreateDataUserWithImage(c *gin.Context) {
	UserID, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId Tidak Valid"})
		return
	}
	DaerahID := c.PostForm("daerahId")
	WilayahID := c.PostForm("wilayahId")
	JabatanStrukturalID := c.PostForm("jabatanStrukturalId")
	NamaLengkap := c.PostForm("nama_lengkap")
	Alamat := c.PostForm("alamat")
	TanggalLahir, err := time.Parse("2006-01-02", c.PostForm("tanggalLahir"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tanggal Tidak Valid"})
		return
	}
	NIK := c.PostForm("nik")
	TempatLahir := c.PostForm("tempatLahir")
	Pekerjaan := c.PostForm("pekerjaan")
	StatusPerkawinan := c.PostForm("statusPerkawinan")
	Agama := c.PostForm("agama")
	Status := c.PostForm("status")
	var DaerahIDUInt *uint
	if DaerahID != "null" {
		DaerahIDInt, err := strconv.Atoi(DaerahID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "DaerahID Tidak Valid"})
			return
		}
		// Only assign the pointer if the value is valid
		DaerahIDUInt = new(uint)
		*DaerahIDUInt = uint(DaerahIDInt)
	}
	WilayahIDInt, err := strconv.Atoi(WilayahID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "WilayahID Tidak Valid"})
		return
	}

	JabatanStrukturalIDInt, err := strconv.Atoi(JabatanStrukturalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JabatanStrukturalID Tidak Valid"})
		return
	}

	user := models.DataAnggota{
		UserID:              uint(UserID),
		DaerahID:            DaerahIDUInt,                 // convert to uint
		WilayahID:           uint(WilayahIDInt),           // convert to uint
		JabatanStrukturalID: uint(JabatanStrukturalIDInt), // convert to uint
		NamaLengkap:         NamaLengkap,
		Alamat:              Alamat,
		TanggalLahir:        TanggalLahir,
		NIK:                 NIK,
		TempatLahir:         TempatLahir,
		Pekerjaan:           Pekerjaan,
		StatusPerkawinan:    StatusPerkawinan,
		Agama:               Agama,
		Status:              Status,
	}

	// Save the DataUser to the database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Handle file upload for fileFoto3x4
	if file, err := c.FormFile("fileFoto3x4"); err == nil {
		// Get keteranganFoto3x4 using PostForm
		keteranganFoto3x4 := c.PostForm("keteranganFoto3x4")

		// Generate a new file name based on user ID and timestamp (to avoid conflicts)
		newFileName := fmt.Sprintf("%d_%s", user.ID, time.Now().Format("20060102150405"))

		// Get the file extension of the uploaded file
		ext := filepath.Ext(file.Filename)

		// Combine the new file name with the file extension
		destination := fmt.Sprintf("uploads/data-anggota/pas-foto/%s%s", newFileName, ext)

		// Save the file with the new name
		if err := c.SaveUploadedFile(file, destination); err == nil {
			image := models.ImageDataAnggota{
				DataUserID: user.ID,
				ImageURL:   destination,
				Keterangan: keteranganFoto3x4,
			}
			if err := config.DB.Create(&image).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save Foto 3x4"})
			return
		}
	}

	// Handle file upload for fileFotoKtp
	if file, err := c.FormFile("fileFotoKtp"); err == nil {
		// Get keteranganFotoKtp using PostForm
		keteranganFotoKtp := c.PostForm("keteranganFotoKtp")

		// Generate a new file name based on user ID and timestamp (to avoid conflicts)
		newFileName := fmt.Sprintf("%d_%s", user.ID, time.Now().Format("20060102150405"))

		// Get the file extension of the uploaded file
		ext := filepath.Ext(file.Filename)

		// Combine the new file name with the file extension
		destination := fmt.Sprintf("uploads/data-anggota/ktp/%s%s", newFileName, ext)

		// Save the file with the new name
		if err := c.SaveUploadedFile(file, destination); err == nil {
			image := models.ImageDataAnggota{
				DataUserID: user.ID,
				ImageURL:   destination,
				Keterangan: keteranganFotoKtp,
			}
			if err := config.DB.Create(&image).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save Foto KTP"})
			return
		}
	}

	// Successfully saved data and images
	c.JSON(http.StatusOK, gin.H{"message": "Data and images successfully saved"})
}

// GetUserData - Get a user by ID with their image
func GetUserData(c *gin.Context) {
	id := c.Param("id")
	var user models.DataAnggota

	// Fetch user by ID and include related ImageUser records
	if err := config.DB.Preload("ImageUsers").Preload("Wilayah").Preload("Daerah").Preload("JabatanStruktural").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
func GetUserDataByIdUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id Tidak Valid"})
		return
	}
	var user models.DataAnggota

	// Fetch user by ID and include related ImageUser records
	if err := config.DB.Preload("ImageUsers").Preload("Wilayah").Preload("Daerah").Preload("JabatanStruktural").Where(models.DataAnggota{UserID: uint(id)}).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data user Tidak Ditemukan"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserData - Update user details and upload a new image
func UpdateUserData(c *gin.Context) {
	id := c.Param("id")
	var user models.DataAnggota

	// Fetch existing user by ID
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind JSON body to DataUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update user details in database
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	// Handle file upload (if exists)
	file, err := c.FormFile("file")
	if err == nil {
		// Save the file to the "uploads" directory
		destination := fmt.Sprintf("uploads/%s", file.Filename)
		if err := c.SaveUploadedFile(file, destination); err == nil {
			// Update image in ImageUser table
			image := models.ImageDataAnggota{
				DataUserID: user.ID,
				ImageURL:   destination,
				Keterangan: c.DefaultPostForm("keterangan", "No description"),
			}
			if err := config.DB.Create(&image).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving image"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error uploading file"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

// DeleteUserData - Delete user and their image(s)
func DeleteUserData(c *gin.Context) {
	id := c.Param("id")
	var user models.DataAnggota

	// Fetch user by ID
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete all associated images
	if err := config.DB.Where("data_user_id = ?", user.ID).Delete(&models.ImageDataAnggota{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting images"})
		return
	}

	// Delete the user from database
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User and associated images deleted successfully"})
}

// func CreateUserData(c *gin.Context) {
// 	var user models.DataUser
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	config.DB.Create(&user)

// 	file, err := c.FormFile("file")
// 	if err == nil {
// 		destination := fmt.Sprintf("uploads/%s", file.Filename)
// 		if err := c.SaveUploadedFile(file, destination); err == nil {
// 			image := models.ImageUser{
// 				DataUserID: user.ID,
// 				ImageURL:   destination,
// 				Keterangan: c.PostForm("keterangan"),
// 			}
// 			config.DB.Create(&image)
// 		}
// 	}

// 	c.JSON(http.StatusOK, user)
// }

// // CreateDataUser untuk membuat data pengguna baru
// func CreateDataUser(c *gin.Context) {
// 	var dataUser models.DataUser
// 	if err := c.ShouldBindJSON(&dataUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Simpan data pengguna ke database
// 	if err := config.DB.Create(&dataUser).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "DataUser creation failed"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "DataUser created successfully", "dataUser": dataUser})
// }

// // GetDataUser untuk mendapatkan detail data pengguna berdasarkan ID
// func GetDataUserById(c *gin.Context) {
// 	dataUserID := c.Param("id")
// 	var dataUser models.DataUser

// 	if err := config.DB.First(&dataUser, dataUserID).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "DataUser not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dataUser)
// }

// // UpdateDataUser untuk memperbarui data pengguna
// func UpdateDataUser(c *gin.Context) {
// 	dataUserID := c.Param("id")
// 	var dataUser models.DataUser

// 	// Bind data dari request body
// 	if err := c.ShouldBindJSON(&dataUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Mencari DataUser berdasarkan ID dan mengupdate
// 	if err := config.DB.Model(&dataUser).Where("id = ?", dataUserID).Updates(dataUser).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update DataUser"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "DataUser updated successfully"})
// }

// // DeleteDataUser untuk menghapus data pengguna berdasarkan ID
// func DeleteDataUser(c *gin.Context) {
// 	dataUserID := c.Param("id")
// 	var dataUser models.DataUser

// 	if err := config.DB.Delete(&dataUser, dataUserID).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete DataUser"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "DataUser deleted successfully"})
// }
