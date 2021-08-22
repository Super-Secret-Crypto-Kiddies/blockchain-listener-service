package main

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {

	godotenv.Load()

	wallet, err := hdwallet.NewFromMnemonic(os.Getenv("SEED"))
	if err != nil {
		panic(err)
	}

	fmt.Println(wallet)

    app := fiber.New(fiber.Config{ Prefork: true })

    app.Post("/create-payment-session", func(c *fiber.Ctx) error {

		payload := struct {
			Crypto string `json:"crypto"` // Ticker symol of the cryptocurrency, e.g. ETH
			Meta   string `json:"meta"`   // Metadata includes merchant specified data, e.g. userId or productId
		}{}
	
		if err := c.BodyParser(&payload); err != nil {
			return c.SendStatus(500) // Return status 500 if the JSON payload is not unserializable
		}

		return c.SendStatus(200)
    })

    app.Listen("localhost:1337")
}