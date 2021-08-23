package routes

import (
	"github.com/gofiber/fiber/v2"
)

func CreatePaymentSession(c *fiber.Ctx) error {
	payload := struct {
		Crypto string `json:"crypto"` // Ticker symol of the cryptocurrency, e.g. ETH
		Meta   string `json:"meta"`   // Metadata includes merchant specified data, e.g. userId or productId
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(500) // Return status 500 if the JSON payload is not unserializable
	}

	

	return c.SendStatus(200)
}