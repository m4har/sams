package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/m4har/sams/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Bearer token format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		userID, err := utils.ValidateToken(parts[1])
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user ID in context
		c.Set("user_id", userID)
		c.Next()
	}
}
