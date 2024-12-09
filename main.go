package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/ziyacivan/go-easy-qr/router"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", router.GeneralRoute(app))
	v1.Post("/generate", router.GenerateCode(app))

	log.Fatal(app.Listen(":3000"))
}
