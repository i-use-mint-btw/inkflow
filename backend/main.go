package main

import (
	"log"
	"net/http"
	"slices"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"

	_ "github.com/lib/pq"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		allowedOrigins := []string{"http://localhost:5500"}
		origin := r.Header.Get("Origin")
		return slices.Contains(allowedOrigins, origin)
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	sock, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal("Failed to upgrade connection")
	}
	defer sock.Close()

	for {
		msgType, msg, err := sock.ReadMessage()

		if err != nil {
			log.Print("Failed to parse message from client")
		}

		if msgType == websocket.CloseMessage {
			log.Print("client disconnected")
		}

		log.Print("Message from client: ", msg)
	}
}

func main() {
	app := fiber.New()

	app.Get("/api/", func (c *fiber.Ctx) error {
		return c.SendString("Welcome to the home route")
	})

	app.Options("/api/document/:id/", adaptor.HTTPHandlerFunc(websocketHandler))

	log.Fatal(app.Listen(":2680"))
}