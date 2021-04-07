package transactions

import (
	dbInstance "backend/database"
	customDatabaseErrors "backend/database/errors"
	"backend/models/pagination"
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
	categoryID := uint(1)

	transaction = transactionsModel.Transaction{
		UserID:     27,
		CategoryID: &categoryID,
		Amount:     250,
		Title:      "Testing transactions",
		Type:       false,
		Date:       currentTime,
	}
	err = CreateTransaction(&transaction)
	createdCategoryID = transaction.ID

	transaction = transactionsModel.Transaction{}
	err = dbInstance.GetDBConnection().Where("user_id = 27 AND title = 'Testing transactions'").First(&transaction).Error
	assert.Equal(t, err, nil)
	assert.Equal(t, transaction.UserID, uint(27))
	assert.Equal(t, transaction.CategoryID, &categoryID)
	assert.Equal(t, transaction.Amount, 250)
	assert.Equal(t, transaction.Title, "Testing transactions")
	assert.Equal(t, transaction.Type, false)
	assert.Equal(t, transaction.Date.Unix(), currentTime.Unix())
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

func TestGetTransactionsByUser(t *testing.T) {
	itemsPerPage := 3
	data := pagination.Data{
		SortDesc:     []bool{true},
		SortBy:       []string{"created_at"},
		Page:         0,
		ItemsPerPage: itemsPerPage,
	}
	transactions, numberOfRecords := GetTransactionsByUser(&data, 27)
	assert.Equal(t, len(transactions), itemsPerPage)
	assert.Equal(t, numberOfRecords, int64(4))

	var manualTransactions []transactionsModel.Transaction
	dbInstance.GetDBConnection().Where("user_id = 27").Order("created_at DESC").
		Limit(itemsPerPage).Find(&manualTransactions)
	for i := 0; i < itemsPerPage; i++ {
		assert.Equal(t, transactions[i].ID, manualTransactions[i].ID)
	}
}
