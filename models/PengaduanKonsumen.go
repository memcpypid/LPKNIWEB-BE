package models

import "time"

type PengaduanKonsumen struct {
	ID        uint               `gorm:"primary_key" json:"id"`
	Nama      string             `gorm:"type:varchar(100);not null" json:"nama"`
	Email     string             `gorm:"type:varchar(100);not null" json:"email"`
	Judul     string             `gorm:"type:varchar(255);not null" json:"judul"`
	Deskripsi string             `gorm:"type:text;not null" json:"deskripsi"`
	Status    string             `gorm:"type:varchar(50);default:'PENDING'" json:"status"`
	Publish   bool               `gorm:"type:boolean;default:false" json:"publish"`
	Teruskan  bool               `gorm:"type:boolean;default:false" json:"teruskan"`
	Media     []MediaPengaduan   `gorm:"foreignKey:PengaduanID" json:"media"`
	WilayahID *uint              `json:"wilayah_id,omitempty" gorm:"type:bigint"`
	Wilayah   *Wilayah           `json:"wilayah,omitempty" gorm:"foreignKey:WilayahID"`
	DaerahID  *uint              `json:"daerah_id,omitempty" gorm:"type:bigint"`
	Daerah    *Daerah            `json:"daerah,omitempty" gorm:"foreignKey:DaerahID"`
	JabatanID uint               `json:"jabatan_id" gorm:"type:bigint;not null"` // Foreign key to JabatanStruktural
	Jabatan   *JabatanStruktural `gorm:"foreignKey:JabatanID" json:"jabatan"`    // One-to-many relationship
	CreatedAt time.Time          `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time          `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// Model MediaPengaduan
type MediaPengaduan struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Tipe        string `gorm:"type:varchar(50);not null" json:"tipe"`      // Tipe media (gambar, video, dokumen)
	Filepath    string `gorm:"type:varchar(255);not null" json:"filepath"` // Path file media
	Deskripsi   string `gorm:"type:text" json:"deskripsi"`                 // Deskripsi media
	PengaduanID uint   `gorm:"not null" json:"pengaduan_id"`               // Foreign key ke PengaduanKonsumen
}
