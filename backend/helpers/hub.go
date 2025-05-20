package helpers

import (
	"fmt"
	"log"
)

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	defer func() {
		log.Println("Hub closed!")
	}()

	log.Println("Hub opened and running")

	for {
		select {
		case client := <-h.Register:
			fmt.Println("Client connected!")
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				fmt.Println("Client disconnected!")
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			log.Print(string(message))
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
