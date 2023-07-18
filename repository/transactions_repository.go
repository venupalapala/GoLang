package repository

import (
	"fmt"
	"go-app/team1/entity"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TransactionsRepository interface {
	GetTransactionsByAccountId(accountId string) []entity.Transactions
	GetTransactionsByAccountIdAndStatus(accountId string, status string) []entity.Transactions
	GetTransactionsByAccountIdFromParticularDate(accountId string, dateTime time.Time) (transactions []entity.Transactions)
	GetTransactionsByAccountIdFromParticularDateAndStatus(accountId string, dateTime time.Time, status string) (transactions []entity.Transactions)
	GetTransacationsByAccountIdBetweenDateTime(accountId string, fromDateTime time.Time, toDateTime time.Time) (transactions []entity.Transactions)
	GetTransacationsByAccountIdBetweenDateTimeAndStatus(accountId string, fromDateTime time.Time, toDateTime time.Time, status string) (transactions []entity.Transactions)

	GetTransactionByTransactionId(transactionId string) (transaction entity.Transactions)
}

type transactionsRepository struct {
	transactions []entity.Transactions
}

func New() TransactionsRepository {
	return &transactionsRepository{}
}

func (repository *transactionsRepository) GetTransactionsByAccountId(accountId string) []entity.Transactions {

	dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	fmt.Printf("Value of Account id  %s", accountId)

	rows, err := sqlDB.Query("SELECT * FROM transaction.transactions WHERE acc_id = $1", accountId)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// An Transactions slice to hold data from returned rows.
	var transactionRows []entity.Transactions

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var transaction entity.Transactions
		if err := rows.Scan(&transaction.TransactionId, &transaction.AccountId, &transaction.TransactionTimeStamp,
			&transaction.Status, &transaction.Amount, &transaction.MerchantName, &transaction.MerchantId,
			&transaction.TransactionType, &transaction.TransactionDetails); err != nil {
			return transactionRows
		}
		transactionRows = append(transactionRows, transaction)
	}
	if err = rows.Err(); err != nil {
		return transactionRows
	}
	return transactionRows
}

func (repository *transactionsRepository) GetTransactionsByAccountIdAndStatus(accountId string, status string) []entity.Transactions {

	dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	fmt.Printf("Value of Account id  %s", accountId)

	rows, err := sqlDB.Query("SELECT * FROM transaction.transactions WHERE acc_id = $1 and status = $2", accountId, status)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// An Transactions slice to hold data from returned rows.
	var transactionRows []entity.Transactions

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var transaction entity.Transactions
		if err := rows.Scan(&transaction.TransactionId, &transaction.AccountId, &transaction.TransactionTimeStamp,
			&transaction.Status, &transaction.Amount, &transaction.MerchantName, &transaction.MerchantId,
			&transaction.TransactionType, &transaction.TransactionDetails); err != nil {
			return transactionRows
		}
		transactionRows = append(transactionRows, transaction)
	}
	if err = rows.Err(); err != nil {
		return transactionRows
	}
	return transactionRows
}

func (repository *transactionsRepository) GetTransactionsByAccountIdFromParticularDate(accountId string, dateTime time.Time) []entity.Transactions {

	dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	fmt.Printf("Value of Account id  %s", accountId)

	rows, err := sqlDB.Query("SELECT * FROM transaction.transactions WHERE acc_id = $1 and tx_ts >= $2", accountId, dateTime)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// An Transactions slice to hold data from returned rows.
	var transactionRows []entity.Transactions

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var transaction entity.Transactions
		if err := rows.Scan(&transaction.TransactionId, &transaction.AccountId, &transaction.TransactionTimeStamp,
			&transaction.Status, &transaction.Amount, &transaction.MerchantName, &transaction.MerchantId,
			&transaction.TransactionType, &transaction.TransactionDetails); err != nil {
			return transactionRows
		}
		transactionRows = append(transactionRows, transaction)
	}
	if err = rows.Err(); err != nil {
		return transactionRows
	}
	return transactionRows
}

func (repository *transactionsRepository) GetTransactionsByAccountIdFromParticularDateAndStatus(accountId string, dateTime time.Time, status string) []entity.Transactions {

	dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()

	rows, err := sqlDB.Query("SELECT * FROM transaction.transactions WHERE acc_id = $1 and tx_ts >= $2 and status = $3", accountId, dateTime, status)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// An Transactions slice to hold data from returned rows.
	var transactionRows []entity.Transactions

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var transaction entity.Transactions
		if err := rows.Scan(&transaction.TransactionId, &transaction.AccountId, &transaction.TransactionTimeStamp,
			&transaction.Status, &transaction.Amount, &transaction.MerchantName, &transaction.MerchantId,
			&transaction.TransactionType, &transaction.TransactionDetails); err != nil {
			return transactionRows
		}
		transactionRows = append(transactionRows, transaction)
	}
	if err = rows.Err(); err != nil {
		return transactionRows
	}
	return transactionRows
}

func (repository *transactionsRepository) GetTransacationsByAccountIdBetweenDateTime(accountId string, fromDateTime time.Time, toDateTime time.Time) []entity.Transactions {

	dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()

	rows, err := sqlDB.Query("SELECT * FROM transaction.transactions WHERE acc_id = $1 and tx_ts >= $2 and tx_ts < $3 ", accountId, fromDateTime, toDateTime)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// An Transactions slice to hold data from returned rows.
	var transactionRows []entity.Transactions

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var transaction entity.Transactions
		if err := rows.Scan(&transaction.TransactionId, &transaction.AccountId, &transaction.TransactionTimeStamp,
			&transaction.Status, &transaction.Amount, &transaction.MerchantName, &transaction.MerchantId,
			&transaction.TransactionType, &transaction.TransactionDetails); err != nil {
			return transactionRows
		}
		transactionRows = append(transactionRows, transaction)
	}
	if err = rows.Err(); err != nil {
		return transactionRows
	}
	return transactionRows
}

func (repository *transactionsRepository) GetTransacationsByAccountIdBetweenDateTimeAndStatus(accountId string, fromDateTime time.Time, toDateTime time.Time, status string) []entity.Transactions {

	dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()

	rows, err := sqlDB.Query("SELECT * FROM transaction.transactions WHERE acc_id = $1 and  tx_ts >= $2 and tx_ts < $3 and status = $4", accountId, fromDateTime, toDateTime, status)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// An Transactions slice to hold data from returned rows.
	var transactionRows []entity.Transactions

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var transaction entity.Transactions
		if err := rows.Scan(&transaction.TransactionId, &transaction.AccountId, &transaction.TransactionTimeStamp,
			&transaction.Status, &transaction.Amount, &transaction.MerchantName, &transaction.MerchantId,
			&transaction.TransactionType, &transaction.TransactionDetails); err != nil {
			return transactionRows
		}
		transactionRows = append(transactionRows, transaction)
	}
	if err = rows.Err(); err != nil {
		return transactionRows
	}
	return transactionRows
}

func (repository *transactionsRepository) GetTransactionByTransactionId(transactionId string) entity.Transactions {

	dsn := "host=localhost user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	fmt.Printf("Transaction id : %s", transactionId)
	rows, err := sqlDB.Query("SELECT * FROM transaction.transactions WHERE tx_id = $1 ", transactionId)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	// An Transactions slice to hold data from returned rows.
	var transactionRow entity.Transactions
	for rows.Next() {
		// Loop through rows, using Scan to assign column data to struct fields.
		var transaction entity.Transactions
		if err := rows.Scan(&transaction.TransactionId, &transaction.AccountId, &transaction.TransactionTimeStamp,
			&transaction.Status, &transaction.Amount, &transaction.MerchantName, &transaction.MerchantId,
			&transaction.TransactionType, &transaction.TransactionDetails); err != nil {
			return transactionRow
		}
		transactionRow = transaction
	}
	if err = rows.Err(); err != nil {
		return transactionRow
	}
	return transactionRow
}
