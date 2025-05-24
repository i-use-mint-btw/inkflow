package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/i-use-mint-btw/api"
	"github.com/i-use-mint-btw/concurrent"
	"github.com/i-use-mint-btw/globals"
	"github.com/i-use-mint-btw/services"
)

func WebsocketHandler(c *websocket.Conn) {
	defer func() {
		log.Println("Websocket closed")
	}()

	client := &concurrent.Client{
		Conn: c,
		Send: make(chan *concurrent.Broadcast, 256),
		Hub:  globals.GlobalHub,
	}

	// Try to see if the document already has content in it
	message, err := services.ReadDocument(c.Params("id"))

	if err != nil {
		log.Print("User tried to edit a document that does not exist")
		c.Close()
		return
	}

	err = c.WriteMessage(websocket.TextMessage, []byte(message))

	if err != nil {
		log.Print("Failed to send markdown ")
	}

	globals.GlobalHub.Register <- client

	var wg sync.WaitGroup
	wg.Add(2)
	go client.WritePump(&wg)
	go client.ReadPump(&wg)
	wg.Wait()
}

func PostDocument(c *fiber.Ctx) error {
	var document api.CreateDocumentDTO
	err := c.BodyParser(&document)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.NewFailureResponse("Poorly formatted request"))
	}

	documentID, err := services.CreateDocument(document.Title)

	if err != nil {
		log.Println("Failed to create document: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(api.NewFailureResponse("Internal server error"))
	}

	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"success": true,
		"message": "Document Created successfully",
		"data": fiber.Map{
			"id":    documentID,
			"title": document.Title,
		},
	})
}
