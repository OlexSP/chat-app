package handlers

import (
	"backend/internal/services"
	"net/http"

	"github.com/gorilla/websocket"
)

// ChatHandler handles chat-related operations.
type ChatHandler struct {
	chatService *services.ChatService
	upgrader    websocket.Upgrader
}

// NewChatHandler creates a new ChatHandler instance.
func NewChatHandler(chatService *services.ChatService) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // You should implement a more secure origin check.
			},
		},
	}
}

// HandleWebSocketConnection handles WebSocket connections for chat communication.
// @Summary Establish a WebSocket connection for chat communication
// @Description Upgrade the HTTP connection to a WebSocket connection for real-time chat communication
// @Produce  json
// @Success 101 {string} string "WebSocket connection established"
// @Failure 400 {string} string "Failed to establish WebSocket connection"
// @Router /ws [get]
func (h *ChatHandler) HandleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to establish WebSocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	// Handle incoming and outgoing messages for the WebSocket connection.
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// Process the received message, e.g., store it in the database or broadcast it to other clients.

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			break
		}
	}
}
