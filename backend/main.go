package main

import (
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/i-use-mint-btw/handlers"
	"github.com/i-use-mint-btw/middleware"
	"github.com/i-use-mint-btw/storage"
)

func main() {
	err := storage.InitDB()

	if err != nil {
		log.Fatal("Failed to initialize database. err: ", err)
	}

	clientURL := os.Getenv("CLIENT_URL")

	app := fiber.New()

	app.Use(cors.New(cors.Config{AllowOrigins: clientURL}))
	app.Use("/api/document/edit", middleware.EnforceWebsocketConnection)

	app.Post("/api/document/create", handlers.CreateDocument)
	app.Get("/api/document/edit/:id", websocket.New(handlers.EditDocument))

	app.Get("/api/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the home route")
	})

	log.Fatal(app.Listen(":2680"))
}

// /api/document/create - Creates a new document and returns the id
// /api/document/:id/edit - Edits a document with a given id
