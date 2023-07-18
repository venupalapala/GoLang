package service

import (
	"go-app/team1/entity"
	"go-app/team1/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionsService interface {
	GetTransactionsByAccountId(accountId string, status string, ctx *gin.Context) []entity.Transactions
	GetTransactionsByAccountIdFromDateTimeAndStatus(accountId string, dateTime time.Time, status string, ctx *gin.Context) []entity.Transactions
	GetTransacationsByAccountIdBetweenDateTime(accountId string, fromDateTime time.Time, toDateTime time.Time, status string, ctx *gin.Context) []entity.Transactions
	GetTransactionByTransactionId(transactionId string) entity.Transactions
}

type service struct {
	repository repository.TransactionsRepository
}

func New(repository repository.TransactionsRepository) TransactionsService {
	return &service{
		repository: repository,
	}
}

func (service *service) GetTransactionsByAccountId(accountId string, status string, ctx *gin.Context) []entity.Transactions {
	var transactions []entity.Transactions
	if len(status) > 0 {
		transactions = service.repository.GetTransactionsByAccountIdAndStatus(accountId, status, ctx)
	} else {
		transactions = service.repository.GetTransactionsByAccountId(accountId, ctx)
	}
	return transactions
}

func (service *service) GetTransactionByTransactionId(transactionId string) (transactions entity.Transactions) {
	var transaction entity.Transactions
	transaction = service.repository.GetTransactionByTransactionId(transactionId)
	return transaction
}

func (service *service) GetTransactionsByAccountIdFromDateTimeAndStatus(accountId string, dateTime time.Time, status string, ctx *gin.Context) []entity.Transactions {
	var transactions []entity.Transactions
	if len(status) > 0 {
		transactions = service.repository.GetTransactionsByAccountIdFromParticularDateAndStatus(accountId, dateTime, status, ctx)
	} else {
		transactions = service.repository.GetTransactionsByAccountIdFromParticularDate(accountId, dateTime, ctx)
	}
	return transactions
}

func (service *service) GetTransacationsByAccountIdBetweenDateTime(accountId string, fromDateTime time.Time, toDateTime time.Time, status string, ctx *gin.Context) []entity.Transactions {
	var transactions []entity.Transactions
	if len(status) > 0 {
		transactions = service.repository.GetTransacationsByAccountIdBetweenDateTimeAndStatus(accountId, fromDateTime, toDateTime, status, ctx)
	} else {
		transactions = service.repository.GetTransacationsByAccountIdBetweenDateTime(accountId, fromDateTime, toDateTime, ctx)
	}
	return transactions
}
