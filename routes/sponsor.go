package routes

import (
	"LPKNI/lpkni_project/controllers"

	"github.com/gin-gonic/gin"
)

func SponsorRoutes(router *gin.Engine) {
	router.POST("/sponsor/daftar", controllers.CreatePengajuanSponsor)
	router.GET("/sponsor/:id", controllers.GetPengajuanSponsorByID)
	router.GET("/sponsor", controllers.GetAllPengajuanSponsor)
	router.PUT("/sponsor", controllers.UpdatePengajuanSponsor)
	router.DELETE("/sponsor/:id", controllers.DeletePengajuanSponsor)
	router.POST("/sponsor/:id/upload", controllers.UploadProposal)
}
