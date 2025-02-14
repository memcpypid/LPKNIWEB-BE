package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username       string    `gorm:"size:100;not null;unique" json:"username"`
	EmailAdmin     string    `gorm:"size:50;not null;unique" json:"email_admin"`
	Password       string    `gorm:"size:255;not null" json:"password"`
	FirstName      string    `gorm:"size:50;not null" json:"first_name"`
	LastName       string    `gorm:"size:50;not null" json:"last_name"`
	Provinsi       string    `gorm:"size:50;not null" json:"provinsi"`
	NIP            string    `gorm:"size:50;not null" json:"nip"`
	SuratIjinAkses string    `gorm:"size:255" json:"surat_ijin_akses"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// type Wilayah struct {
// 	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
// 	Nama         string    `gorm:"size:100;not null" json:"nama"`
// 	EmailWilayah string    `gorm:"size:50;not null" json:"email_wilayah"`
// 	Password     string    `gorm:"size:255;not null" json:"password"`
// 	AdminID      uint      `gorm:"not null" json:"admin_id"`
// 	Admin        Admin     `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"admin"`
// 	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
// 	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
// 	Daerahs      []Daerah  `gorm:"foreignKey:WilayahID" json:"daerahs"`
// }

// type Daerah struct {
// 	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
// 	Nama        string    `gorm:"size:100;not null" json:"nama"`
// 	EmailDaerah string    `gorm:"size:50;not null" json:"email_daerah"`
// 	WilayahID   uint      `gorm:"not null" json:"wilayah_id"`
// 	Password    string    `gorm:"size:255;not null" json:"password"`
// 	Wilayah     Wilayah   `gorm:"foreignKey:WilayahID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"wilayah"`
// 	Anggota     []Anggota `gorm:"foreignKey:DaerahID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"anggota"`
// 	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
// 	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
// }

func (u *Admin) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// func (u *Wilayah) HashPassword() error {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(hashedPassword)
// 	return nil
// }

// func (u *Daerah) HashPassword() error {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(hashedPassword)
// 	return nil
// }
