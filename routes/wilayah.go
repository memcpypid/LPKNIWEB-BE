package routes

import (
	"LPKNI/lpkniService/controllers"

	"github.com/gin-gonic/gin"
)

func wilayah(api *gin.RouterGroup) {
	api.POST("/wilayah", controllers.CreateWilayah)       // Create User
	api.GET("/wilayah", controllers.GetAllWilayah)   // Get User by ID
	api.GET("/wilayah/:id", controllers.GetWilayahByID)   // Get User by ID
	api.PUT("/wilayah/:id", controllers.UpdateWilayah)    // Update User
	api.DELETE("/wilayah/:id", controllers.DeleteWilayah) // Delete User
}