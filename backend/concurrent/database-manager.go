package concurrent

import "github.com/i-use-mint-btw/services"

type DatabaseManager struct {
	Incoming chan *DatabaseMessage
}

type DatabaseMessage struct {
	messageType *int // could potentially accept different kinds of messages like updates, deletions etc (may need if extending functionality)
	documentUpdate *DocumentUpdate
}

type DocumentUpdate struct {
	documentID string
	content []byte
}

func NewDatabaseManager() *DatabaseManager {
	return &DatabaseManager{
		Incoming: make(chan *DatabaseMessage),
	}
}

func (dm *DatabaseManager) Run() {
	for {
		select {
		case message, ok := <- dm.Incoming: 
			// the hub closed the channel
			if !ok {
				return
			}
			services.UpdateDocument(message.documentUpdate.documentID, message.documentUpdate.content)
		}
	}
}