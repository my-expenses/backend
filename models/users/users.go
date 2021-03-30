package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName      string `gorm:"not null" form:"fullName" validate:"required"`
	Email         string `gorm:"unique;not null" form:"email" validate:"required,email"`
	EmailVerified bool   `gorm:"default:0"`
}
