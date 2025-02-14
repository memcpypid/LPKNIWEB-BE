package models

import (
	"time"
)

type Aduan struct {
	ID           uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerName string         `gorm:"size:100;not null" json:"customer_name"`
	Email        string         `gorm:"size:100;not null" json:"email"`
	NoHp         string         `gorm:"size:15;not null" json:"no_hp"`
	Deskripsi    string         `gorm:"type:text;not null" json:"deskripsi"`
	TanggalAduan string         `gorm:"not null" json:"tanggal_aduan"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	AduanDokumen []AduanDokumen `gorm:"foreignKey:AduanID;constraint:OnDelete:CASCADE" json:"aduan_dokumen"`
}

type AduanDokumen struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	AduanID  uint   `gorm:"not null" json:"aduan_id"`
	FileURL  string `gorm:"type:varchar(255);not null" json:"file_url"`
	FileName string `gorm:"type:varchar(255);not null" json:"file_name"`
}
