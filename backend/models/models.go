package models

type Document struct {
	Title string `json:"title"`
}

type DocumentUpdate struct {
	DocumentID string `json:"documentID"`
	Content string `json:"content"`
}

type BroadcastMessage struct {
	Message string `json:"message"`
}