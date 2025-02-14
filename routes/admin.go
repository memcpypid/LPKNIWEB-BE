package routes

import (
	"LPKNI/lpkni_project/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	router.GET("/admin", controllers.GetAllAdmins)
	router.POST("/admin", controllers.RegisterAdmin)
	router.GET("/admin/:id", controllers.GetAdminByID)
	router.DELETE("/admin/:id", controllers.DeleteAdmin)
	router.POST("/admin/:id/upload", controllers.UploadSuratIjinAkses)
	router.POST("/admin/login", controllers.LoginAdmin)
}
