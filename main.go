package main

import (
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/foxnut/go-hdwallet"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"blockchain-listener-service/routes"
	"blockchain-listener-service/wallet"
	"blockchain-listener-service/database"
)

func main() {
	godotenv.Load()
	database.Connect()

	wallet.CreateWallet(hdwallet.ETH)

	app := fiber.New(fiber.Config{ Prefork: true })
	app.Use(recover.New())

    app.Post("/create-payment-session", routes.CreatePaymentSession)
	
    app.Listen("localhost:1337")
}