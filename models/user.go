package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	JenisPendaftaran string `json:"jenisPendaftaran" gorm:"type:varchar(20);not null"`
	FirstName        string `json:"firstName" gorm:"type:varchar(50);not null"`
	LastName         string `json:"lastName" gorm:"type:varchar(50);not null"`
	ContactNo        string `json:"contactNo" gorm:"type:varchar(20);not null"`
	EmailUs          string `json:"emailUs" gorm:"type:varchar(100);unique;not null"`
	Username         string `json:"username" gorm:"type:varchar(50);unique;not null"`
	Password         string `json:"password" gorm:"type:varchar(255);not null"`
	Daerah           string `json:"daerah" gorm:"type:varchar(50)"`
	Wilayah          string `json:"wilayah" gorm:"type:varchar(50)"`

}


func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}


func (user *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
