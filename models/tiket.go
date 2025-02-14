package models

import (
	"time"
)

type Tiket struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaPengguna      string    `gorm:"size:50; not null" json:"nama_pengguna"`
	EmailPengguna     string    `gorm:"size:100;not null" json:"email_pengguna"`
	Subjek            string    `gorm:"size:255;not null" json:"subjek"`
	Deskripsi         string    `gorm:"type:text;not null" json:"deskripsi"`
	Status            string    `gorm:"type:enum('open', 'in_progress', 'resolved', 'closed');default:'open';not null" json:"status"`
	TanggalBuat       time.Time `gorm:"autoCreateTime" json:"tanggal_buat"`
	TanggalDiperbarui time.Time `gorm:"autoUpdateTime" json:"tanggal_diperbarui"`
}
