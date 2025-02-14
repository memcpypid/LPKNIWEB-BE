package routes

import (
	"LPKNI/lpkni_project/controllers"

	"github.com/gin-gonic/gin"
)

func AnggotaRoutes(router *gin.Engine) {
	router.POST("/anggota/daftar", controllers.PendaftaranAnggota)
	router.POST("/anggota/login/", controllers.LoginAnggota)
	router.GET("/anggota/:id", controllers.GetAnggotaByID)
	router.GET("/anggota", controllers.GetAllAnggota)
	router.POST("/anggota/:id/uploadsurat", controllers.UploadSuratIjinAksesUs)
}
