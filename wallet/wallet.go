package wallet

import (
	"blockchain-listener-service/database"

	"github.com/dongri/go-mnemonic"
	"github.com/foxnut/go-hdwallet"
)

var MasterKey *hdwallet.Key

func InitializeWallet() {
	seedPhrase := database.SeedPhrase{Seed: CreateSeedPhrase()}
	database.DB.FirstOrCreate(&seedPhrase)

	master, err := hdwallet.NewKey(
		hdwallet.Mnemonic(seedPhrase.Seed),
	)

	if err != nil {
		panic(err)
	} else {
		MasterKey = master
	}
}

func CreateSeedPhrase() string {
	phrase, _ := mnemonic.GenerateMnemonic(256, mnemonic.LanguageEnglish)
	return phrase
}

func CreateWallet(coin uint32) database.Wallet {
	index := database.GetWalletIndex(coin)
	wallet, _ := MasterKey.GetWallet(hdwallet.CoinType(coin), hdwallet.AddressIndex(index))
	address, _ := wallet.GetAddress()
	key := wallet.GetKey().PrivateHex()

	account := database.Wallet{
		Currency:      coin,
		PublicAddress: address,
		PrivateKey:    key,
	}

	database.DB.Create(&account)

	return account
}
