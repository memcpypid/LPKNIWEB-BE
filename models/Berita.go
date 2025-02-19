package models

import "time"

// Model Berita
type Berita struct {
	ID        uint             `gorm:"primary_key" json:"id"`
	Judul     string           `gorm:"type:varchar(255);not null" json:"judul"`
	Konten    string           `gorm:"type:text;not null" json:"konten"`
	Status    string           `gorm:"type:varchar(50);default:'DRAFT'" json:"status"`
	Tanggal   time.Time        `gorm:"type:datetime;not null" json:"tanggal"`
	Penulis   string           `gorm:"type:varchar(100);not null" json:"penulis"`
	DaerahID  *uint            `json:"daerahId,omitempty" gorm:"type:bigint"` // Foreign Key ke Daerah (opsional)
	Daerah    Daerah           `json:"daerah,omitempty" gorm:"foreignKey:DaerahID"`
	WilayahID *uint            `json:"wilayahId,omitempty" gorm:"type:bigint"` // Foreign Key ke Wilayah (opsional)
	Wilayah   Wilayah          `json:"wilayah,omitempty" gorm:"foreignKey:WilayahID"`
	Media     []MediaBerita    `gorm:"foreignkey:BeritaID" json:"media"`                   // Relasi satu ke banyak dengan Media
	Kategori  []KategoriBerita `gorm:"many2many:berita_kategori;not null" json:"kategori"` // Relasi banyak ke banyak dengan Kategori
	CreatedAt time.Time        `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time        `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// Model Media
type MediaBerita struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Tipe      string `gorm:"type:varchar(50);not null" json:"tipe"`
	BeritaID  uint   `gorm:"not null" json:"berita_id"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
	Filepath  string `gorm:"type:varchar(255)" json:"filepath"`
}

// Model Kategori
type KategoriBerita struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Nama      string    `gorm:"type:varchar(100);unique;not null" json:"nama"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}
