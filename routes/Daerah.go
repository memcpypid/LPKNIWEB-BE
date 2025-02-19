package routes

import (
	"LPKNI/lpkniService/controllers"

	"github.com/gin-gonic/gin"
)

func Daerah(api *gin.RouterGroup) {
	api.POST("/daerah", controllers.CreateDaerah)                    // Create Daerah
	api.GET("/daerah", controllers.GetAllDaerah)                     // Get Daerah by ID
	api.GET("/daerah/wilayah/:id", controllers.GetDaerahByIDWilayah) // Get Daerah by ID
	api.GET("/daerah/:id", controllers.GetDaerahByID)                // Get Daerah by ID
	api.PUT("/daerah/:id", controllers.UpdateDaerah)                 // Update Daerah
	api.DELETE("/daerah/:id", controllers.DeleteDaerah)              // Delete Daerah
}
