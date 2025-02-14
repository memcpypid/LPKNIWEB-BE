package routes

import (
	"LPKNI/lpkni_project/controllers"

	"github.com/gin-gonic/gin"
)

func NewsRoutes(router *gin.Engine) {
	router.GET("/news", controllers.GetAllNews)
	router.POST("/news", controllers.CreateNews)
	router.GET("/news/:id", controllers.GetNewsByID)
	router.PUT("/news/:id", controllers.UpdateNews)
	router.DELETE("/news/:id", controllers.DeleteNews)

}
