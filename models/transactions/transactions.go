package transactions

import (
	usersModel "backend/models/users"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	UserID      uint            `json:"userID" form:"userID"`
	User        usersModel.User `json:"user" validate:"required,nostructlevel"`
	CategoryID  *uint           `json:"categoryID" form:"categoryID"`
	Category    Category        `json:"category" validate:"required,nostructlevel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount      int             `json:"amount" form:"amount" validate:"required,gt=0"`
	Title       string          `json:"title" form:"transactionTitle" gorm:"size:255" validate:"required"`
	Type        bool            `json:"type" form:"type"`
	ReceiptPath string          `json:"receiptPath" form:"receiptPath"`
	Date        time.Time       `json:"date" form:"date" validate:"required,lte"`
}
