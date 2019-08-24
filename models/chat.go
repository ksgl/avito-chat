package models

import "time"

// easyjson:json
type Chat struct {
	ChatID    int32     `json:"chat_id"`
	Name      string    `json:"name"`
	Users     []int32   `json:"users,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// easyjson:json
type ChatsArr []*Chat

// easyjson:json
type ChatMessages struct {
	ChatID int32 `json:"chat"`
}
