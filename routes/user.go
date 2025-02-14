package routes

import (
	"LPKNI/lpkni_project/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users/register", controllers.RegisterUser)
	router.POST("/users/login", controllers.LoginUser)
	router.GET("/users/:id", controllers.GetUserByID)
	router.GET("/users", controllers.GetAllUsers)
	router.PUT("users/update", controllers.UpdateUserProfile)
}
