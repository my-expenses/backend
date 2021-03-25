package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `gorm:"not null" form:"fullName"`
	Email string `gorm:"unique;not null" form:"email"`
	EmailVerified bool `gorm:"default:0"`
}
