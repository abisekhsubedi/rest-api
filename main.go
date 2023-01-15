package main

import (
	"fmt"

	"github.com/abisekhsubedi/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// loadEnv loads the .env file
func loadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

// initApp initializes the app
func initApp() error {
	err := loadENV()
	if err != nil {
		return err
	}

	// start mongodb database
	err = database.StartMongoDB()
	if err != nil {
		return err
	}
	return nil
}

func main() {

	// initialize app
	err := initApp()
	if err != nil {
		panic(err)
	}

	// defer database connection
	defer database.CloseMongoDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./index.html")
	})
	app.Listen(":3000")
}
