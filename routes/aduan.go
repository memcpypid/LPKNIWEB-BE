package routes

import (
	"LPKNI/lpkni_project/controllers"

	"github.com/gin-gonic/gin"
)

func AduanRoutes(router *gin.Engine) {
	router.GET("/aduan", controllers.GetAllAduan)
	router.POST("/aduan", controllers.CreateAduan)
	router.GET("/aduan/:id", controllers.GetAduanByID)
	router.PUT("/aduan/:id", controllers.UpdateAduan)
	router.DELETE("/aduan/:id", controllers.DeleteAduan)

	router.POST("/aduan/:id/dokumen", controllers.UploadDokumenAduan)
}
