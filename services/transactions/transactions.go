package transactions

import (
	transactionsDBInteractions "backend/database/transactions"
	paginationData "backend/models/pagination"
	transactionsModel "backend/models/transactions"
)

func CreateTransaction(transaction *transactionsModel.Transaction) error {
	if *transaction.CategoryID == 0 {
		transaction.CategoryID = nil
	}
	return transactionsDBInteractions.CreateTransaction(transaction)
}

func GetTransactionsByUser(data *paginationData.Data, userID uint) ([]transactionsModel.Transaction, int64) {
	return transactionsDBInteractions.GetTransactionsByUser(data, userID)
}

func DeleteTransaction(transactionID, userID uint) error {
	return transactionsDBInteractions.DeleteTransaction(transactionID, userID)
}