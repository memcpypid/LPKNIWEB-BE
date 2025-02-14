package controllers

import (
	"LPKNI/lpkni_project/config"
	"LPKNI/lpkni_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}


func RegisterUser(c *gin.Context) {
	var user models.User

	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	
	var existingUser models.User
	if err := config.DB.Where("username = ? OR email_us = ?", user.Username, user.EmailUs).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username atau email sudah digunakan"})
		return
	}

	
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password: " + err.Error()})
		return
	}
	user.Password = hashedPassword

	
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User successfully registered",
		"user":    user,
	})
}


type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}


func LoginUser(c *gin.Context) {
	var input LoginRequest

	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format: " + err.Error()})
		return
	}

	var user models.User

	
	if err := config.DB.Where("email_us = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	
	userResponse := gin.H{
		"id":               user.ID,
		"jenisPendaftaran": user.JenisPendaftaran,
		"firstName":        user.FirstName,
		"lastName":         user.LastName,
		"contactNo":        user.ContactNo,
		"emailUs":          user.EmailUs,
		"username":         user.Username,
		"daerah":           user.Daerah,
		"wilayah":          user.Wilayah,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    userResponse,
	})
}


func GetAllUsers(c *gin.Context) {
	var users []models.User

	
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}


func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}



func UpdateUserProfile(c *gin.Context) {
	
	email := c.DefaultQuery("email", "")
	if email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email is required"})
		return
	}

	var input struct {
		FullName string `json:"fullName"`
	}

	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	if err := config.DB.Model(&models.User{}).Where("email_us = ?", email).Update("first_name", input.FullName).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
