package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/m4har/sams/internal/application/services"
	"github.com/m4har/sams/internal/infrastructure/db"
	"github.com/m4har/sams/internal/infrastructure/persistence"
	"github.com/m4har/sams/internal/interfaces/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Initialize database connection
	dbConn, err := db.NewPostgresConnection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		log.Printf("Check your database configuration in .env file:")
		log.Printf("DB_HOST=%s", os.Getenv("DB_HOST"))
		log.Printf("DB_PORT=%s", os.Getenv("DB_PORT"))
		log.Printf("DB_NAME=%s", os.Getenv("DB_NAME"))
		log.Printf("DB_USER=%s", os.Getenv("DB_USER"))
		log.Fatalf("Application cannot start without database connection")
	}
	defer dbConn.Close()

	log.Println("Successfully connected to database")

	userRepo := persistence.NewUserRepository(dbConn)
	authService := services.NewAuthService(userRepo)
	server := http.NewServer(authService)

	log.Println("Starting server on :8080")
	if err := server.Start(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
