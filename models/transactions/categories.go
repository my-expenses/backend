package transactions

import (
	usersModel "backend/models/users"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	UserID      uint            `json:"userID" form:"userID" gorm:"uniqueIndex:idx_user_category"`
	User        usersModel.User `json:"user" validate:"required,nostructlevel"`
	Title       string          `json:"title" form:"title" gorm:"uniqueIndex:idx_user_category;size:50" validate:"required,lte=30"`
}
