package main

import (
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/i-use-mint-btw/concurrent"
	"github.com/i-use-mint-btw/globals"
	"github.com/i-use-mint-btw/handlers"
	"github.com/i-use-mint-btw/middleware"
	"github.com/i-use-mint-btw/storage"
)

func main() {
	err := storage.InitDB()

	if err != nil {
		log.Fatal("Failed to initialize database. err: ", err)
	}

	app := fiber.New()
	globals.GlobalHub = concurrent.NewHub()

	go globals.GlobalHub.Run()
	go globals.GlobalHub.DatabaseManager.Run()

	app.Use(cors.New(cors.Config{AllowOrigins: os.Getenv("ALLOWED_ORIGINS")}))
	app.Use("/api/document/edit", middleware.EnforceWebsocketConnection)
	app.Post("/api/document/create", handlers.PostDocument)
	app.Get("/api/document/edit/:id", websocket.New(handlers.WebsocketHandler))
	app.Get("/api/document/read/:id", handlers.GetDocument)

	app.Get("/api/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the home route")
	})

	log.Fatal(app.Listen(":2680"))
}
