package services

import (
	"backend/internal/models"
	"crypto/sha1"
	"errors"
	"fmt"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	tokenTTL   = 12 * time.Hour
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

// UserRepository represents the interface for interacting with user data in the database.
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
}

// AuthService handles user authentication.
type AuthService struct {
	userRepo UserRepository
}

// NewAuthService creates a new AuthService instance.
func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

// RegisterUser registers a new user.
func (s *AuthService) RegisterUser(input models.UserInput) (*models.User, error) {
	// Validate user input, such as checking for empty fields or invalid email format.
	if input.Username == "" || input.Password == "" || input.Email == "" {
		return nil, errors.New("username, password, and email are required")
	}

	// Check if the user already exists in the database.
	existingUser, _ := s.userRepo.GetUserByUsername(input.Username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// If the user does not exist, create a new user in the database.
	input.Password = generatePasswordHash(input.Password)

	newUser := &models.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		ID:       uuid.NewV4().String(),
	}

	err := s.userRepo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// LoginUser authenticates a user and returns a token.
func (s *AuthService) LoginUser(username, password string) (string, error) {
	// Validate the input, such as checking for empty fields.
	if username == "" || password == "" {
		return "", errors.New("username and password are required")
	}

	// Check if the user exists in the database and if the password is correct.
	password = generatePasswordHash(password)

	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil || user.Password != password {
		return "", errors.New("invalid username or password")
	}

	// Generate and return an authentication token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

// generatePasswordHash generates a password hash
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
