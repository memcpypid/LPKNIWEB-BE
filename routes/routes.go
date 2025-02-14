package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	UserRoutes(router)
	AnggotaRoutes(router)
	NewsRoutes(router)
	KeuanganRoutes(router)
	AduanRoutes(router)
	TiketRoutes(router)
	ResponRoutes(router)
	SponsorRoutes(router)
	PiutangRoutes(router)
	AdminRoutes(router)
	PendaftaranRoutes(router)

	return router
}
