package models

// Room represents a chat room in the chat application.
type Room struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// RoomInput represents the input data for creating or updating a chat room.
type RoomInput struct {
	Name string `json:"name"`
}
