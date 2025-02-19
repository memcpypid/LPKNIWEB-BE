package models

type Wilayah struct {
	ID          uint     `json:"id_wilayah" gorm:"primaryKey;column:id_wilayah"`
	NamaWilayah string   `json:"nama_wilayah" gorm:"type:varchar(50);not null;unique"` // Nama Provinsi, harus unik
	KodeWilayah string   `json:"kode_wilayah" gorm:"type:varchar(50);not null"`        // Kode Wilayah
	Daerah      []Daerah `json:"-" gorm:"foreignKey:WilayahID"`
}
