package models

import "time"

// News represents the structure of a news item
type News struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Date        time.Time `json:"date" gorm:"not null"`
	Category    string    `json:"category" gorm:"type:varchar(100);not null"`
	ImageURL    string    `json:"image_url" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
