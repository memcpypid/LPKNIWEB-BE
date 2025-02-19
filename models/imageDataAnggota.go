package models

import "time"

type ImageDataAnggota struct {
	ID         uint        `json:"id" gorm:"primaryKey;column:id_image_user"`
	DataUserID uint        `json:"dataUserId" gorm:"type:bigint;not null"`       // Foreign Key ke tabel DataUser
	DataUser   DataAnggota `json:"-" gorm:"foreignKey:DataUserID"`               // Relasi dengan DataUser
	ImageURL   string      `json:"imageUrl" gorm:"type:varchar(255);not null"`   // URL atau path gambar
	Keterangan string      `json:"keterangan" gorm:"type:varchar(100);not null"` // Keterangan gambar (misal: "KTP", "KK", dll)
	CreatedAt  time.Time   `json:"createdAt" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt  time.Time   `json:"updatedAt" gorm:"type:timestamp;default:current_timestamp"`
}
