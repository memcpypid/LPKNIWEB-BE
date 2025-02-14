package models

type JabatanStruktural struct {
	ID               uint   `json:"id" gorm:"primaryKey;column:id_jabatan_struktural"`
	Nama             string `json:"nama" gorm:"type:varchar(50);not null"` // Nama jabatan struktural, misalnya "Kepala", "Sekretaris"
	MaksimumAnggota  int    `json:"maksimumAnggota" gorm:"type:int;not null"` // Batas maksimum anggota yang dapat memiliki jabatan ini
}
