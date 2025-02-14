package routes

import (
	"LPKNI/lpkni_project/controllers"

	"github.com/gin-gonic/gin"
)

func ResponRoutes(router *gin.Engine) {

	router.GET("/respon", controllers.GetAllRespon)
	router.GET("/respon/:id", controllers.GetResponByID)
	router.POST("/respon", controllers.CreateRespon)
	router.PUT("/respon/:id", controllers.UpdateRespon)
	router.DELETE("/respon/:id", controllers.DeleteRespon)
	router.GET("/respon/tiket/:tiket_id", controllers.GetResponByTiketID)
}
