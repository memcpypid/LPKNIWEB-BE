package models

import (
	"time"
)

type PengajuanSponsor struct {
	ID              uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	ApplicantName   string     `gorm:"size:100;not null" json:"applicant_name"`
	ApplicantEmail  string     `gorm:"size:100;not null" json:"applicant_email"`
	PhoneNumber     string     `gorm:"size:15;not null" json:"phone_number"`
	SponsorName     string     `gorm:"size:100;not null" json:"sponsor_name"`
	AmountRequested float64    `gorm:"not null" json:"amount_requested"`
	SponsorshipType string     `gorm:"size:50;not null" json:"sponsorship_type"`
	Justification   string     `gorm:"type:text;not null" json:"justification"`
	Status          string     `gorm:"type:enum('pending', 'approved', 'rejected');default:'pending';not null" json:"status"`
	StartDate       *time.Time `json:"start_date,omitempty"`
	EndDate         *time.Time `json:"end_date,omitempty"`
	Attachment      string     `gorm:"size:255" json:"attachment,omitempty"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
