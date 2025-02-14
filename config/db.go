package config

import (
	"LPKNI/lpkni_project/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "root:password@tcp(127.0.0.1:3306)/lpkni-project?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Koneksi ke database gagal: ", err)
	} else {
		fmt.Println("Berhasil terhubung ke database")
	}
	if err := db.AutoMigrate(

		&models.User{},
		&models.Pendaftaran{},
		&models.News{},

	); err != nil {
		log.Fatalf("Gagal melakukan migrasi: %v", err)
	}

	log.Println("Migrasi database berhasil!")
	DB = db
}
