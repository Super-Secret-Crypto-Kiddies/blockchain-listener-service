package txstore

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

type Store struct {
	db gorm.DB
}

func Create() Store {
	db, err := gorm.Open(sqlite.Open("./.store.db"))
	if err != nil {
		panic("failed to connect to db")
	}
	db.AutoMigrate(&Transaction{})
	t := Store{db: *db}
	return t
}

func (t Store) AddTransaction(curr string, toAddr string, fromAddr string, amt float32, tm time.Time) {
	t.db.Select(curr, toAddr, fromAddr, amt, tm)
}

func (t Store) Clear() {
	t.db.Delete(Transaction{})
}
