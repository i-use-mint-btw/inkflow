package main

import (
	"log"
	"os"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/i-use-mint-btw/handlers"
	"github.com/i-use-mint-btw/helpers"
	"github.com/i-use-mint-btw/middleware"
	"github.com/i-use-mint-btw/storage"
)

func main() {
	err := storage.InitDB()

	if err != nil {
		log.Fatal("Failed to initialize database. err: ", err)
	}

	app := fiber.New()
	hub := helpers.NewHub()
	go hub.Run()

	app.Use(cors.New(cors.Config{AllowOrigins: os.Getenv("ALLOWED_ORIGINS")}))
	app.Use("/api/document/edit", middleware.EnforceWebsocketConnection)
	app.Post("/api/document/create", handlers.CreateDocument)
	app.Get("/api/document/edit/:id", websocket.New(func(c *websocket.Conn) {
		defer func() {
			log.Println("Websocket closed")
		}()

		client := &helpers.Client{
			Conn: c,
			Send: make(chan []byte, 256),
			Hub:  hub,
		}

		hub.Register <- client

		var wg sync.WaitGroup
		wg.Add(2)
		go client.WritePump(&wg)
		go client.ReadPump(&wg)
		wg.Wait()
	}))

	app.Get("/api/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the home route")
	})

	log.Fatal(app.Listen(":2680"))
}