package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ziyacivan/go-easy-qr/qrcode"
)

func GeneralRoute(app *fiber.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Easy QR API",
			"routes": fiber.Map{
				"generate": c.BaseURL() + "/api/v1/generate",
			},
		})
	}
}

func GenerateCode(app *fiber.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		url := c.Query("data")

		if url == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Data must be provided as query parameter"})
		}

		qr, err := qrcode.GenerateQRCode(url)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Type("png").Send(qr)
	}
}
