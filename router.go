package main

import (
	"github.com/abisekhsubedi/rest-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func generateApp() *fiber.App {
	app := fiber.New()

	// create a health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK 👌🏼")
	})

	// create the library groups and routes
	libGroup := app.Group("/library")
	libGroup.Get("/", handlers.TestHandler)
	libGroup.Get("/html", handlers.HTMLHandler)

	return app
}
