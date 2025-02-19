package routes

import (
	"LPKNI/lpkniService/controllers"

	"github.com/gin-gonic/gin"
)

func Jabatan(api *gin.RouterGroup) {
	api.POST("/jabatan", controllers.CreateJabatan)                    // Create Jabatan
	api.GET("/jabatan", controllers.GetAllJabatan)                     // Get Jabatan all
	api.GET("/jabatan/:id", controllers.GetJabatanById)                // Get Jabatan by ID
	api.PUT("/jabatan/:id", controllers.UpdateJabatan)                 // Update Jabatan
	api.DELETE("/jabatan/:id", controllers.DeleteJabatan)              // Delete Jabatan
	api.GET("/penggunaa-jabatan/:id", controllers.CountJabatanMembers) // Get Jabatan by ID
}
