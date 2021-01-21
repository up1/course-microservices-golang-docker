package demo

import "github.com/gofiber/fiber/v2"

func Start() {
	app := fiber.New()

	app.Get("/", demo)

	app.Listen(":3000")
}
