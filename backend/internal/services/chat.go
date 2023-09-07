package services

import (
	"backend/internal/models"
	"errors"
	"time"
)

// MessageRepository represents the interface for interacting with message data in the database.
type MessageRepository interface {
	CreateMessage(message *models.Message) error
	GetMessagesByRoomID(roomID int64) ([]*models.Message, error)
}

// ChatService handles chat interactions.
type ChatService struct {
	messageRepo MessageRepository
}

// NewChatService creates a new ChatService instance.
func NewChatService(messageRepo MessageRepository) *ChatService {
	return &ChatService{
		messageRepo: messageRepo,
	}
}

// SendMessage sends a chat message to a room.
func (s *ChatService) SendMessage(input models.MessageInput) (*models.Message, error) {
	// Validate the input, by checking for empty fields.
	if input.UserID == "" || input.RoomID == 0 || input.Content == "" {
		return nil, errors.New("user ID, room ID, and content are required")
	}

	// Create a new message in the database.
	newMessage := &models.Message{
		UserID:    input.UserID,
		RoomID:    input.RoomID,
		Content:   input.Content,
		Timestamp: time.Now(),
	}

	err := s.messageRepo.CreateMessage(newMessage)
	if err != nil {
		return nil, err
	}

	return newMessage, nil
}

// GetMessages retrieves chat messages for a room.
func (s *ChatService) GetMessages(roomID int64) ([]*models.Message, error) {
	// Retrieve messages from the database for the specified room.
	messages, err := s.messageRepo.GetMessagesByRoomID(roomID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
