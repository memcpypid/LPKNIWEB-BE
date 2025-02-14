package models

import (
	"gorm.io/gorm"
)

type Pendaftaran struct {
	gorm.Model
	Nama            string `json:"nama" gorm:"type:varchar(100);not null"`
	Nik             string `json:"nik" gorm:"type:varchar(20);not null;unique"`
	NoTelepon       string `json:"noTelepon" gorm:"type:varchar(20);not null"`
	Email           string `json:"email" gorm:"type:varchar(100);not null;unique"`
	TipePendaftaran string `json:"tipePendaftaran" gorm:"type:varchar(50);not null"`
	Provinsi        string `json:"provinsi,omitempty" gorm:"type:varchar(50)"`
	KotaKabupaten   string `json:"kotaKabupaten,omitempty" gorm:"type:varchar(50)"`
	Jabatan         string `json:"jabatan" gorm:"type:varchar(50);not null"`
	Alamat          string `json:"alamat" gorm:"type:varchar(255);not null"`
	Foto3x4         string `json:"foto3x4" gorm:"type:varchar(255)"`
	FotoKtp         string `json:"fotoKtp" gorm:"type:varchar(255)"`
}

