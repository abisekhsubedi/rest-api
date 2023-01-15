package main

import (
	"context"

	"github.com/abisekhsubedi/rest-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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

	// routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./index.html")
	})
	app.Post("/api", func(c *fiber.Ctx) error {
		// Write a todo to the application
		sampleDoc := bson.M{"name": "Sample todo"}
		collection := database.GetCollection("todos")
		nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
