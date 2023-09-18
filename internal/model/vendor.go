package model

import (
	"gorm.io/gorm"
	"time"
)

// Vendor struct with fields for vendorname, password, email, and verification
type Vendor struct {
	gorm.Model
	Vendorname string         `gorm:"not null;column:vendorname"`
	Password   string         `gorm:"not null;column:password" json:"-"`
	Email      string         `gorm:"unique;not null;column:email"`
	Status     string         `gorm:"not null;column:status;default:'pending'"` // check with enum vendor status
	VerifiedAt *time.Time     `gorm:"column:verified_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index;column:deleted_at"`
}
