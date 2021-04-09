package transactions

import (
	dbInstance "backend/database"
	customDatabaseErrors "backend/database/errors"
	"backend/database/scopes"
	paginationData "backend/models/pagination"
	transactionsModel "backend/models/transactions"
	"gorm.io/gorm"
	"time"
)

func CreateTransaction(transaction *transactionsModel.Transaction) error {
	return dbInstance.GetDBConnection().Create(transaction).Error
}

func GetTransactionsByUser(data *paginationData.Data, startOfMonth, endOfMonth time.Time, userID uint) ([]transactionsModel.Transaction, int64) {
	transactions := make([]transactionsModel.Transaction, 0)
	query := dbInstance.GetDBConnection().Where("user_id = ? AND date BETWEEN ? AND ?", userID, startOfMonth, endOfMonth)
	numberOfRecords := countTransactions(query)
	query.Scopes(scopes.Paginate(data)).Find(&transactions)
	return transactions, numberOfRecords
}

func countTransactions(db *gorm.DB) int64 {
	totalTransactions := int64(0)
	db.Model(&transactionsModel.Transaction{}).Count(&totalTransactions)
	return totalTransactions
}

func DeleteTransaction(transactionID, loggedInUserID uint) error {
	db := dbInstance.GetDBConnection().Unscoped().Where("id = ? AND user_id = ?",
		transactionID, loggedInUserID).Delete(&transactionsModel.Transaction{})
	if db.RowsAffected == 0 {
		return &customDatabaseErrors.RecordNotFoundError{}
	}
	return db.Error
}