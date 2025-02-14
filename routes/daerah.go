package routes

import (
	"LPKNI/lpkniService/controllers"

	"github.com/gin-gonic/gin"
)

func daerah(api *gin.RouterGroup) {
	api.POST("/daerah", controllers.CreateDaerah)       // Create User
	api.GET("/daerah", controllers.GetAllDaerah)   // Get User by ID
	api.GET("/daerah/:id", controllers.GetDaerahByID)   // Get User by ID
	api.PUT("/daerah/:id", controllers.UpdateDaerah)    // Update User
	api.DELETE("/daerah/:id", controllers.DeleteDaerah) // Delete User
}