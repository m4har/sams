// internal/domain/dto/user_dto.go
package dto

import "time"

// UserDTO represents the data structure for user responses
// @Description User information
type UserDTO struct {
	// User ID
	ID uint `json:"id" example:"1"`
	// User email address
	Email string `json:"email" example:"user@example.com"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at"`
}

// RegisterUserDTO represents the data structure for user registration
type RegisterUserDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginUserDTO represents the data structure for user login
type LoginUserDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponseDTO represents the data structure for login response
type LoginResponseDTO struct {
	Token string  `json:"token"`
	User  UserDTO `json:"user"`
}
