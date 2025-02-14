package models

import (
	"time"
)

type Piutang struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Subjek           string    `gorm:"size:255;not null" json:"subjek"`
	Deskripsi        string    `gorm:"type:text;not null" json:"deskripsi"`
	TanggalPengaduan time.Time `gorm:"autoCreateTime" json:"tanggal_pengaduan"`
	Status           string    `gorm:"type:enum('pending', 'in_progress', 'resolved');default:'pending';not null" json:"status"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
