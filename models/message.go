package models

// easyjson:json
type Message struct {
	MessageID int32  `json:"message_id"`
	Chat      int32  `json:"chat"`
	Author    int32  `json:"author"`
	Text      string `json:"text"`
}
