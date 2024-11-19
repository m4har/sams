package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m4har/sams/internal/application/services"
	"github.com/m4har/sams/internal/interfaces/http/handlers"
	"github.com/m4har/sams/internal/interfaces/http/middleware"
)

func SetupRoutes(r *gin.Engine, authService *services.AuthService) {
	authHandler := handlers.NewAuthHandler(authService)

	// Public endpoints
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
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes
		protected := api.Group("/protect")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/me", authHandler.GetProfile)
		}
	}
}
