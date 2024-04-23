package chat

import (
	"time"

	"github.com/google/uuid"
)

const (
	MessageTypePresence = "presence"
	MessageTypeMessage  = "message"
	MessageTypeTyping   = "typing"
	MessageTypeStatus   = "status"
	MessageTypeReset    = "reset"

	MessageTypeOffline = "0"
	MessageTypeOnline  = "1"
)

type Message struct {
	Text       string    `json:"text"`
	Id         uuid.UUID `json:"id"`
	Type       string    `json:"type"`
	Timestamp  time.Time `json:"ts"`
	Sender     uuid.UUID `json:"sender"`
	SenderName string    `json:"sender_name"`

	// this, however, will be the roomId instead of userId as a room may contain more than 2 people.
	Receiver uuid.UUID `json:"room"`
}
