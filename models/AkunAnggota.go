package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AkunAnggota struct {
	ID           uint       `json:"id_user" gorm:"primaryKey;column:id_user"`
	Email        string     `json:"email" gorm:"type:varchar(100);unique;not null"` // Email untuk login
	Password     string     `json:"password" gorm:"type:varchar(255);not null"`     // Password terenkripsi
	NamaDepan    string     `json:"nama_depan" gorm:"type:varchar(255);not null"`
	NamaBelakang string     `json:"nama_belakang" gorm:"type:varchar(255)"`
	NoHp         string     `json:"no_hp" gorm:"type:varchar(255);unique;not null"`
	Role         string     `json:"role" gorm:"type:varchar(20);not null"` // Role pengguna: 'Admin', 'User', 'Manager', dll.
	LastLogin    *time.Time `json:"lastLogin,omitempty"`                   // Waktu terakhir login
	CreatedAt    time.Time  `json:"createdAt" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"type:timestamp;default:current_timestamp"`
	// Relasi One-to-One dengan DataUser
	// DataUser DataAnggota `json:"dataUser" gorm:"foreignKey:UserID;references:ID"`
}

func (AkunAnggota) TableName() string {
	return "akun_anggota" // New table name
}
func (user *AkunAnggota) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *AkunAnggota) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
