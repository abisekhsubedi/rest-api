package handlers

import "github.com/gofiber/fiber/v2"

func HTMLHandler(c *fiber.Ctx) error {
	return c.SendFile("../index.html")
}
	