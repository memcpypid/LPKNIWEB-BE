package models

import "time"

type DataAnggota struct {
	ID                  uint               `json:"id_data_anggota" gorm:"primaryKey;column:id_data_anggota"`
	UserID              uint               `json:"userId" gorm:"type:bigint;not null"`    // Foreign Key ke tabel User
	DaerahID            *uint              `json:"daerahId,omitempty" gorm:"type:bigint"` // Foreign Key ke Daerah (opsional)
	Daerah              Daerah             `json:"daerah,omitempty" gorm:"foreignKey:DaerahID"`
	WilayahID           uint               `json:"wilayahId" gorm:"type:bigint;not null"` // Foreign Key ke Wilayah (wajib)
	Wilayah             Wilayah            `json:"wilayah" gorm:"foreignKey:WilayahID"`
	JabatanStrukturalID uint               `json:"jabatanStrukturalId" gorm:"type:bigint;not null"` // Foreign Key ke JabatanStruktural
	JabatanStruktural   JabatanStruktural  `json:"jabatanStruktural" gorm:"foreignKey:JabatanStrukturalID"`
	NamaLengkap         string             `json:"nama_lengkap" gorm:"type:varchar(255);not null"`                                   // Wajib
	Alamat              string             `json:"alamat" gorm:"type:varchar(255);not null"`                                         // Wajib
	TanggalLahir        time.Time          `json:"tanggalLahir" gorm:"type:date;not null"`                                           // Wajib
	NIK                 string             `json:"nik" gorm:"type:varchar(20);unique;not null"`                                      // Wajib
	TempatLahir         string             `json:"tempatLahir" gorm:"type:varchar(100);not null"`                                    // Wajib
	Pekerjaan           string             `json:"pekerjaan" gorm:"type:varchar(100);not null"`                                      // Wajib
	StatusPerkawinan    string             `json:"statusPerkawinan" gorm:"type:varchar(50);not null"`                                // Wajib
	Agama               string             `json:"agama" gorm:"type:varchar(50);not null"`                                           // Wajib
	Status              string             `json:"status" gorm:"type:enum('PENDING', 'SUCCESS', 'CANCEL');default:PENDING;not null"` // Wajib
	CreatedAt           time.Time          `json:"createdAt" gorm:"type:timestamp;default:current_timestamp;not null"`
	UpdatedAt           time.Time          `json:"updatedAt" gorm:"type:timestamp;default:current_timestamp;not null"`
	ImageUsers          []ImageDataAnggota `json:"imageUsers" gorm:"foreignKey:DataUserID"`
}
