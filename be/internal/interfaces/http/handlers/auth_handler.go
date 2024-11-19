package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m4har/sams/internal/application/services"
	"github.com/m4har/sams/internal/domain/dto"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterUserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.Register(req.Email, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginUserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetProfile retrieves the authenticated user's profile
// @Summary Get user profile
// @Description Get currently authenticated user's profile
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.UserDTO
// @Failure 401 {object} map[string]string
// @Router /me [get]
func (h *AuthHandler) GetProfile(c *gin.Context) {
	// Get userID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Get user from repository
	user, err := h.authService.GetUserByID(uint(userID.(uint)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user profile"})
		return
	}

	// Return user profile
	c.JSON(http.StatusOK, dto.UserDTO{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
