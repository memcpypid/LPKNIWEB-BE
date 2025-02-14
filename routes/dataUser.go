package routes

import (
	"LPKNI/lpkniService/controllers"

	"github.com/gin-gonic/gin"
)

func DatauserRoutes(api *gin.RouterGroup) {
	api.POST("/data-user", controllers.CreateDataUser)       // Create User
	api.GET("/data-user/:id", controllers.GetDataUserById)   // Get User by ID
	api.PUT("/data-user/:id", controllers.UpdateDataUser)    // Update User
	api.DELETE("/data-user/:id", controllers.DeleteDataUser) // Delete User
}