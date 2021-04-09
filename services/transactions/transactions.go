package transactions

import (
	transactionsDBInteractions "backend/database/transactions"
	paginationData "backend/models/pagination"
	transactionsModel "backend/models/transactions"
	"github.com/jinzhu/now"
	"time"
)

func CreateTransaction(transaction *transactionsModel.Transaction) error {
	if *transaction.CategoryID == 0 {
		transaction.CategoryID = nil
	}
	return transactionsDBInteractions.CreateTransaction(transaction)
}

func GetTransactionsByUser(data *paginationData.Data, month time.Time, userID uint) ([]transactionsModel.Transaction, int64) {
	startOfMonth := now.With(month).BeginningOfMonth()
	endOfMonth := now.With(month).EndOfMonth()
	return transactionsDBInteractions.GetTransactionsByUser(data, startOfMonth, endOfMonth, userID)
}

func GetGroupedTransactions(userID uint, month time.Time) []map[string]interface{} {
	startOfMonth := now.With(month).BeginningOfMonth()
	endOfMonth := now.With(month).EndOfMonth()
	return transactionsDBInteractions.GetGroupedTransactions(userID, startOfMonth, endOfMonth)
}

func DeleteTransaction(transactionID, userID uint) error {
	return transactionsDBInteractions.DeleteTransaction(transactionID, userID)
}