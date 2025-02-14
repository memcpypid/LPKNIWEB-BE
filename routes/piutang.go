package routes

import (
	"LPKNI/lpkni_project/controllers"

	"github.com/gin-gonic/gin"
)

func PiutangRoutes(router *gin.Engine) {

	router.GET("/piutang/:id", controllers.GetPiutangByID)
	router.POST("/piutang", controllers.CreatePiutang)
	router.DELETE("/piutang/:id", controllers.DeletePiutang)
}
