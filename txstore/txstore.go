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
	Timestamp   time.Time
}

type TxStore struct {
	db gorm.DB
}

func Create() TxStore {
	db, err := gorm.Open(sqlite.Open("./.store.db"))
	if err != nil {
		panic("failed to connect to db")
	}
	db.AutoMigrate(&Transaction{})
	t := TxStore{db: *db}
	return t
}

func (t TxStore) AddTransaction(curr string, toAddr string, fromAddr string, amt float32, tm time.Time) {
	t.db.Select(curr, toAddr, fromAddr, amt, tm)
}

func (t TxStore) Clear() {
	t.db.Delete(Transaction{})
}
