package main

import (
	"blockchain-listener-service/database"
	"blockchain-listener-service/routes"
	"blockchain-listener-service/wallet"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	database.Connect("./.store.db")
	wallet.InitializeWallet()

	app := fiber.New(fiber.Config{Prefork: true})
	app.Use(recover.New())

	app.Post("/create-payment-session", routes.CreatePaymentSession)

	app.Listen("localhost:1337")
}
