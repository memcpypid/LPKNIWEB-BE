package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	protected := router.Group("/api")

	userRoutes(protected)
	daerah(protected)
	wilayah(protected)
	// router.POST("/api/login", controllers.Login)
	router.Static("/uploads", "./uploads")
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Apotek Management API!",
		})
	})
}
