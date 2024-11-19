package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m4har/sams/internal/application/services"
	"github.com/m4har/sams/internal/interfaces/http/handlers"
)

func SetupRoutes(r *gin.Engine, authService *services.AuthService) {
	authHandler := handlers.NewAuthHandler(authService)

	// @Summary     Test API
	// @Description Health check endpoint
	// @Tags        health
	// @Produce     json
	// @Success     200 {object} map[string]string
	// @Router      /test [get]
	r.GET("/api/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API is working",
			"status":  "success",
		})
	})

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			// @Summary     Register new user
			// @Description Register a new user with email and password
			// @Tags        auth
			// @Accept      json
			// @Produce     json
			// @Param       input body dto.RegisterUserDTO true "Register User Input"
			// @Success     201 {object} map[string]string
			// @Failure    400 {object} map[string]string
			// @Failure    500 {object} map[string]string
			// @Router      /auth/register [post]
			auth.POST("/register", authHandler.Register)

			// @Summary     Login user
			// @Description Login with email and password
			// @Tags        auth
			// @Accept      json
			// @Produce     json
			// @Param       input body dto.LoginUserDTO true "Login User Input"
			// @Success     200 {object} dto.LoginResponseDTO
			// @Failure    400 {object} map[string]string
			// @Failure    401 {object} map[string]string
			// @Router      /auth/login [post]
			auth.POST("/login", authHandler.Login)
		}
	}
}
