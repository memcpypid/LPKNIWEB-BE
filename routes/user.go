package routes

import (
	"LPKNI/lpkniService/controllers"

	"github.com/gin-gonic/gin"
)

func userRoutes(api *gin.RouterGroup) {
	api.POST("/user", controllers.CreateUser)       // Create User
	api.GET("/user/:id", controllers.GetUserById)   // Get User by ID
	api.PUT("/user/:id", controllers.UpdateUser)    // Update User
	api.DELETE("/user/:id", controllers.DeleteUser) // Delete User
}
