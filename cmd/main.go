package main

import (
	"LPKNI/lpkni_project/config"
	"LPKNI/lpkni_project/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	config.ConnectDatabase()

	// Inisialisasi router Gin
	r := gin.Default()
	routes.PendaftaranRoutes(r)


	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // Ubah sesuai domain frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Tambahkan pesan selamat datang di root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Selamat datang di sistem LPKNI",
		})
	})

	// Tambahkan rute untuk API pengguna
	routes.UserRoutes(r)
r.Static("uploads","./uploads")
	// **Tambahkan rute untuk API berita**
	routes.NewsRoutes(r) // <-- Ini yang hilang

	// Jalankan server
	log.Println("Server berjalan di port 5000")
	if err := r.Run(":3000"); err != nil {
		log.Fatal("Server gagal dijalankan: ", err)
	}
}
