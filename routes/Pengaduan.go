package routes

import (
	"LPKNI/lpkniService/controllers"

	"github.com/gin-gonic/gin"
)

func Pengaduan(api *gin.RouterGroup) {
	api.POST("/pengaduan", controllers.CreatePengaduan)       // Create User
	api.GET("/pengaduan", controllers.GetAllPengaduan)        // Get User by ID
	api.GET("/pengaduan/:id", controllers.GetPengaduanById)   // Get User by ID
	api.PUT("/pengaduan/:id", controllers.UpdatePengaduan)    // Update User
	api.DELETE("/pengaduan/:id", controllers.DeletePengaduan) // Delete User
}
