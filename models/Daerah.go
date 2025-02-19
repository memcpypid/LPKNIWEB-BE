// models/daerah.go
package models

type Daerah struct {
	ID         uint     `json:"id_daerah" gorm:"primaryKey;column:id_daerah"`        // ID daerah (Primary Key)
	NamaDaerah string   `json:"nama_daerah" gorm:"type:varchar(50);not null;unique"` // Nama Daerah, harus unik
	KodeDaerah string   `json:"kode_daerah" gorm:"type:varchar(50);not null"`        // Kode Daerah (misalnya '11.01' untuk Aceh Selatan)
	WilayahID  uint     `json:"-" gorm:"type:bigint;not null"`                       // Foreign Key untuk Wilayah (ID wilayah)
	Wilayah    *Wilayah `json:"-" gorm:"foreignKey:WilayahID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
