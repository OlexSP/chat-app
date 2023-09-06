package models

// User represents a user in the chat application.
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
}

// UserInput represents the input data for creating or updating a user.
type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
