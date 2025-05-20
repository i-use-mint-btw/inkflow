package helpers

import (
	"log"
	"sync"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/i-use-mint-btw/handlers"
)

// Accept an http request
// Open a websocket connection to that client
// Allow the client to send messages through the socket
// Then broadcast the message to all other clients

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
	Hub  *Hub
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Reads messages from the websocket connection and broadcasts them
func (c *Client) ReadPump(wg *sync.WaitGroup) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
		wg.Done()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	documentID := c.Conn.Params("id")

	for {
		_, message, err := c.Conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		err = handlers.UpdateDocument(documentID, message)

		if err != nil {
			log.Printf("Failed to persist client message (%v) in the database", message)
			break
		}

		content, err := handlers.ReadDocument(documentID)

		if err != nil {
			log.Println("Failed to read document")
			break
		}

		c.Hub.Broadcast <- []byte(content)
	}
}

// Reads broadcasted messages and writes them to the websocket
func (c *Client) WritePump(wg *sync.WaitGroup) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
		wg.Done()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// ...existing code...
