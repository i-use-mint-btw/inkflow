package concurrent

import (
	"fmt"
	"log"

)

type Broadcast struct {
	client *Client // The client who actually broadcasted this message
	message []byte
}

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan *Broadcast
	Register   chan *Client
	Unregister chan *Client
	DatabaseManager *DatabaseManager
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan *Broadcast),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		DatabaseManager: NewDatabaseManager(),
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
		case broadcast := <-h.Broadcast:
			for client := range h.Clients {
				if broadcast.client == client {
					continue
				}

				select {
				case client.Send <- broadcast:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}

			dbMessage := &DatabaseMessage{
				documentUpdate: &DocumentUpdate{
					documentID: broadcast.client.Conn.Params("id"),
					content: broadcast.message,
				},
			}

			h.DatabaseManager.Incoming <- dbMessage
		}
	}
}
