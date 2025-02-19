package routes

import (
	"LPKNI/lpkniService/controllers"

	"github.com/gin-gonic/gin"
)

func Berita(api *gin.RouterGroup) {
	api.POST("/berita", controllers.CreateBerita)       // Create User
	api.GET("/berita", controllers.GetAllBerita)   // Get User by ID
	api.GET("/berita/:id", controllers.GetBeritaByID)   // Get User by ID
	api.PUT("/berita/:id", controllers.UpdateBerita)    // Update User
	api.DELETE("/berita/:id", controllers.DeleteBerita) // Delete User
	
	api.POST("/berita/kategori", controllers.CreateKategori)       // Create User
	api.GET("/berita/kategori", controllers.GetAllKategori)   // Get User by ID
	api.GET("/berita/kategori/:id", controllers.GetKategoriByID)   // Get User by ID
	api.PUT("/berita/kategori/:id", controllers.UpdateKategori)    // Update User
	api.DELETE("/berita/kategori/:id", controllers.DeleteKategori) // Delete User
}