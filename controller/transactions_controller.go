package controller

import (
	"test/entity"
	"test/service"
	"time"
)

type TransactionsController interface {
	FindTransactionsByAccountId(accountId string, status string) []entity.Transactions
	FindTransactionsByAccountIdFromDateTimeAndStatus(accountId string, dateTime time.Time, status string) []entity.Transactions
	FindTransacationsByAccountIdBetweenDateTimeAndStatus(accountId string, fromDateTime time.Time, toDateTime time.Time, status string) []entity.Transactions
	FindTransactionByTransactionId(transactionId string) (transactions entity.Transactions)
}

type controller struct {
	service service.TransactionsService
}

func New(service service.TransactionsService) TransactionsController {
	return &controller{
		service: service,
	}
}

func (transactionsController *controller) FindTransactionsByAccountId(accountId string, status string) []entity.Transactions {

	var transactions []entity.Transactions

	transactions = transactionsController.service.GetTransactionsByAccountId(accountId, status)

	return transactions

}

func (transactionsController *controller) FindTransactionByTransactionId(transactionId string) (transactions entity.Transactions) {
	var transaction entity.Transactions
	transaction = transactionsController.service.GetTransactionByTransactionId(transactionId)
	return transaction
}

func (transactionsController *controller) FindTransactionsByAccountIdFromDateTimeAndStatus(accountId string, dateTime time.Time, status string) []entity.Transactions {
	var transactions []entity.Transactions

	transactions = transactionsController.service.GetTransactionsByAccountIdFromDateTimeAndStatus(accountId, dateTime, status)

	return transactions
}

func (transactionsController *controller) FindTransacationsByAccountIdBetweenDateTimeAndStatus(accountId string, fromDateTime time.Time, toDateTime time.Time, status string) []entity.Transactions {
	var transactions []entity.Transactions

	transactions = transactionsController.service.GetTransacationsByAccountIdBetweenDateTime(accountId, fromDateTime, toDateTime, status)

	return transactions
}
