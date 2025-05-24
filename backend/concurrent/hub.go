package concurrent

import (
	"fmt"
	"log"
)

type Broadcast struct {
	client  *Client // The client who actually broadcasted this message
	message []byte
}

type Hub struct {
	Clients         map[*Client]bool
	Broadcast       chan *Broadcast
	Register        chan *Client
	Unregister      chan *Client
	DatabaseManager *DatabaseManager
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:       make(chan *Broadcast),
		Register:        make(chan *Client),
		Unregister:      make(chan *Client),
		Clients:         make(map[*Client]bool),
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
			addr := client.Conn.RemoteAddr().String()
			fmt.Printf("Client %v connected!\n", addr[len(addr)-2:])
			h.Clients[client] = true

		case client := <-h.Unregister:
			addr := client.Conn.RemoteAddr().String()
			fmt.Printf("Client %v disconnected!\n", addr[len(addr)-2:])

			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

		case broadcast := <-h.Broadcast:
			// Queues a database update that will happen in the background
			dbMessage := &DatabaseMessage{
				documentUpdate: &DocumentUpdate{
					documentID: broadcast.client.Conn.Params("id"),
					content:    broadcast.message,
				},
			}

			h.DatabaseManager.Incoming <- dbMessage

			if len(h.Clients) <= 1 {
				break
			}

			for client := range h.Clients {
				addr := client.Conn.RemoteAddr().String()

				if broadcast.client == client {
					from := broadcast.client.Conn.RemoteAddr().String()
					log.Printf("From: %v = %v", from[len(from)-2:], string(broadcast.message[:11]))
					continue
				}
				log.Println("To: ", addr[len(addr)-2:])

				select {
				case client.Send <- broadcast:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
