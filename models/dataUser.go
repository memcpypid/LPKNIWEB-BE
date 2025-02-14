package models

import "time"

type DataUser struct {
	ID                  uint      `json:"id" gorm:"primaryKey;column:id_data_user"`
	UserID              uint      `json:"userId" gorm:"type:bigint;not null"`  // Foreign Key ke tabel User
	DaerahID            *uint     `json:"daerahId,omitempty" gorm:"type:bigint"`  // Foreign Key ke Daerah (opsional)
	Daerah              Daerah   `json:"daerah,omitempty" gorm:"foreignKey:DaerahID"`
	WilayahID           *uint     `json:"wilayahId,omitempty" gorm:"type:bigint"`  // Foreign Key ke Wilayah (opsional)
	Wilayah             Wilayah  `json:"wilayah,omitempty" gorm:"foreignKey:WilayahID"`
	JabatanStrukturalID uint      `json:"jabatanStrukturalId" gorm:"type:bigint;not null"`  // Foreign Key ke JabatanStruktural
	JabatanStruktural   JabatanStruktural `json:"jabatanStruktural" gorm:"foreignKey:JabatanStrukturalID"`
	Alamat              string    `json:"alamat" gorm:"type:varchar(255);"` // Alamat pengguna
	TanggalLahir        *time.Time `json:"tanggalLahir,omitempty" gorm:"type:date"` // Tanggal lahir pengguna
	NIK                 string    `json:"nik" gorm:"type:varchar(20);unique"`  // Nomor Induk Kependudukan (NIK)
	NoKK                string    `json:"noKK" gorm:"type:varchar(20)"`        // Nomor Kartu Keluarga (KK)
	TempatLahir         string    `json:"tempatLahir" gorm:"type:varchar(100)"` // Tempat lahir
	Pekerjaan           string    `json:"pekerjaan" gorm:"type:varchar(100)"`   // Pekerjaan pengguna
	StatusPerkawinan    string    `json:"statusPerkawinan" gorm:"type:varchar(50)"` // Status perkawinan (misal: menikah, belum menikah)
	Agama               string    `json:"agama" gorm:"type:varchar(50)"`
	Status              string    `json:"status" gorm:"type:enum('PENDING', 'SUCCESS', 'CANCEL');not null"`
	CreatedAt           time.Time `json:"createdAt" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt           time.Time `json:"updatedAt" gorm:"type:timestamp;default:current_timestamp"`
	ImageUsers []ImageUser `json:"imageUsers" gorm:"foreignKey:DataUserID"` // Satu DataUser memiliki banyak ImageUser
}




