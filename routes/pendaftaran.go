package routes

import (
	"LPKNI/lpkni_project/controllers"

	"github.com/gin-gonic/gin"
)

func PendaftaranRoutes(router *gin.Engine) {

	router.GET("/pendaftaran/:id", controllers.GetPendaftaranByID)
	router.POST("/pendaftaran", controllers.CreatePendaftaran)
	router.DELETE("/pendaftaran/:id", controllers.DeletePendaftaran)
	router.PUT("/pendaftaran/:id", controllers.UpdatePendaftaran)
	router.POST("/pendaftaran/:id", controllers.UploadFoto)
	router.GET("/pendaftaran", controllers.GetAllPendaftaran)
}
