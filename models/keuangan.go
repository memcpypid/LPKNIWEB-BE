package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	Pemasukan   = "pemasukan"
	Pengeluaran = "pengeluaran"
)

type Keuangan struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AnggotaID      uint      `gorm:"not null" json:"anggota_id"`
	Anggota        Anggota   `gorm:"foreignKey:AnggotaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"anggota"`
	Nominal        float64   `gorm:"not null" json:"nominal"`
	TipeTransaksi  string    `gorm:"type:enum('pemasukan', 'pengeluaran');not null" json:"tipe_transaksi"`
	Deskripsi      string    `gorm:"type:text" json:"deskripsi"`
	BuktiTransaksi string    `gorm:"size:255" json:"bukti_transaksi"`
	Bulan          int       `gorm:"not null" json:"bulan"`
	Tahun          int       `gorm:"not null" json:"tahun"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (k *Keuangan) ValidateTipeTransaksi() error {
	if k.TipeTransaksi != Pemasukan && k.TipeTransaksi != Pengeluaran {
		return fmt.Errorf("tipe transaksi harus 'pemasukan' atau 'pengeluaran'")
	}
	return nil
}

func (k *Keuangan) BeforeCreate(tx *gorm.DB) (err error) {

	t := time.Now()
	k.Bulan = int(t.Month())
	k.Tahun = t.Year()
	return
}
