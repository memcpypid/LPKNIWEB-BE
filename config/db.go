package config

import (
	"LPKNI/lpkniService/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "root:password@tcp(127.0.0.1:3306)/lpkni_web_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Koneksi ke database gagal: ", err)
	} else {
		fmt.Println("Berhasil terhubung ke database")
	}
	if err := db.AutoMigrate(
		&models.AkunAnggota{},
		&models.DataAnggota{},
		&models.Wilayah{},
		&models.Daerah{},
		&models.ImageDataAnggota{},
		&models.Berita{},
		&models.MediaBerita{},
		&models.KategoriBerita{},
		&models.PengaduanKonsumen{},
		&models.MediaPengaduan{},
		&models.SessionLogin{},
	); err != nil {
		log.Fatalf("Gagal melakukan migrasi: %v", err)
	}

	log.Println("Migrasi database berhasil!")
	DB = db
}
