package transactions

import (
	dbInstance "backend/database"
	customDatabaseErrors "backend/database/errors"
	transactionsModel "backend/models/transactions"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

var createdTransactionID uint

func TestCreateTransaction(t *testing.T) {
	var transaction transactionsModel.Transaction
	err := dbInstance.GetDBConnection().Where("user_id = 27 AND title = 'Testing transactions'").First(&transaction).Error
	assert.Equal(t, transaction.ID, uint(0))
	assert.Error(t, err, gorm.ErrRecordNotFound)
	currentTime := time.Now()

	transaction = transactionsModel.Transaction{
		UserID:     27,
		CategoryID: 95,
		Amount:     250,
		Title:      "Testing transactions",
		Type:       false,
		PaidAt:     currentTime,
	}
	err = CreateTransaction(&transaction)
	createdCategoryID = transaction.ID

	transaction = transactionsModel.Transaction{}
	err = dbInstance.GetDBConnection().Where("user_id = 27 AND title = 'Testing transactions'").First(&transaction).Error
	assert.Equal(t, err, nil)
	assert.Equal(t, transaction.UserID, uint(27))
	assert.Equal(t, transaction.CategoryID, uint(95))
	assert.Equal(t, transaction.Amount, uint(250))
	assert.Equal(t, transaction.Title, "Testing transactions")
	assert.Equal(t, transaction.Type, false)
	assert.Equal(t, transaction.PaidAt.Unix(), currentTime.Unix())
}

func TestDeleteTransaction(t *testing.T) {
	var transaction transactionsModel.Transaction
	err := dbInstance.GetDBConnection().Where("user_id = 27 AND title = 'Testing transactions'").First(&transaction).Error
	assert.Equal(t, err, nil)

	transactionID := transaction.ID

	err = DeleteTransaction(transactionID, 27)
	assert.Equal(t, err, nil)

	err = DeleteTransaction(transactionID, 27)
	assert.Equal(t, err, &customDatabaseErrors.RecordNotFoundError{})

	transaction = transactionsModel.Transaction{}
	err = dbInstance.GetDBConnection().Where("user_id = 27 AND title = 'Testing transactions'").First(&transaction).Error
	assert.Equal(t, err, gorm.ErrRecordNotFound)
}
