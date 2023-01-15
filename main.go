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
		sampleDoc := []interface{}{
			bson.M{"name": "React js book", "completed": false, "type": "book"},
			bson.M{"name": "Javascript the weird part", "completed": false, "type": "video course"},
			bson.M{"name": "laracast PHP course", "completed": false, "type": "video course"},
		}
		collection := database.GetCollection("todos")
		nDoc, err := collection.InsertMany(context.TODO(), sampleDoc)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error inserting todo")
		}
		// send down info about the todo
		return c.JSON(nDoc)
	})

	app.Listen(":3000")
}
