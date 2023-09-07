package models

import "time"

// Message represents a chat message in the chat application.
type Message struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	RoomID    int64     `json:"room_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// MessageInput represents the input data for creating a message.
type MessageInput struct {
	UserID  string `json:"user_id"`
	RoomID  int64  `json:"room_id"`
	Content string `json:"content"`
}
