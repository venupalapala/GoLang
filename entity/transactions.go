package entity

import (
	"time"
)

type Transactions struct {
	TxId         string    `db:"tx_id"`
	AccId        string    `db:"acc_id"`
	TxTs         time.Time `db:"tx_ts"`
	Status       string    `db:"status"`
	Amount       uint      `db:"amount"`
	Merchantname string    `db:"merchantname"`
	MerchantId   string    `db:"merchant_id"`
	TxType       string    `db:"tx_type"`
	TxDetails    string    `db:"tx_details"`
}
