package model

import (
	"gorm.io/gorm"
)

// User struct with fields for username, password, phone, and verification
type User struct {
	gorm.Model
	Username  string         `gorm:"unique;not null;column:username"`
	Password  string         `gorm:"not null;column:password"`
	Phone     string         `gorm:"unique;not null;column:phone"`
	Session   string         `gorm:"column:session"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}
