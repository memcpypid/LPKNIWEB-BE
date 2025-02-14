package models

import "time"

type Berita struct {
	ID        uint        `gorm:"primary_key"`
	Judul     string      `gorm:"type:varchar(255);not null"`
	Konten    string      `gorm:"type:text;not null"`
	Status    string      `gorm:"type:varchar(50);default:'draft'"`
	Tanggal   time.Time   `gorm:"type:datetime;not null"` // Gunakan time.Time untuk tanggal
	Penulis   string      `gorm:"type:varchar(100);not null"`
	Media     []MediaBerita     `gorm:"foreignkey:BeritaID"` // Relasi satu ke banyak dengan Media
	Kategoris []KategoriBerita  `gorm:"foreignkey:BeritaID"` // Relasi satu ke banyak dengan Kategori
	CreatedAt time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP"` // Waktu pembuatan
	UpdatedAt time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // Waktu pembaruan
}

// Model Media
type MediaBerita struct {
	ID        uint   `gorm:"primary_key"`
	Link      string `gorm:"type:varchar(255);not null"` // Link YouTube atau URL gambar/video
	Tipe      string `gorm:"type:varchar(50);not null"`  // Tipe media (image, video, youtube)
	BeritaID  uint   `gorm:"not null"`                  // ID Berita sebagai foreign key
	Deskripsi string `gorm:"type:text"`                 // Deskripsi opsional untuk media
	Filepath  string `gorm:"type:varchar(255)"`         // Untuk menyimpan lokasi file jika berupa gambar/video
}

// Model Kategori
type KategoriBerita struct {
	ID        uint    `gorm:"primary_key"`
	Nama      string  `gorm:"type:varchar(100);unique;not null"`
	BeritaID  uint    `gorm:"not null"`  // Foreign key untuk relasi satu ke banyak dengan Berita
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
