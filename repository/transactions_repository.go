package repository

import (
	"go-app/team1/entity"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type TransactionsRepository interface {
	GetTransactionsByAccountId(accountId string, ctx *gin.Context) []entity.Transactions
	GetTransactionsByAccountIdAndStatus(accountId string, status string, ctx *gin.Context) []entity.Transactions
	GetTransactionsByAccountIdFromParticularDate(accountId string, dateTime time.Time, ctx *gin.Context) (transactions []entity.Transactions)
	GetTransactionsByAccountIdFromParticularDateAndStatus(accountId string, dateTime time.Time, status string, ctx *gin.Context) (transactions []entity.Transactions)
	GetTransacationsByAccountIdBetweenDateTime(accountId string, fromDateTime time.Time, toDateTime time.Time, ctx *gin.Context) (transactions []entity.Transactions)
	GetTransacationsByAccountIdBetweenDateTimeAndStatus(accountId string, fromDateTime time.Time, toDateTime time.Time, status string, ctx *gin.Context) (transactions []entity.Transactions)

	GetTransactionByTransactionId(transactionId string) (transaction entity.Transactions)
}

type transactionsRepository struct {
	transactions []entity.Transactions
}

func New() TransactionsRepository {
	return &transactionsRepository{}
}

var dsn = "host=pgdb user=postgres password=postgres dbname=workshop sslmode=disable"

func (repository *transactionsRepository) GetTransactionsByAccountId(accountId string, ctx *gin.Context) []entity.Transactions {

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "transaction.",
			SingularTable: false,
		},
	})

	if err != nil {
		panic(err.Error())
	}

	pageSizeStr := ctx.Query("pageSize")
	pageSize := 50
	if pageSizeStr != "" {
		convPageSize, _ := strconv.Atoi(pageSizeStr)
		pageSize = convPageSize
	}
	pageStr := ctx.Query("page")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * pageSize

	var transactionRows []entity.Transactions
	db.Where("acc_id = ?", accountId).Limit(pageSize).Offset(offset).Find(&transactionRows)

	return transactionRows
}

func (repository *transactionsRepository) GetTransactionsByAccountIdAndStatus(accountId string, status string, ctx *gin.Context) []entity.Transactions {

	// dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "transaction.",
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	pageSizeStr := ctx.Query("pageSize")
	pageSize := 50
	if pageSizeStr != "" {
		convPageSize, _ := strconv.Atoi(pageSizeStr)
		pageSize = convPageSize
	}
	pageStr := ctx.Query("page")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * pageSize

	var transactionRows []entity.Transactions
	db.Where("acc_id = ? and status = ?", accountId, status).Limit(pageSize).Offset(offset).Find(&transactionRows)

	return transactionRows
}

func (repository *transactionsRepository) GetTransactionsByAccountIdFromParticularDate(accountId string, dateTime time.Time, ctx *gin.Context) []entity.Transactions {

	// dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "transaction.",
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	pageSizeStr := ctx.Query("pageSize")
	pageSize := 50
	if pageSizeStr != "" {
		convPageSize, _ := strconv.Atoi(pageSizeStr)
		pageSize = convPageSize
	}
	pageStr := ctx.Query("page")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * pageSize

	var transactionRows []entity.Transactions
	db.Where("acc_id = ? and tx_ts >= ?", accountId, dateTime).Limit(pageSize).Offset(offset).Find(&transactionRows)

	return transactionRows
}

func (repository *transactionsRepository) GetTransactionsByAccountIdFromParticularDateAndStatus(accountId string, dateTime time.Time, status string, ctx *gin.Context) []entity.Transactions {

	// dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "transaction.",
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	pageSizeStr := ctx.Query("pageSize")
	pageSize := 50
	if pageSizeStr != "" {
		convPageSize, _ := strconv.Atoi(pageSizeStr)
		pageSize = convPageSize
	}
	pageStr := ctx.Query("page")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * pageSize

	var transactionRows []entity.Transactions
	db.Where("acc_id = ? and tx_ts >= ? and status =?", accountId, dateTime, status).Limit(pageSize).Offset(offset).Find(&transactionRows)

	return transactionRows
}

func (repository *transactionsRepository) GetTransacationsByAccountIdBetweenDateTime(accountId string, fromDateTime time.Time, toDateTime time.Time, ctx *gin.Context) []entity.Transactions {

	// dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "transaction.",
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	pageSizeStr := ctx.Query("pageSize")
	pageSize := 50
	if pageSizeStr != "" {
		convPageSize, _ := strconv.Atoi(pageSizeStr)
		pageSize = convPageSize
	}
	pageStr := ctx.Query("page")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * pageSize

	var transactionRows []entity.Transactions
	db.Where("acc_id = ? and tx_ts >= ? and tx_ts <= ? ", accountId, fromDateTime, toDateTime).Limit(pageSize).Offset(offset).Find(&transactionRows)

	return transactionRows
}

func (repository *transactionsRepository) GetTransacationsByAccountIdBetweenDateTimeAndStatus(accountId string, fromDateTime time.Time, toDateTime time.Time, status string, ctx *gin.Context) []entity.Transactions {

	// dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "transaction.",
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	pageSizeStr := ctx.Query("pageSize")
	pageSize := 50
	if pageSizeStr != "" {
		convPageSize, _ := strconv.Atoi(pageSizeStr)
		pageSize = convPageSize
	}
	pageStr := ctx.Query("page")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * pageSize

	var transactionRows []entity.Transactions
	db.Where("acc_id = ? and tx_ts >= ? and tx_ts <= ? and status =?", accountId, fromDateTime, toDateTime, status).Limit(pageSize).Offset(offset).Find(&transactionRows)

	return transactionRows
}

func (repository *transactionsRepository) GetTransactionByTransactionId(transactionId string) entity.Transactions {

	// dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "transaction.",
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err.Error())
	}

	var transactionRow entity.Transactions
	db.Where("tx_id = ?", transactionId).Find(&transactionRow)

	return transactionRow
}
