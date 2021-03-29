package transactions

import (
	usersModel "backend/models/users"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID      uint            `json:"userID" form:"userID"`
	User        usersModel.User `json:"user"`
	CategoryID  uint            `json:"categoryID" form:"categoryID"`
	Category    Category        `json:"category"`
	Amount      uint            `json:"amount" form:"amount"`
	Description string          `json:"description" form:"description" gorm:"size:255"`
	Type        bool            `json:"type" form:"type"`
	ReceiptPath string          `json:"receiptPath" form:"receiptPath"`
}
