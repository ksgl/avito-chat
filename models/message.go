package models

import "time"

// easyjson:json
type Message struct {
	MessageID int32     `json:"message_id"`
	Chat      int32     `json:"chat"`
	Author    int32     `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// easyjson:json
type MessagesArr []*Message
