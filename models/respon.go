package models

import (
	"time"
)

type Respon struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	TiketID       uint      `gorm:"not null" json:"tiket_id"`
	Tiket         Tiket     `gorm:"foreignKey:TiketID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"tiket"`
	CSNama        string    `gorm:"size:100;not null" json:"cs_nama"`
	Respon        string    `gorm:"type:text;not null" json:"respon"`
	TanggalRespon time.Time `gorm:"autoCreateTime" json:"tanggal_respon"`
}
