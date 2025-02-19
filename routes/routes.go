package routes

import (
	"LPKNI/lpkniService/handlers"
	"LPKNI/lpkniService/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	protected := router.Group("/api")

	userRoutes(protected)
	Daerah(protected)
	Wilayah(protected)
	Berita(protected)
	Pengaduan(protected)

	Jabatan(protected)
	protected.Use(middleware.VerifyJWT())
	{
		// Contoh rute yang dilindungi
		protected.GET("/profile", func(c *gin.Context) {
			user, _ := c.Get("user")
			akun, _ := c.Get("data_anggota")
			c.JSON(200, gin.H{
				"user":         user,
				"data_anggota": akun,
			})
		})
		DatauserRoutes(protected)
	}
	router.POST("/api/auth/login", handlers.Login)
	router.POST("/api/auth/logout", handlers.Logout)
	router.Static("/uploads", "./uploads")
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Apotek Management API!",
		})
	})
}
