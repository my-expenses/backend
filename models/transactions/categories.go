package transactions

import (
	usersModel "backend/models/users"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	UserID      uint            `json:"userID" form:"userID" gorm:"uniqueIndex:idx_user_category"`
	User        usersModel.User `json:"user"`
	Title       string          `json:"title" form:"title" gorm:"uniqueIndex:idx_user_category;size:50"`
	Description string          `json:"description" form:"description" gorm:"size:255"`
}
