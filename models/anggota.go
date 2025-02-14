package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Anggota struct {
	AnggotaID uint `json:"anggota_id" gorm:"primaryKey"`
	// DaerahID     *uint      `json:"daerah_id" gorm:"index;null"`
	// Daerah       Daerah     `json:"daerah" gorm:"foreignKey:DaerahID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL"`
	NamaAnggota    string     `gorm:"size:50;not null" json:"nama_anggota"`
	EmailAnggota   string     `gorm:"size:50;not null;unique" json:"email_anggota"`
	Password       string     `gorm:"size:255;not null" json:"password"`
	NoHp           string     `gorm:"size:15;not null" json:"no_hp"`
	Provinsi       string     `gorm:"type:text" json:"provinsi"`
	TempatLahir    string     `gorm:"size:50" json:"tempat_lahir"`
	TanggalLahir   string     `gorm:"not null" json:"tanggal_lahir"`
	JenisKelamin   string     `gorm:"type:enum('laki-laki','perempuan');default:'laki-laki';not null" json:"jenis_kelamin"`
	NomorKTP       string     `gorm:"size:20" json:"nomor_ktp"`
	Pekerjaan      string     `gorm:"size:50" json:"pekerjaan"`
	SuratIjinAkses string     `gorm:"size:255" json:"surat_ijin_akses"`
	CreatedAt      time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	Keuangans      []Keuangan `gorm:"foreignKey:AnggotaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"keuangans"`
}

func (a *Anggota) BeforeCreate(tx *gorm.DB) (err error) {

	if a.Password != "" {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		a.Password = string(hashedPassword)
	}

	return nil
}

func (a *Anggota) VerifikasiPassword(password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	return err == nil
}
