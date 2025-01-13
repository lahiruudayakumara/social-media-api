package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"social-media-api/config"
	"social-media-api/routes"
	"social-media-api/utils"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get and validate JWT_SECRET
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatalf("JWT_SECRET environment variable is not set")
	}
	utils.SetJWTSecret(jwtSecret)

	// Initialize Database
	config.InitDB()

	// Create Gin Router
	r := gin.Default()

	// Set Up Routes
	routes.SetupAuthRoutes(r)
	routes.SetupPostRoutes(r)
	routes.SetupUserRoutes(r)

	// Run Server
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
