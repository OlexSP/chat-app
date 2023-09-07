package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"encoding/json"
	"net/http"
)

// UserHandler handles user-related operations.
type UserHandler struct {
	authService *services.AuthService
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler(authService *services.AuthService) *UserHandler {
	return &UserHandler{
		authService: authService,
	}
}

// RegisterUser handles user registration.
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userInput models.UserInput

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.authService.RegisterUser(userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// LoginUser handles user login.
func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userInput models.UserInput

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.authService.LoginUser(userInput.Username, userInput.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
