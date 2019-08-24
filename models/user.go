package models

// easyjson:json
type User struct {
	Username string `json:"username"`
	UserID   int32  `json:"user_id"`
}

// easyjson:json
type UserChat struct {
	UserID int32 `json:"user"`
}
