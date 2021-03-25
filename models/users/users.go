package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
	EmailVerified bool `gorm:"default:0"`
}
