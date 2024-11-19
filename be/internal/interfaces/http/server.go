package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/m4har/sams/docs" // Import swagger docs
	"github.com/m4har/sams/internal/application/services"
	"github.com/m4har/sams/internal/interfaces/http/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           SAMS API
// @version         1.0
// @description     This is a sample server for SAMS application.
// @host            localhost:8080
// @BasePath        /api/v1
// @schemes         http
// @produce         json
// @consumes        json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type Server struct {
	router      *gin.Engine
	authService *services.AuthService
}

func NewServer(authService *services.AuthService) *Server {
	server := &Server{
		router:      gin.Default(),
		authService: authService,
	}

	// Add swagger route
	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(server.router, authService)
	return server
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
