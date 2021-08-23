package database

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Currency    string // e.g. ETH, BTC, etc.
	ToAddress   string
	FromAddress string
	Amount      float32
	TxID        string // transaction hash on blockchain
	Metadata    string // stringified JSON for whatever the merchant's doing
	Timestamp   time.Time
}

type Wallet struct {
	Currency   string
	PublicKey  string
	PrivateKey string
}

func Connect() gorm.DB {
	db, err := gorm.Open(sqlite.Open("./.store.db"))

	if err != nil {
		panic("failed to connect to db")
	}

	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&Wallet{})

	return *db
}
