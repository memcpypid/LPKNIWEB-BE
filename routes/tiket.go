package routes

import (
	"LPKNI/lpkni_project/controllers"

	"github.com/gin-gonic/gin"
)

func TiketRoutes(router *gin.Engine) {
	router.GET("/tiket", controllers.GetAllTiket)
	router.POST("/tiket", controllers.CreateTiket)
	router.GET("/tiket/:id", controllers.GetTiketByID)
	router.PUT("/tiket/:id/status", controllers.UpdateTiket)
	router.DELETE("/tiket/:id", controllers.DeleteTiket)

}
