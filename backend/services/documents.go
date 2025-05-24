package services

import (
	"log"

	"github.com/i-use-mint-btw/storage"
)

func CreateDocument(title string) (string, error) {
	row := storage.DB.QueryRow("INSERT INTO documents (title) VALUES ($1) RETURNING id", title)

	var id string
	err := row.Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdateDocument(id string, content []byte) error {
	_, err := storage.DB.Exec("UPDATE documents SET content=$1 WHERE id=$2", content, id)

	if err != nil {
		log.Print("Failed to update document")
		return err
	}

	return nil
}

func ReadDocument(id string) (string, error) {
	row := storage.DB.QueryRow("SELECT content FROM documents WHERE id=$1", id)

	var content string
	err := row.Scan(&content)

	if err != nil {
		log.Print("Failed to read document")
		return "", err
	}

	return content, nil
}

func ReadDocumentTitle(id string) (string, error) {
	row := storage.DB.QueryRow("SELECT title FROM documents WHERE id=$1", id)

	var content string
	err := row.Scan(&content)

	if err != nil {
		log.Print("Failed to read document")
		return "", err
	}

	return content, nil
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
