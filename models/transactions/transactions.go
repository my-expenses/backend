package transactions

import (
	usersModel "backend/models/users"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	UserID      uint            `json:"userID" form:"userID"`
	User        usersModel.User `json:"user"`
	CategoryID  uint            `json:"categoryID" form:"categoryID"`
	Category    Category        `json:"category"`
	Amount      uint            `json:"amount" form:"amount" validator:"required"`
	Title       string          `json:"title" form:"title" gorm:"size:255" validator:"required"`
	Type        bool            `json:"type" form:"type" validator:"required"`
	ReceiptPath string          `json:"receiptPath" form:"receiptPath"`
	PaidAt      time.Time       `json:"paidAt" form:"paidAt" validation:"required,lte"`
}
