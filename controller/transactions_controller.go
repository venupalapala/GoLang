package controller

import (
	"go-app/team1/entity"
	"go-app/team1/service"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionsController interface {
	FindTransactionsByAccountId(accountId string, status string, ctx *gin.Context) []entity.Transactions
	FindTransactionsByAccountIdFromDateTimeAndStatus(accountId string, dateTime time.Time, status string, ctx *gin.Context) []entity.Transactions
	FindTransacationsByAccountIdBetweenDateTimeAndStatus(accountId string, fromDateTime time.Time, toDateTime time.Time, status string, ctx *gin.Context) []entity.Transactions
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

func (transactionsController *controller) FindTransactionsByAccountId(accountId string, status string, ctx *gin.Context) []entity.Transactions {

	var transactions []entity.Transactions

	transactions = transactionsController.service.GetTransactionsByAccountId(accountId, status, ctx)

	return transactions

}

func (transactionsController *controller) FindTransactionByTransactionId(transactionId string) (transactions entity.Transactions) {
	var transaction entity.Transactions
	transaction = transactionsController.service.GetTransactionByTransactionId(transactionId)
	return transaction
}

func (transactionsController *controller) FindTransactionsByAccountIdFromDateTimeAndStatus(accountId string, dateTime time.Time, status string, ctx *gin.Context) []entity.Transactions {
	var transactions []entity.Transactions

	transactions = transactionsController.service.GetTransactionsByAccountIdFromDateTimeAndStatus(accountId, dateTime, status, ctx)

	return transactions
}

func (transactionsController *controller) FindTransacationsByAccountIdBetweenDateTimeAndStatus(accountId string, fromDateTime time.Time, toDateTime time.Time, status string, ctx *gin.Context) []entity.Transactions {
	var transactions []entity.Transactions

	transactions = transactionsController.service.GetTransacationsByAccountIdBetweenDateTime(accountId, fromDateTime, toDateTime, status, ctx)

	return transactions
}
