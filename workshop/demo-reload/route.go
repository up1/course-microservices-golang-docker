package demo

import "github.com/gofiber/fiber/v2"

func demo(c *fiber.Ctx) error {
	return c.SendString("Hello..., World 👋!")
}
