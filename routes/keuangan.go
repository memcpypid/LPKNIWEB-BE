package routes

import (
	"LPKNI/lpkni_project/controllers"

	"github.com/gin-gonic/gin"
)

func KeuanganRoutes(router *gin.Engine) {
	router.GET("/keuangan", controllers.GetAllKeuangan)
	router.POST("/keuangan", controllers.CreateKeuangan)
	router.GET("/keuangan/:id", controllers.GetKeuanganByID)
	router.PUT("/keuangan/:id", controllers.UpdateKeuangan)
	router.DELETE("/keuangan/:id", controllers.DeleteKeuangan)
	router.POST("/keuangan/:id/buktitransaksi", controllers.UploadBuktiTransaksi)
	router.GET("/keuangan/statistik", controllers.GetKeuanganStats)

}
