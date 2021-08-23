package wallet

import (
	"os"
	"fmt"
	"github.com/foxnut/go-hdwallet"
)

func CreateWallet(crypto uint32) {
	master, err := hdwallet.NewKey(
        hdwallet.Mnemonic(os.Getenv("SEED")),
    )
    if err != nil {
        panic(err)
    }

	wallet, _ := master.GetWallet(hdwallet.CoinType(crypto))
    address, _ := wallet.GetAddress()
	
	fmt.Println(address, wallet.GetKey().PrivateHex())
}