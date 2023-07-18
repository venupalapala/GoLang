package entity

import (
	"time"
)

type Transactions struct {
	TransactionId        string    `db:"tx_id"`
	AccountId            string    `db:"acc_id"`
	TransactionTimeStamp time.Time `db:"tx_ts"`
	Status               string    `db:"status"`
	Amount               uint      `db:"amount"`
	MerchantName         string    `db:"merchantname"`
	MerchantId           string    `db:"merchant_id"`
	TransactionType      string    `db:"tx_type"`
	TransactionDetails   string    `db:"tx_details"`
}
