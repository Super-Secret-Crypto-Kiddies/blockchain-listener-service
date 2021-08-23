package routes

import (
	"blockchain-listener-service/wallet"

	"github.com/foxnut/go-hdwallet"
	"github.com/gofiber/fiber/v2"
)

var CoversionMap = map[string]uint32{
	"BTC":        hdwallet.BTC,
	"BTCTestnet": hdwallet.BTCTestnet,
	"LTC":        hdwallet.LTC,
	"DOGE":       hdwallet.DOGE,
	"DASH":       hdwallet.DASH,
	"ETH":        hdwallet.ETH,
	"ETC":        hdwallet.ETC,
	"BCH":        hdwallet.BCH,
	"QTUM":       hdwallet.QTUM,
}

func CreatePaymentSession(c *fiber.Ctx) error {
	payload := struct {
		Crypto string `json:"crypto"` // Ticker symol of the cryptocurrency, e.g. ETH
		Meta   string `json:"meta"`   // Metadata includes merchant specified data, e.g. userId or productId
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(500) // Return status 500 if the JSON payload is not unserializable
	}

	account := wallet.CreateWallet(CoversionMap[payload.Crypto])

	return c.JSON(&fiber.Map{
		"address": account.PublicAddress,
	})
}

func SpawnNewListener() {

}
