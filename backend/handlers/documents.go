package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/i-use-mint-btw/storage"
)

type Document struct {
	Title string `json:"title"`
}

func CreateDocument(c *fiber.Ctx) error {
	document := new(Document)
	err := c.BodyParser(document)

	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error": "Failed to parse json",
		})
	}

	row := storage.DB.QueryRow("INSERT INTO documents (title) VALUES ($1) RETURNING id", document.Title)
	var id string
	err = row.Scan(&id)

	if err != nil {
		log.Println(err)
		c.JSON(&fiber.Map{
			"success": false,
			"error": "Internal",
		})
	}

	log.Print(id)

	return c.JSON(&fiber.Map{
		"success": true,
		"id": id,
	})
}

func EditDocument(c *websocket.Conn) {
	documentID := c.Params("id")

	log.Println("Connection from: ", c.RemoteAddr())

	for {
		mt, msg, err := c.ReadMessage()

		if err != nil {
			log.Print("Error when reading message", msg)
			break
		}

		if mt == websocket.CloseMessage {
			c.Close()
			break
		}

		_, err = storage.DB.Exec("UPDATE documents SET content=$1 WHERE id=$2", msg, documentID)

		if err != nil {
			log.Print("Failed to update document")
		}

		row := storage.DB.QueryRow("SELECT content FROM documents WHERE id=$1", documentID)

		if row.Err() == sql.ErrNoRows {
			log.Print(err)
			c.Close()
			break
		}

		

		c.WriteMessage(websocket.TextMessage, []byte(msg))
	}
}

/*
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
*/
