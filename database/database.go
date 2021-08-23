package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	Confirmed = iota
	Pending
	Partial
	Failed
)

type Transaction struct {
	gorm.Model
	Currency    uint32 // e.g. ETH, BTC, etc according to hdwallet's enum.
	ToAddress   string
	FromAddress *string
	Amount      *float32
	TxID        *string // transaction hash on blockchain
	Status      uint
	PrivateKey  string
	WalletID    uint
}

type Wallet struct {
	gorm.Model
	Currency      uint32
	PublicAddress string
	PrivateKey    string
	Metadata      *string // stringified JSON for whatever the merchant's doing
	Transactions  []Transaction
}

type SeedPhrase struct {
	gorm.Model
	Seed string
}

type WalletIndex struct {
	gorm.Model
	Currency uint32
	Index    uint64
}

var DB *gorm.DB

func Connect(path string) {
	connection, err := gorm.Open(sqlite.Open(path))

	DB = connection

	if err != nil {
		panic("failed to connect to db")
	}

	DB.AutoMigrate(&Transaction{})
	DB.AutoMigrate(&Wallet{})
	DB.AutoMigrate(&SeedPhrase{})
	DB.AutoMigrate(&WalletIndex{})
}

func GetWalletIndex(currency uint32) uint64 {
	var index uint64
	DB.Where(WalletIndex{Currency: currency}).FirstOrCreate(&index, WalletIndex{Index: 0})
	return index
}
