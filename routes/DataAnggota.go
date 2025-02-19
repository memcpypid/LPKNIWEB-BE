package routes

import (
	"LPKNI/lpkniService/controllers"

	"github.com/gin-gonic/gin"
)

func DatauserRoutes(api *gin.RouterGroup) {
	api.POST("/data-anggota", controllers.CreateDataUserWithImage) // Create User
	// api.POST("/data-user", controllers.CreateDataUser)       // Create User
	api.GET("/data-anggota/:id", controllers.GetUserData)              // Get DataAnggota by ID
	api.GET("/data-anggota/user/:id", controllers.GetUserDataByIdUser) // Get AkunAnggota by ID
	api.PUT("/data-anggota/:id", controllers.UpdateUserData)           // Update User
	api.DELETE("/data-anggota/:id", controllers.DeleteUserData)        // Delete User
}
